import "./Greetings.css"
import {useNavigate} from 'react-router-dom'
import {ReactComponent as Chef} from "../assets/Chef Prepares Food.svg"
import {ReactComponent as Stars} from "../assets/stars.svg"
import {ReactComponent as Negr} from "../assets/negr.svg"
import {ReactComponent as GitHub} from "../assets/github.svg"
import {ReactComponent as Telegram} from "../assets/telegram.svg"
import {ReactComponent as Copyright} from "../assets/copyright.svg"
import {ReactComponent as Logo} from "../assets/Frame 6.svg"

const Greetings = () =>{
    
    const navigate = useNavigate();

    const aboutRedirect = () => {
        navigate("/")
    }

    const recieptsRedirect = () => {
        navigate("/main")
    }

    const getStartRedirect = () => {
        navigate("/sign-up")
    }

    return(
        <div className="greetings"> 

            <div className="header">
                <div className="culina-name">
                    <h1 className="culina">Culina</h1>
                </div>
                <div className="middle">
                    <a onClick={() => {aboutRedirect()}}>About</a>
                    <a onClick={() => {recieptsRedirect()}}>Recipes</a>
                    <a href="#">Pricing</a>
                </div>
                <div className="get-start">
                    <button onClick={() => {getStartRedirect()}}>Get Started</button>            
                </div>
            </div>
            
            <div className="main">
                <div className="main-left-side">
                    <h1 className="main-title">Become<br/> cooking <br/>magician</h1>
                    <Stars/>
                    <p className="left-side-text">5/5 our users making <br/> their first food perfectly</p>
                </div>
                <div className="main-mid-side">
                    <Chef className="chef-svg"/>
                </div>
                <div className="main-right-side">
                    <h2 className="right-side-number">50</h2>
                    <p className="right-side-text-users">current-users</p>
                    <p className="right-side-line"></p>
                    <h2 className="right-side-number">210</h2>
                    <p className="right-side-text-recipes">recipes for all</p>
                    <p className="">need</p>                
                </div>
            </div>

            <div className="info">
                <div className="info-left-side">
                    <div className="black-container">
                        <Negr/>
                        <h2 className="black-container-name">Ainsley Harriott</h2>
                        <p className="black-container-text">Everybody’s got the ability to cook. <br/>They just have to be shown!</p>
                    </div>
                </div>
                <div className="info-right-side">
                    <h2 className="info-right-side-text">Let's cook <br/> together!</h2>
                    <button className="info-right-side-btn" onClick={() => {getStartRedirect()}}>Join us</button>
                </div>
            </div>

            {/* <div className="footer">
                <div className="footer-left-side">
                    <Logo className="footer-left-side-logo"/>
                    <div className="footer-left-side-icons"> 
                        <a href = "#">    
                            <Telegram className="footer-left-side-icon"/>
                        </a>
                        <a href = "#">  
                            <GitHub className="footer-left-side-icon"/>
                        </a>  
                    </div>
                </div>
                <div className="footer-right-side">
                    <div className="contact-us">
                        <h1>Contact Us</h1>
                    </div>
                    <div className="footer-email">
                        <p className = "p-email">qwaq.dev@gmail.com</p>
                        <p className = "p-email">gird.class@gmail.com</p>
                    </div>
                </div>
            </div> */}
            <p className="footer-line"></p>
            <div className="under-footer">
                <div className="left-side">
                    <Copyright className="copyright-svg"/> 
                </div>
                <div className="under-footer-right-side">
                    <a href = "#" >Terms of Service</a>
                </div>
            </div>
        </div>

    );
}

export default Greetings