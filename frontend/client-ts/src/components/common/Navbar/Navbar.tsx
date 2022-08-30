
import { Link } from "react-router-dom"
import st from "./Navbar.module.scss"
import { FC, useState } from "react";
import { Modal } from "../Modal/Modal";
import { RegistrationForm } from "../../../modals/account/AccountForms";
import { LoginForm } from "../../../modals/account/AccountForms";

interface ILink {
    icon: string
    title: string
    path: string
}

interface PropTypes{
    links: ILink[]
    authorized: boolean
    organizationName?: string
    children?: React.ReactNode
};

export const Navbar: FC<PropTypes> = ({links, authorized, organizationName='Wains', children}) =>{
    const [openRegModal, setRegModal] = useState(false)
    const [openLoginModal, setLoginModal] = useState(false)

    function openLoginAndReg(){
        setRegModal(!openRegModal)
        setLoginModal(!openLoginModal)
    }

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
        </div>
        {authorized ? 
        <div id={st["wrapper-account"]}>
        <div id={st["user-account"]}>
            <Link to={"#"} id={st["user-avatar"]}></Link>
            <div id={st["user-info"]}>
                <i id={st["bell"]} className="fa-solid fa-bell"></i>
                <Link to={"#"} id={st["user-name"]}>Ульяна Романова</Link>
                <div id={st["user-group"]}>АиВТ/261</div>
            </div>
        </div> </div>: 
        <div id={st["wrapper-account"]}>
            <div id={st["login-buttons"]}>
                <button id={st["button-login"]} onClick={() => setLoginModal(true)}>Войти</button>
                <button id={st["button-registration"]} onClick={() => setRegModal(true)}>Регистрация</button>
            </div>
        </div>
        }
        
        {openRegModal &&
            <Modal cross={true} onClose={() => setRegModal(false)}>
                <RegistrationForm sendingEvent={function(){}} setRegModal={() => setRegModal} openLoginAndReg={openLoginAndReg}/>
            </Modal>}
            {openLoginModal &&
            <Modal cross={true} onClose={() => setLoginModal(false)}>
                <LoginForm sendingEvent={function(){}} setLoginModal={() => setRegModal} openLoginAndReg={openLoginAndReg}/>
            </Modal>}
    </div>
    )
}
