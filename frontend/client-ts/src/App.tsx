import React from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import { LoginPage } from './pages/LoginPage';
import { RegisterPage } from './pages/RegisterPage';
import { Account } from './pages/Account/Account';
import { Navbar } from './components/Navbar/Navbar';



const App = () => {
  let links = [
    {icon: "ff", title: "Something1"},
    {icon: "cc", title: "Something2"},
  ]

  if(true){
    return (
    <main>
      <Navbar/>
      <Routes>
        <Route path="/account" element={<Account />} />
        <Route path="*" element={<Navigate to="/account" />} />
      </Routes>
    </main>
    )
  }
   
  return ( 
    <Routes>
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="*" element={<Navigate to="/login" />} />
    </Routes>
  )
}

export default App;
