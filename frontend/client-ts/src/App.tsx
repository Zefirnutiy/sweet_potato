import { DepartmentManagement } from './pages/DepartmentManagement/DepartmentManagement';
import { Navbar } from './components/common/Navbar/Navbar';
import { InformationCard } from './components/cards/Cards';
import { Navigate, Route, Routes } from 'react-router-dom';
import { changePages, changeSchema } from './middleware/change';
import { useState } from 'react';
import { useAuth } from './hooks/auth.hook';
import { AuthContext } from './context/authContext';



const App = () => {
  
  changeSchema()
  const organizationLinks = [
    {icon: "fa fa-cubes", title: "Курсы", path: "#"},
    {icon: "fa fa-brain", title: "Тесты", path: "#"},
    {icon: "fa fa-square-poll-horizontal", title: "Результаты", path: "#"},
    {icon: "fa fa-address-book", title: "Управление", path: "/department"},
    {icon: "fa fa-gear", title: "Аккаунт", path: "#"},
    {icon: "fa fa-circle-info", title: "Информафия", path: "#"},
    {icon: "fa fa-chart-column", title: "Статистика", path: "#"},
  ]

  const clientLinks = [
    {icon: "fa fa-cubes", title: "Курсы", path: "#"},
    {icon: "fa fa-brain", title: "Тесты", path: "#"},
    {icon: "fa fa-square-poll-horizontal", title: "Результаты", path: "#"},
    {icon: "fa fa-gear", title: "Аккаунт", path: "#"},
    {icon: "fa fa-circle-info", title: "Информафия", path: "#"},
  ]

  const guestLinks = [
    {icon: "fa fa-cubes", title: "Главная", path: "#"},
    {icon: "fa fa-brain", title: "Отзывы", path: "#"},
    {icon: "fa fa-square-poll-horizontal", title: "Документация API", path: "#"},
    {icon: "fa fa-address-book", title: "Скриншоты", path: "#"},
  ]
  const [page, setPage] = useState("organization")
    changePages(setPage)
    const {login, logout, token, userId} = useAuth()
    const isAuthenticated = !!token
    
    if(page === "organization" || isAuthenticated){

      return (
      <AuthContext.Provider value={{ login, token, userId, logout, isAuthenticated }}>
        <Navbar links={organizationLinks} authorized={true}>
          <InformationCard title="ТЫ ОРГАНИЗАЦИЯ" message="Храни тебя бог"/>
        </Navbar>
        <Routes>
          <Route path="/" element={<DepartmentManagement/>} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </AuthContext.Provider>
      )
    }
  
    if(page === "user"){
      return (
        <>
          <Navbar links={clientLinks} authorized={true}>
          <InformationCard title="ТЫ ПОЛЬЗОВАТЕЛЬ" message="Храни тебя бог"/>
          </Navbar>
          <Routes>
            <Route path="/" element={<div></div>} />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
        </>
      )
    }

    if(page === "guest"){
    
      return (
        <AuthContext.Provider value={{ login, token, userId, logout, isAuthenticated }}>
          <Navbar links={guestLinks} authorized={false}>
          <InformationCard title="ТЫ ГОСТЬ" message="Храни тебя бог"/>
          </Navbar>
          <Routes>
            <Route path="/" element={<div></div>} />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
        </AuthContext.Provider>
      )
    }
    return null
}

export default App;


