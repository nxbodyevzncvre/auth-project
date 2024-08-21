import './App.css';
import React, {useState} from 'react';
import {
  BrowserRouter,
  Route,
  Routes,
 } from "react-router-dom";

import SignIn from "./components/SignIn/SignIn";
import SignUp from './components/SignUp/SignUp';
import Greetings from "./components/Greetings/Greetings";
import Main from './components/Main/Main';

function App() {
  return (
    <BrowserRouter>

      <div className="container">
       <Routes>
          <Route path = "/" element = {<Greetings/>}></Route> 
          <Route path = "/sign-in" element = {<SignIn/>}></Route>
          <Route path = "/sign-up" element = {<SignUp/>}></Route>
          <Route path = "/main" element = {<Main/>}></Route>
       </Routes>
      </div>
    </BrowserRouter>

  );
}

export default App;
