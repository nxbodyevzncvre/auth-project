import "./Main.css"
import Card from "./Card/Card"
import Select from "../Select/Select"
import Filter from "../Filter/Filter"
import {useEffect, useState} from "react";
import { useNavigate } from "react-router-dom"
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
import axios from "axios"



const Main = () =>{
    const navigate = useNavigate();
    const [cardInfo, setCardInfo] = useState([]);
    const [username, setUsername] = useState("")
    const getInfo = async() =>{
        try{
            const token = localStorage.getItem("token");
            const headers = {Authorization: `Bearer ${token}`}
            const response = await axios.get("http://localhost:8080/profile", {headers})
            setUsername(response.data.username)
        }catch(err){                                                                                                                                                                            
            console.err(err)
        }
    }
    const PrevArrow = (props) =>{
        const {className, onClick} = props;
        return(
            <div onClick = {onClick} className = {`arrow ${className}`}>
                <AiOutlineArrowLeft class="arrows" style = {{color: "white"}}/>
            </div>
        )
    }
    useEffect(() =>{
        const token = localStorage.getItem("token");
        if(!token){
            navigate("/");
            console.log("you are not authed")
        }
        getInfo()
        // axios.get("http://localhost:8080/cards")
        //     .then(data => setCardInfo(data.data))
        //     .catch(err => console.err(err))
        

        
    }, [])
    const NextArrow = (props) =>{
        const {className, onClick} = props;
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
    

    const aboutRedirect = () => {
        navigate("/")
    }

    const recieptsRedirect = () => {
        navigate("/main")
    }

    const cardRedirect = () => {
        navigate("/")
    }

    return(
        <div className="main-window">
            <div className="header-main">
                <div className="culina-name">
                    <Logo_title className = "culina-logo"/>
                </div>
                <div className="middle-main">
                    <a onClick={() => {aboutRedirect()}}>About</a>
                    <a onClick={() => {recieptsRedirect()}}>Recipes</a>
                    <a href="#">Pricing</a>
                </div>
                <div className="user-info">
                    <div className="user-logo">
                        <Delete_asap className = "logo-svg"/>

                    </div>
                    <div className="user-data">
                        <p className = "user-name">{username}</p>
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
                        <Select values = {["Newest", "Latest", "Rating"]}/>
                    </div>
            </div>

            <div className="main-block">
                <div className="main-filter-block">
                    <div className="main-filter-block-top">
                        <h2 className="main-filter-block-title">Filter by:</h2>
                        <a href="#" className="main-filter-block-clear">clear filters</a>
                    </div>
                    <div className="select-filter">
                        <Filter nameFilter="Difficult:" values = {["5 stars", "4 stars", "3 stars", "2 stars", "1 stars"]}/>
                        <Filter nameFilter="Type:" values = {["Breakfast", "Drink", "Lanch", "Dinner", "Another"]}/>
                        <Filter nameFilter="Time:" values = {["20 min", "30 min", "60 min", "90 min", "120 min", "120 min +"]}/>
                    </div>
                </div>
                <div className="main-card-block">       
                    {cardInfo.map(el => {
                        return <Card dish_name={el.dish_name} dish_rating = {el.dish_rating} dish_creator = {el.dish_creator} dish_descr = {el.dish_descr} dish_types = {el.dish_types}/>
                    })}
                </div>
            </div>
        </div>
    );
}

export default Main