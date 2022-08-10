import React from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import { LoginPage } from './pages/LoginPage';
import { RegisterPage } from './pages/RegisterPage';
import { Account } from './pages/Account/Account';
import { Navbar } from './components/Navbar/Navbar';
import { InformationCard } from './components/InformationCard/InformationCard';



const App = () => {

  if(false){

  let links = [
    {icon: "fa fa-cubes", title: "Курсы", path: "#"},
    {icon: "fa fa-brain", title: "Тесты", path: "#"},
    {icon: "fa fa-square-poll-horizontal", title: "Результаты", path: "#"},
    {icon: "fa fa-address-book", title: "Пользователи", path: "#"},
    {icon: "fa fa-gear", title: "Аккаунт", path: "#"},
    {icon: "fa fa-circle-info", title: "Информмафия", path: "#"},
    {icon: "fa fa-chart-column", title: "Статистика", path: "#"},
  ]

  if(true){

    return (
    <main>
      <Navbar links={links}>
        <InformationCard title="ТЫ ПИДОР" message="Ты таким родился."/>
      </Navbar>
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
