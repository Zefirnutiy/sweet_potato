
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
        <div id={st["nav-content"]}>
            {/* <button onClick={() => setRegModal(true)}>Выйти</button>
            {openRegModal &&
            <Modal headerText={"Вы точно хотите выйти?"} onClose={() => setRegModal(false)}>
                
                <button>Точно</button>
                <button onClick={() => setRegModal(false)}>Нет</button>

            </Modal>} */}
            <div id={st["login-buttons"]}>
                <button id={st["button-login"]} onClick={() => setLoginModal(true)}>Войти</button>
                <button id={st["button-registration"]} onClick={() => setRegModal(true)}>Регистрация</button>
            </div>
            {openRegModal &&
            <Modal>
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
                            <button onClick={()=> setRegModal(false)}>Готово</button>
                            <p>У вас уже есть аккаунт? <button onClick={()=> setLoginModal(true)}>Войти</button></p>
                        </div>
                    </div>
                </div>
            </Modal>}
            {openLoginModal &&
            <Modal>
                <div className={st["wrapper-login"]}>
                    <div className={st["welcome"]}>ffffffffffffffff</div>
                    <div className={st["registration-form"]}>
                        <label>Эл. почта</label>
                        <input type="text"/>
                        <label>Пароль</label>
                        <input type="text"/>
                    </div>
                    <button onClick={()=> setLoginModal(false)}>Готово</button>
                </div>
            </Modal>}
            {children}
        </div>
    </div>
    )
}
