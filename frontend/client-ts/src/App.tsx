import { DepartmentManagement } from './pages/DepartmentManagement/DepartmentManagement';
import { Navbar } from './components/common/Navbar/Navbar';
import { InformationCard } from './components/cards/Cards';
import { Navigate, Route, Routes } from 'react-router-dom';
import { UserManagement } from './pages/UserManagement/UserManagement';



const App = () => {
  let links = [
    {icon: "fa fa-cubes", title: "Курсы", path: "#"},
    {icon: "fa fa-brain", title: "Тесты", path: "#"},
    {icon: "fa fa-square-poll-horizontal", title: "Результаты", path: "#"},
    {icon: "fa fa-address-book", title: "Пользователи", path: "/users"},
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
        <Route path="/department" element={<DepartmentManagement/>} />
        <Route path="/users" element={<UserManagement/>} />
        <Route path="*" element={<Navigate to="/department" />} />
      </Routes>
    </main>
    )
  }

  return ( 
    <h1>asdasd</h1>
  )
}

export default App;