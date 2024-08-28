package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/lib/pq"
	"github.com/nxbodyevzncvre/mypackage/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *sql.DB
func Init() error{
	conf := config.GetConfig()
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
	conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Dbname)

	log.Printf("Connecting to database with DSN: %s", dsn)

	DB, err = sql.Open("postgres", dsn)
	
	if err != nil{
		log.Printf("error opening db %v", err)
	}

	if err = DB.Ping(); err !=nil{
		fmt.Printf("error pingign db %s", err)
		return err
	}
	log.Println("success")
	return nil
}
func MongoClient() *mongo.Client{
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_SRV_RECORD")).SetServerAPIOptions(serverAPIOptions)	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err)
	}
	return client

}
