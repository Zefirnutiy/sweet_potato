import axios from "axios";
import React, { FC, useContext, useState } from "react";
import { AuthContext } from "../../context/authContext";
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
    /*function eventComplite(){
        setLoginModal(false)
        sendingEvent()
    }*/

    const auth = useContext(AuthContext)
    const [form, setForm] = useState({email: '', password: '', title: '', passwordVerification: ''})

    function changeHandler (event: React.ChangeEvent<HTMLInputElement>)  {
        setForm({...form, [event.target.name]: event.target.value})
    }
    
    const loginHandler = async () => {
        try {
            await axios.post('http://localhost:8080/auth/login', {...form})
            .then((request) =>{
                auth.login(request.data.token, request.data.organization.id)
                setLoginModal(false)
            })
        } catch (e) {console.log(e)}

    }

    return(
        <div className={st["wrapper-registration"]}>
            <div className={st["welcome"]}>
            </div>
            <div className={st["registration"]}>
                <h3>Вход</h3>
                <div className={st["registration-form"]}>
                    <div>Эл. почта</div>
                    <input type="text" onChange={changeHandler}/>
                    <div>Пароль</div>
                    <input type="password" onChange={changeHandler}/>
                    <button className={st["button-done"]} onClick={loginHandler}>Войти</button>
                    <p>У вас нет аккаунта? <button className={st["button-link"]} onClick={()=> openLoginAndReg()}>Регистрация</button></p>
                </div>
            </div>
        </div>
    )
}

export const RegistrationForm: FC<PropTypesRegister> = ({sendingEvent, setRegModal, openLoginAndReg}) => {
    /*function eventComplite(){
        setRegModal(false)
        sendingEvent()
    }*/

    const auth = useContext(AuthContext)
    const [form, setForm] = useState({email: '', password: '', title: '', passwordVerification: ''})

    function changeHandler (event: React.ChangeEvent<HTMLInputElement>)  {
        setForm({...form, [event.target.name]: event.target.value})
    }
    
    const registerHandler = async () => {
        try {
            await axios.post('http://localhost:8080/auth/register', {...form})
            .then((request) =>{
                auth.login(request.data.token, request.data.organization.id)
                setRegModal(false)
            })
        } catch (e) {console.log(e)}

    }


    return(
        <div className={st["wrapper-registration"]}>
        <div className={st["welcome"]}>
        </div>
        <div className={st["registration"]}>
            <h3>Регистрация</h3>
            <div className={st["registration-form"]}>
                <div>Название организации</div>
                <input type="text" name="title" onChange={changeHandler}/>
                <div>Эл. почта</div>
                <input type="text" name="email" onChange={changeHandler}/>
                <div>Пароль</div>
                <input type="password" name="password" onChange={changeHandler}/>
                <div>Повтор пароля</div>
                <input type="password" name="passwordVerification"/>
                <button className={st["button-done"]} onClick={registerHandler}>Готово</button>
                <p>У вас уже есть аккаунт? <button className={st["button-link"]} onClick={()=> openLoginAndReg()}>Войти</button></p>
            </div>
        </div>
    </div>
    )
}
   
