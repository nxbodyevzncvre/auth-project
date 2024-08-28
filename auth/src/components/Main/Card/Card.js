import "./Card.css"

const Card = (props) => {

    return(
        <div className="card">
            <div className="card-img">
                <img src="#" alt="#" />
            </div>
            <div className="card-title">
                {props.dish_name}
            </div>
            <div className="card-bot">
                <div className="card-rate">
                    <span>Rate: </span>
                    <span>{props.dish_rating}</span>
                </div>
                <div className="card-user">
                    <span>User: </span>
                    <span>{props.dish_creator}</span>
                </div>
            </div>
        </div>
    )
}

export default Card