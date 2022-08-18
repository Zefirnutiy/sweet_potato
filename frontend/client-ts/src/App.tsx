import { DepartmentManagement } from './pages/DepartmentManagement/DepartmentManagement';
import { Navbar } from './components/Navbar/Navbar';
import { InformationCard } from './components/InformationCard/InformationCard';
import { Navigate, Route, Routes } from 'react-router-dom';
import { UserManagement } from './pages/UserManagement/UserManagement';



const App = () => {
  let links = [
    {icon: "fa fa-cubes", title: "Курсы", path: "#"},
    {icon: "fa fa-brain", title: "Тесты", path: "#"},
    {icon: "fa fa-square-poll-horizontal", title: "Результаты", path: "#"},
    {icon: "fa fa-address-book", title: "Пользователи", path: "/departmentManagement"},
    {icon: "fa fa-gear", title: "Аккаунт", path: "#"},
    {icon: "fa fa-circle-info", title: "Информафия", path: "#"},
    {icon: "fa fa-chart-column", title: "Статистика", path: "#"},
  ]

  if(true){
    return (
    <main>
      <Navbar links={links}>
        <InformationCard title="ТЫ ПИДОР" message="Ты таким родился."/>
      </Navbar>
      <Routes>
        <Route path="/departmentManagement" element={<DepartmentManagement/>} />
        <Route path="/userManagement" element={<UserManagement/>} />
        <Route path="*" element={<Navigate to="/departmentManagement" />} />
      </Routes>
    </main>
    )
  }

  return ( 
    <h1>asdasd</h1>
  )
}

export default App;