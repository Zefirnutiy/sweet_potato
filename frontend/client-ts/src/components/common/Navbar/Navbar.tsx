
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
    const [openLoginModal, setLoginModal] = useState(false)
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
        <div id={st["login-buttons"]}>
            <button id={st["button-login"]} onClick={() => setLoginModal(true)}>Войти</button>
            <button id={st["button-registration"]} onClick={() => setRegModal(true)}>Регистрация</button>
        </div>
        <div id={st["nav-content"]}>
            {children}
        </div>
        <div id={st["card-user-account"]}>
            <Link to={"#"} id={st["card-user-avatar"]}></Link>
            <div id={st["card-user-info"]}>
                <i id={st["bell"]} className="fa-solid fa-bell"></i>
                <Link to={"#"} id={st["card-user-name"]}>Ульяна Романова</Link>
                <div id={st["card-user-group"]}>АиВТ/261</div>
            </div>
        </div>
        {openRegModal &&
            <Modal cross={true} onClose={() => setRegModal(false)}>
                <div className={st["wrapper-registration"]}>
                    <div className={st["welcome"]}>
                    </div>
                    <div className={st["registration"]}>
                        <h3>Регистрация</h3>
                        <div className={st["registration-form"]}>
                            <div>Название организации</div>
                            <input type="text"/>
                            <div>Эл. почта</div>
                            <input type="text"/>
                            <div>Пароль</div>
                            <input type="text"/>
                            <div>Повтор пароля</div>
                            <input type="text"/>
                            <button className={st["button-done"]} onClick={()=> setRegModal(false)}>Готово</button>
                            <p>У вас уже есть аккаунт? <button className={st["button-link-login"]} onClick={()=> setLoginModal(true)}>Войти</button></p>
                        </div>
                    </div>
                </div>
            </Modal>}
            {openLoginModal &&
            <Modal cross={true} onClose={() => setLoginModal(false)}>
                <div className={st["wrapper-registration"]}>
                    <div className={st["welcome"]}>
                    </div>
                    <div className={st["registration"]}>
                        <h3>Вход</h3>
                        <div className={st["registration-form"]}>
                            <div>Эл. почта</div>
                            <input type="text"/>
                            <div>Пароль</div>
                            <input type="text"/>
                            <button className={st["button-done"]} onClick={()=> setLoginModal(false)}>Войти</button>
                            <p>У вас нет аккаунта? <button className={st["button-link-login"]} onClick={()=> setRegModal(true)}>Регистрация</button></p>
                        </div>
                    </div>
                </div>
            </Modal>}
    </div>
    )
}
