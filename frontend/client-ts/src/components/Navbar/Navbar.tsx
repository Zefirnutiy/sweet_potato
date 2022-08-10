
import { Link } from "react-router-dom"
import { InformationCard } from '../InformationCard/InformationCard';
import "./Navbar.css"
import "../../fonts/fontawesome-free-6.1.2-web/css/all.css"

type PropTypes = {
    links?: object
    organizationName?: string
  };

export const Navbar: React.FC<PropTypes> = ({links, organizationName='Wains'}) =>{
    return (
        <div id="navbar">
        <div id="logo">
           {organizationName}
        </div>
        <div id="nav">
            <Link to={"#"}><i className="fa fa-cubes"></i> Курсы</Link>
            <Link to={"#"}><i className="fa fa-brain"></i> Тесты</Link>
            <Link to={"#"}><i className="fa fa-square-poll-horizontal"></i> Результаты</Link>
            <Link to={"#"}><i className="fa fa-address-book"></i> Пользователи</Link>
            <Link to={"#"}><i className="fa fa-gear"></i> Аккаунт</Link>
            <Link to={"#"}><i className="fa fa-circle-info"></i> Информмафия</Link>
            <Link to={"#"}><i className="fa fa-chart-column"></i> Статистика</Link>
        </div>
        <div id="navContent">
            <InformationCard title="ТЫ ПИДОР" message="Ты таким родился."/>
        </div>
    </div>
    )
}