import { DepartmentManagement } from './pages/DepartmentManagement/DepartmentManagement';
import { Navbar } from './components/common/Navbar/Navbar';
import { InformationCard } from './components/cards/Cards';
import { Navigate, Route, Routes } from 'react-router-dom';
import { UserManagement } from './pages/UserManagement/UserManagement';
import { mirage } from './middleware/mirage';

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

  if(true){
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

  return ( 
    <h1>asdasd</h1>
  )
}

export default App;


function changeSchema(){
  document.addEventListener('keydown', function(event) {
    if (event.code === 'KeyZ') {
      document.documentElement.style.setProperty("--main-hue", "250" )
      document.documentElement.style.setProperty("--accent-hue", "40" )
    }else if (event.code === 'KeyX') {
      document.documentElement.style.setProperty("--main-hue", "50" )
      document.documentElement.style.setProperty("--accent-hue", "160" )
    }else if (event.code === 'KeyC') {
      document.documentElement.style.setProperty("--main-hue", "150" )
      document.documentElement.style.setProperty("--accent-hue", "240" )
    }else if (event.code === 'KeyV') {
      document.documentElement.style.setProperty("--main-hue", "200" )
      document.documentElement.style.setProperty("--accent-hue", "300" )
    }
  });
}