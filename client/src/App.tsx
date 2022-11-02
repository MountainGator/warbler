import React from 'react';
import './App.scss';
import { BrowserRouter, Routes, Route, redirect } from 'react-router-dom';
import Home from './components/Home/home.component';
import Login from './components/Login/login.component';
import {ApiService} from './services/api.service';

function App() {
  const api = new ApiService();

  React.useEffect(() => {
    const check = async () => {
      const res = await api.checkLogin();
      console.log("check login api res", res)
      if(res.user) {
        redirect("home")
      } else {
        redirect("login")
      }
    }
    check();
  },[])

  return (
    <BrowserRouter>
      <Routes>
        <Route path="home" element={<Home/>} />
        <Route path="login" element={<Login/>} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
