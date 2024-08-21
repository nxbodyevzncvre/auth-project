import "./Header.css"

import {ReactComponent as Logo} from "../assets/Logos/logo.svg"

const Header = () =>{
    return(
        <div className="header">
            <div className="culina-name">
                 <h1 className="culina">Culina</h1>
            </div>
            <div className="middle">
                <a href="#">About</a>
                <a href="#">Recipes</a>
                <a href="#">Pricing</a>
            </div>
            <div className="logo-container">
                <Logo className="logo"/>            
            </div>
            </div>

    )
}

export default Header