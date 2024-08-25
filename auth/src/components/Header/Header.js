import "./Header.css"

import {ReactComponent as Logo} from "../assets/Logos/logo.svg"
import {useNavigate} from "react-router-dom"

const Header = () =>{
    const navigate = useNavigate();

    const aboutRedirect = () =>{
        navigate("/")
    }

    const recieptsRedirect = () =>{
        navigate("/main")
    }
    return(
        <div className="header">
            <div className="culina-name">
                 <h1 className="culina">Culina</h1>
            </div>
            <div className="middle">
                <a onClick={() => {aboutRedirect()}}>About</a>
                <a onClick={() => {recieptsRedirect()}}>Recipes</a>
                <a href="#">Pricing</a>
            </div>
            <div className="logo-container">
                <Logo className="logo"/>            
            </div>
            </div>

    )
}

export default Header