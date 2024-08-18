import "./Profile.css"
import axios from "axios";
import { useEffect, useState } from "react";

const Profile = () =>{
    let [name, setName] = useState("")
    const [loading, setLoading] = useState(true)
    useEffect(() =>{
        const getProfileData = async() =>{
            try{
                const response = axios.get("http://localhost:8080/profile", {
                    headers:{
                        "Authorization" : `Bearer ${localStorage.getItem("token")}`
                    }
                })
                setName((await response).data.username)
                setLoading(false)
            }catch(error){
                console.error(error)
                setLoading(false)
            }
        };

        getProfileData()

    }, [])
    if (loading){
        return <div>loading...</div>
    }
    return(
        <div className="profile">
            <h1>Hello {name}</h1>
            <button>Get Profile</button>
            </div>
        );


}

export default Profile
