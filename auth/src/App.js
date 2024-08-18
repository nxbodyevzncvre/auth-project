import './App.css';
import React from 'react';
import {
  BrowserRouter,
  Route,
  Routes,
 } from "react-router-dom";

import SignIn from "./components/SignIn/SignIn";
import SignUp from './components/SignUp/SignUp';
import Profile from "./components/MyProfile/Profile";
import Greetings from "./components/Greetings/Greetings";

function App() {
  return (
    <BrowserRouter>

      <div className="container">
       <Routes>
          <Route path = "/" element = {<Greetings/>}></Route> 
          <Route path = "/sign-in" element = {<SignIn/>}></Route>
          <Route path = "/sign-up" element = {<SignUp/>}></Route>
          <Route path = "/profile" element = {<Profile/>}></Route>
       </Routes>
      </div>
    </BrowserRouter>

  );
}

export default App;
