import "./SignIn.css"
import {useNavigate} from "react-router-dom"
import {useState, useEffect} from "react";
import axios from "axios";
import  Header from "../Header/Header";


const SignIn = () =>{
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [notFound, setFound] = useState("found")
    const navigate = useNavigate()

    useEffect(() =>{
        localStorage.setItem("token", undefined)
    }, [])
    const Login = async(username, password) =>{
        try{
            const response = await axios.post("http://localhost:8080/login", {username, password})

            if(response.data === "User not found" || response.data === "Password is incorrect"){
                console.error("User not found or password is incorrect")
                setFound("not-found")
                return
               
            }
            console.log(response.data.access_token)
            const token = response.data.access_token
            if (!token){
                setFound("not-found")
            }

            localStorage.setItem("token", token);
            axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
            console.log("success")
            navigate("/main")
            //set token to local storage
            
        }catch(error){
            console.log(`error - ${error}`);
        }
    }
    const handleSubmit = (e) =>{
        e.preventDefault();
        Login(username, password)
    }

    return(
        <div className="signin">
            <Header/>
                <form className = "form-signin" onSubmit= {handleSubmit}>
                    <h1 className = "login">Login</h1>
                    <input 
                        type="text"
                        name="username" 
                        placeholder = "Username" 
                        value = {username}
                        onChange={(e) =>{
                            setUsername(e.target.value)
                        }}
                        className="input-username" />
                    <input 
                        type="password" 
                        name="password" 
                        placeholder = "Password" 
                        value= {password}
                        onChange ={(e) => {
                            setPassword(e.target.value)
                        }}
                        className="input-password" />
                    <button type="submit" onClick={handleSubmit}>Login</button>
                    <div className = {notFound}>User not found!</div>
                    <p>Forgot password? <a className = "forgot-password" href = "#">click here</a></p>
                </form> 
        </div> 
    )
}

export default SignIn;