import "./Card.css"

const Card = (props) => {

    return(
        <div className="card">
            <div>
                {props.dish_name}
            </div>
        </div>
    )
}

export default Card