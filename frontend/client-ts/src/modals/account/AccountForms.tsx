import { FC } from "react";
import st from './AccountForms.module.scss';

interface PropTypesLogin {
    sendingEvent: () => void
    setLoginModal: (set: boolean) => void
    openLoginAndReg: () => void
}

interface PropTypesRegister {
    sendingEvent: () => void
    setRegModal: (set: boolean) => void
    openLoginAndReg: () => void
}

export const LoginForm: FC<PropTypesLogin> = ({sendingEvent, setLoginModal, openLoginAndReg}) => {
    function eventComplite(){
        setLoginModal(false)
        sendingEvent()
    }
    return(
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
                    <button className={st["button-done"]} onClick={eventComplite}>Войти</button>
                    <p>У вас нет аккаунта? <button className={st["button-link"]} onClick={()=> openLoginAndReg()}>Регистрация</button></p>
                </div>
            </div>
        </div>
    )
}

export const RegistrationForm: FC<PropTypesRegister> = ({sendingEvent, setRegModal, openLoginAndReg}) => {
    function eventComplite(){
        setRegModal(false)
        sendingEvent()
    }
    return(
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
                <button className={st["button-done"]} onClick={eventComplite}>Готово</button>
                <p>У вас уже есть аккаунт? <button className={st["button-link"]} onClick={()=> openLoginAndReg()}>Войти</button></p>
            </div>
        </div>
    </div>
    )
}
   
