import "./SignUp.css"
import { useState } from "react"
import Header from "../Header/Header";

const SignUp = () =>{
    const [email, setEmail] = useState("")
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e) =>{
        e.preventDefault();
        const data = {
            Username: username,
            Email: email,
            Password: password};
        console.log(data)
        fetch("http://localhost:8080/register", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(() =>{
            console.log("success")
        })
        .catch((err)=>{
            console.error("USER ALREADY EXISTS", err)
        })
        
}
    return(
        <div className="signup">
        <Header/>
            <form className = "form-signup"onSubmit= {handleSubmit}>
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
                <button type="submit">Login</button>
                <p>Already have an account? <a className = "forgot-password" href = "#">click here</a></p>
            </form> 
    </div> 
    )
}

export default SignUp;