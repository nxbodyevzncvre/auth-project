import "./SignUp.css"
import { useState } from "react"
import {useNavigate} from "react-router-dom"
import Header from "../Header/Header";
import axios from "axios";

const SignUp = () =>{
    const [email, setEmail] = useState("")
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [exists, setExists] = useState('not-exists')
    const navigate = useNavigate();
    const loginRedirect = () =>{
        navigate("/sign-in")
    }
    const handleSubmit = async(e) =>{
        e.preventDefault();
        const data = {
            Username: username,
            Password: password,
            Email: email
        };
        console.log(data)
        try{
            const response = await axios.post("http://localhost:8080/register", {username, password, email});
            if (response.data === "E-mail already exists"){
                console.error("Email already exists")
                setExists("exists")
                return
            }
            navigate("/sign-in")
        }catch(error){
            console.log(`${error}`)
        }
    
}
    return(
        <div className="signup">
        <Header/>
            <form className = "form-signup" onSubmit= {handleSubmit}>
                <h1 className = "register">Sign Up</h1>
                <input 
                    type="text"
                    name="e-mail" 
                    placeholder = "E-mail" 
                    value = {email}
                    onChange={(e) =>{
                        setEmail(e.target.value)
                    }}
                    className="input-username" />
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
                <button type="submit">Sign up</button>
                <div className={exists}>User already exists</div>
                <p>Already have an account? <a className = "forgot-password" onClick={() => loginRedirect()}>click here</a></p>
            </form> 
    </div> 
    )
}

export default SignUp;