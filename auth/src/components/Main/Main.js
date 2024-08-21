import "./Main.css";
import {ReactComponent as Logo_title} from "../assets/Logos/logo_title.svg"
import {ReactComponent as Delete_asap} from "../assets/negr.svg"
import Slider from "react-slick";
import {ReactComponent as First} from "../assets/Slides/FirstSlide.svg"
import {ReactComponent as Second} from "../assets/Slides/SecondSlide.svg"
import {ReactComponent as Third} from "../assets/Slides/ThirdSlide.svg"
import {ReactComponent as Fourth} from "../assets/Slides/FourthSlide.svg"
import {ReactComponent as Fifth} from "../assets/Slides/FifthSlide.svg"
import {ReactComponent as SearchLogo} from "../assets/Logos/search-logo.svg"
import {AiOutlineArrowLeft, AiOutlineArrowRight} from "react-icons/ai"
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css"


const Main = () =>{
    const PrevArrow = (props) =>{
        const {className, style, onClick} = props;
        return(
            <div onClick = {onClick} className = {`arrow ${className}`}>
                <AiOutlineArrowLeft class="arrows" style = {{color: "white"}}/>
            </div>
        )
    }
    const NextArrow = (props) =>{
        const {className, style, onClick} = props;
        return(
            <div onClick = {onClick} className = {`arrow ${className}`}>
                <AiOutlineArrowRight class = "arrows" style = {{color: "white"}}/>
            </div>
        )
    }
    const settings = {
        dots: true,
        infinite: false,
        speed: 300,
        slidesToShow: 1,
        slidesToScroll: 1,
        lazyload: true,
        nextArrow: <NextArrow to ="next"/>,
        prevArrow: <PrevArrow to ="prev"/>
    }
    const data = [
        <First className ="slide"/>,
        <Second className = "slide"/>,
        <Third className = "slide"/>,
        <Fourth className = "slide"/>,
        <Fifth className = "slide"/>
    ]
    
    

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
            <section className = "main-section-container">
                <Slider {...settings}>
                    {data.map((d) =>{
                        return d
                    })}
                </Slider>
            </section>

            <div className="recipe-finder-block">
                    <div className="recipe-filter">
                        <h2 className="recipe-name">Recipe</h2>
                    </div>
                    <div className="recipe-search-panel">
                        <a href = "#" className = "search-logo-a">
                            <SearchLogo className = "search-logo"/>
                        </a>
                       
                        <input className="recipe-search" type="text" placeholder="search"/>
                    </div>
                    <div className="main-sort-by">
                        <select name="sort">
                            <optgroup label="Sort By">
                                <option value="newest">Newest</option>
                                <option value="latets">Latest</option>
                                <option value="rating">Rating</option>
                            </optgroup>
                        </select>
                    </div>
            </div>
            

        </div>
    );
}

export default Main