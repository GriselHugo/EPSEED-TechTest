import React, { useState } from 'react';
import { BrowserRouter } from 'react-router-dom';

import Router from '../Router/Router';
import Navbar from '../Navbar/Navbar';
import Login from '../Login/Login';

import './App.css';

import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  const [isLogged, setIsLogged] = useState(parseInt(localStorage.getItem("isLogged") || 0));

  return (
    isLogged ? (
      <BrowserRouter>
        <Navbar />
        <Router />
      </BrowserRouter>
    ) : (
      <Login setIsLogged={setIsLogged} />
    )
  );
}

export default App;
