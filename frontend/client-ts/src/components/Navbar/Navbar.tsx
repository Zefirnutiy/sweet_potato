
import { Link } from "react-router-dom"
import { InformationCard } from '../InformationCard/InformationCard';
import "./Navbar.css"
import "../../fonts/fontawesome-free-6.1.2-web/css/all.css"
import { FC, useState } from "react";
import { Modal } from "../Modal/Modal";

type ILink = {
    icon: string
    title: string
    path: string
}

type PropTypes = {
    links: ILink[]
    organizationName?: string
    children?: React.ReactNode
  };

export const Navbar: FC<PropTypes> = ({links, organizationName='Wains', children}) =>{
    const [openRegModal, setRegModal] = useState(false)
    return (
        <div id="navbar">
        <div id="logo">
           {organizationName}
        </div>
        <div id="nav">
            {links?.map(link => 
                <Link to={link.path}><i className={link.icon}></i> {link.title}</Link> 
            )}
           
        </div>
        <div id="navContent">
            {children}
            <button onClick={() => setRegModal(true)}>Окно регистрации</button>
            <Modal openModal={openRegModal} setActiveModal={setRegModal} headerText={"Регистрация"}>
                <input type="text" placeholder="Название организации"/>
                <input type="text" placeholder="Эл.почта"/>
                <input type="text" placeholder="Пароль"/>
                <input type="text" placeholder="Повтор пароля"/>
                <button>Готово</button>
            </Modal>
        </div>
    </div>
    )
}

