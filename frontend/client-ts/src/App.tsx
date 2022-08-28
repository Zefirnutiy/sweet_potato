import { DepartmentManagement } from './pages/DepartmentManagement/DepartmentManagement';
import { Navbar } from './components/common/Navbar/Navbar';
import { InformationCard } from './components/cards/Cards';
import { Navigate, Route, Routes } from 'react-router-dom';
import { UserManagement } from './pages/UserManagement/UserManagement';
import { mirage } from './middleware/mirage';
import { changePages, changeSchema } from './middleware/change';
import { useState } from 'react';
import { UserPage } from './pages/UserPage/UserPage';
import { Login } from './pages/Login/Login';
import { Register } from './pages/Register/Register';

const App = () => {
  mirage()
  changeSchema()
  let links = [
    {icon: "fa fa-cubes", title: "Курсы", path: "#"},
    {icon: "fa fa-brain", title: "Тесты", path: "#"},
    {icon: "fa fa-square-poll-horizontal", title: "Результаты", path: "#"},
    {icon: "fa fa-address-book", title: "Управление", path: "/department"},
    {icon: "fa fa-gear", title: "Аккаунт", path: "#"},
    {icon: "fa fa-circle-info", title: "Информафия", path: "#"},
    {icon: "fa fa-chart-column", title: "Статистика", path: "#"},
  ]
  const [page, setPage] = useState("admin")
    changePages(setPage)

    if(page === "admin"){
      return (
      <main>
        <Navbar links={links}>
          <InformationCard title="ТЫ ХОРОШИЙ ЧЕЛОВЕК" message="Храни тебя бог"/>
        </Navbar>
        <Routes>
          <Route path="/department" element={<DepartmentManagement/>} />
          <Route path="/users" element={<UserManagement/>} />
          <Route path="*" element={<Navigate to="/department" />} />
        </Routes>
      </main>
      )
    }
  
    if(page === "register"){
      return (
        <Routes>
          <Route path="/login" element={<Login/>} />
          <Route path="/register" element={<Register/>} />
          <Route path="*" element={<Navigate to="/login" />} />
        </Routes>
      )
    }

    if(page === "user"){
      return (
        <Routes>
          <Route path="/user" element={<UserPage/>} />
          <Route path="*" element={<Navigate to="/user" />} />
        </Routes>
      )
    }
    return null
}

export default App;


