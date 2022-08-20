
import { Link } from "react-router-dom"
import st from "./Navbar.module.scss"
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
        <div id={st["navbar"]}>
        <div id={st["logo"]}>
           <Link to="/departament">{organizationName}</Link>
        </div>
        <div id={st["nav"]}>
            {links?.map(link => 
                <Link to={link.path}><i className={link.icon}></i> {link.title}</Link> 
            )}
        </div>
        <div id={st["nav-content"]}>
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
