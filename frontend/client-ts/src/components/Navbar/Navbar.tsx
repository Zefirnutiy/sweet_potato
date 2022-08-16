
import { Link } from "react-router-dom"
import "./Navbar.css"
import { FC, useState } from "react";
import { Modal } from "../Modal/Modal";

interface ILink {
    icon: string
    title: string
    path: string
}

interface PropTypes{
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
            <button onClick={() => setRegModal(true)}>Выйти</button>
            {openRegModal &&
            <Modal headerText={"Вы точно хотите выйти?"} onClose={() => setRegModal(false)}>
                
                <button>Точно</button>
                <button onClick={() => setRegModal(false)}>Нет</button>

            </Modal>}
        </div>
    </div>
    )
}

