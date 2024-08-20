import "./Main.css";
import {ReactComponent as Logo_title} from "../assets/logo_title.svg"
import {ReactComponent as Delete_asap} from "../assets/negr.svg"
const Main = () =>{
    return(
        <div className="main-window">
            <div className="header-main">
                <div className="culina-name">
                    <Logo_title className = "culina-logo"/>
                </div>
                <div className="middle-main">
                    <a href="#">About</a>
                    <a href="#">Recipes</a>
                    <a href="#">Pricing</a>
                </div>
                <div className="user-info">
                    <div className="user-logo">
                        <Delete_asap className = "logo-svg"/>

                    </div>
                    <div className="user-data">
                        <p className = "user-name">Negr</p>
                        <p className = "user-amount">12 posts</p>
                        </div>            
                </div>
            </div>
            
        </div>
    );
}

export default Main