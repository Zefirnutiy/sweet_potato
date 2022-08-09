
import { Link } from "react-router-dom"
import "./Navbar.css"
import "../../fonts/fontawesome-free-6.1.2-web/css/all.css"

export const Navbar = () =>{
    return (
        <div id="navbar">
        <div id="logo">
           Wains
        </div>
        <div id="nav">
            <Link to={"#"}><i className="fa fa-cubes"></i> Курсы</Link>
            <Link to={"#"}><i className="fa fa-brain"></i> Тесты</Link>
            <Link to={"#"}><i className="fa fa-square-poll-horizontal"></i> Результаты</Link>
            <Link to={"#"}><i className="fa fa-address-book"></i> Пользователи</Link>
            <Link to={"#"}><i className="fa fa-circle-user"></i> Аккаунт</Link>
            <Link to={"#"}><i className="fa-solid fa-circle-info"></i> Информмафия</Link>
            <Link to={"#"}><i className="fa fa-chart-column"></i> Статистика</Link>
        </div>
        <div id="navContent">

        </div>
    </div>
    )
}