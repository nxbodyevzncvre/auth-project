import "./SignIn.css"
import {useState} from "react";
import axios from "axios";
import  Header from "../Header/Header";
const SignIn = () =>{
    

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const Login = async(username, password) =>{
        try{
            const response = await axios.post("http://localhost:8080/login", {username, password})

            if(response.data === "User not found"){
                console.error("User not found")
                return
            }
            console.log(response.data.access_token)
            const token = response.data.access_token

            //set token to local storage
            localStorage.setItem("token", token);
    

            axios.defaults.headers.common["Authorization"] = `Bearer $token`;
            console.log("success",)
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
                    <p>Forgot password? <a className = "forgot-password" href = "#">click here</a></p>
                </form> 
        </div> 
    )
}

export default SignIn;