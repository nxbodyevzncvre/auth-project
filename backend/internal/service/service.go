package service

import (
	"bytes"
	"context"
	"database/sql"
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	// "os"
	//"regexp"
	"strconv"
	"path/filepath"
	"strings"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nxbodyevzncvre/mypackage/internal/config"
	"github.com/nxbodyevzncvre/mypackage/internal/db"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type AuthHandler struct {
	conf        *config.Config
	authStorage *config.AuthStorage
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		conf:        config.GetConfig(),
		authStorage: &config.AuthStorage{DB: &config.Users{Users: make(map[string]config.User)}},
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	if db.DB == nil {
		fmt.Printf("DB not initialized")
		return c.Status(fiber.StatusInternalServerError).SendString("DB not initialized")
	}
	regReq := config.User{}
	if err := c.BodyParser(&regReq); err != nil {
		return c.SendString(err.Error())
	}

	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", regReq.Email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return c.SendString(err.Error())
	}

	if exists {
		return c.SendString("E-mail already exists")
	}

	_, err = db.DB.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", regReq.Username, regReq.Password, regReq.Email)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString("Success")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	//getting username and password
	regReq := config.User{}

	if err := c.BodyParser(&regReq); err != nil {
		return c.SendString(err.Error())
	}
	//querry for username existance
	var user config.User
	err := db.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", regReq.Username).Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendString("User not found")
		}
		return c.SendString(err.Error())
	}

	//if noert exists return error
	//if exists checking password for login
	if regReq.Password != user.Password {
		return c.SendString("Password is incorrect")
	}
	//JWT TOKEN CREATING

	//1. creating data for token (payload chapter)
	payload := jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(config.JwtSecretKey)
	if err != nil {
		logrus.WithError(err).Error("JWTT token signing")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(config.LoginResponse{Token: t})
}

func JwtPayloadFromRequest(c *fiber.Ctx) (jwt.MapClaims, bool) {
	jwtToken, ok := c.Locals(config.ContextKeyUser).(*jwt.Token)
	
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_context_value": c.Context().Value(config.ContextKeyUser),
		}).Error("wrong type of jwt token in Context")
		return nil, false
	}
	payload, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_claims": jwtToken.Claims,
		}).Error("wrong type of jwt token claims")
		return nil, false
	}
	return payload, true

}

func (h *AuthHandler) Profile(c *fiber.Ctx) error {
	jwtPayload, ok := JwtPayloadFromRequest(c)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)

	}
	rows, err := db.DB.Query("SELECT username FROM users WHERE username = $1", jwtPayload["sub"])
	if err != nil {
		return c.SendString("User not found")
	}

	defer rows.Close()
	var username string
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return c.JSON(config.ProfileResponse{
		Username: username,
	})
}
	
func resetUsedIds(client *redis.Client) error{
	_, err := client.Del(context.Background(), "used_ids").Result()
	return err
}

func GetRandomRedisIds(client *redis.Client) (string, error) {
	for {
		count, err := client.SCard(context.Background(), "all_ids").Result()
		if err != nil {
			return "", err
		}
		if count == 0 {
			if err := resetUsedIds(client); err != nil {
				return "", err
			}
		}

		id, err := client.SRandMember(context.Background(), "all_ids").Result()
		if err != nil {
			return "", err
		}

		exists, err := client.SIsMember(context.Background(), "used_ids", id).Result()
		if err != nil {
			return "", err
		}
		if !exists {
			if err := client.SAdd(context.Background(), "used_ids", id).Err(); err != nil {
				return "", err
			}
			return id, nil
		}
	}
}
func GetRandomRedisID(c *fiber.Ctx) error {
	client := db.RedisClient()
	id, err := GetRandomRedisIds(client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"id":    id,
	})
}


func postRedisId(client *redis.Client, id string) error {
	return client.SAdd(context.Background(), "all_ids", id).Err()
	
}

func PostCard(c *fiber.Ctx) error {
    client := db.MongoClient()
    imageDb := client.Database("culina-image")
    dataDb := client.Database("culina-data")
    redisClient := db.RedisClient()

    var cardInfo config.Card
    if err := c.BodyParser(&cardInfo); err != nil {
        log.Printf("Failed to parse body: %v", err)
        return c.Status(fiber.StatusBadRequest).SendString("Failed to fetch data")
    }

    // Загрузка изображения
    fileHeader, err := c.FormFile("image")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":   "Failed to upload image",
        })
    }

    file, err := fileHeader.Open()
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }
    defer file.Close()

    content, err := io.ReadAll(file)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    bucket, err := gridfs.NewBucket(imageDb)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   "Failed to create bucket",
        })
    }

    metadata := bson.M{
        "ext": strings.ToLower(filepath.Ext(fileHeader.Filename)),
    }

    uploadStream, err := bucket.OpenUploadStreamWithID(
        primitive.NewObjectID(),
        fileHeader.Filename,
        options.GridFSUpload().SetMetadata(metadata),
    )
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }
    defer uploadStream.Close()

    fileSize, err := uploadStream.Write(content)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    imageID := uploadStream.FileID.(primitive.ObjectID)
    cardInfo.ImageID = imageID
	err = postRedisId(redisClient, cardInfo.ImageID.Hex())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   "Failed to add ID to Redis",
        })
    }


    coll := dataDb.Collection("_culina-data")
    _, err = coll.InsertOne(context.TODO(), cardInfo)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "error": false,
        "msg":   "Card and image successfully uploaded",
        "image": fiber.Map{
            "id":   imageID.Hex(),
            "name": fileHeader.Filename,
            "size": fileSize,
            "test": cardInfo.ImageID.Hex(),
        },
    })
}


func setResponseHeaders(c *fiber.Ctx, buff *bytes.Buffer, ext string) error {
	switch ext {
	case ".png":
		c.Set("Content-Type", "image/png")
	case ".jpg":
		c.Set("Content-Type", "image/jpg")
	case ".jpeg":
		c.Set("Content-Type", "image/jpeg")
	}
	c.Set("Cache-Control", "public, max-age=31536000")
	c.Set("Content-Length", strconv.Itoa(len(buff.Bytes())))
	return nil
}


func GetDataCard(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	dataDB := db.MongoClient().Database("culina-data")
	filter := bson.D{{"image_id", id}}
	var result config.Card
	coll := dataDB.Collection("_culina-data")
	err = coll.FindOne(c.Context(), filter).Decode(&result)
	if err != nil {
		return c.SendString("Not allowed to find data")
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"dish_name":    result.Dish_name,
			"dish_rating":  result.Dish_rating,
			"dish_creator": result.Dish_creator,
			"dish_descr":   result.Dish_descr,
			"dish_types":   result.Dish_types,
		},
	})
}



func GetImgCard(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid image ID format",
		})
	}

	db := db.MongoClient().Database("culina-image")

	var avatarMetadata bson.M
	err = db.Collection("fs.files").FindOne(c.Context(), bson.M{"_id": id}).Decode(&avatarMetadata)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Avatar not found",
		})
	}

	metadata, ok := avatarMetadata["metadata"].(bson.M)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Image metadata is missing or incorrect",
		})
	}

	ext, ok := metadata["ext"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "File extension not found in metadata",
		})
	}

	var buffer bytes.Buffer
	bucket, err := gridfs.NewBucket(db, options.GridFSBucket().SetName("fs"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to create GridFS bucket",
		})
	}

	_, err = bucket.DownloadToStream(id, &buffer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to download image",
		})
	}

	setResponseHeaders(c, &buffer, ext)

	return c.Send(buffer.Bytes())
}
