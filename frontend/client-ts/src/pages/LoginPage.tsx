/* eslint-disable jsx-a11y/anchor-is-valid */
import { Link } from "react-router-dom"
import "../styles/css/login.css"

export const LoginPage = () => {
    return (
        <main>
        <div className="container">
            <div className="information">
                <h1>Добро пожаловать в систему тестирования</h1>
                <div>
                    При создании тестировочного аккаунта, у вас будут возможности:
                    <ul>
                        <li>
                            Создать 1 курс
                        </li>
                        <li>
                            Создать 1 тест
                        </li>
                        <li>
                            Создать 1 учителя
                        </li>
                        <li>
                            Создать 1 студента
                        </li>
                        <li>
                            Открывать тесты
                        </li>
                        <li>
                            Проверять тесты
                        </li>
                    </ul>
                </div> 
            </div>
        </div>
        <div className="registration-container">
           <div className="switch">
                <input type="radio" id="user" name="switch" readOnly/><label htmlFor="user" className="switch-client-radio">Пользователь</label>
                <input type="radio" id="organization" name="switch" readOnly/><label htmlFor="organization" className="switch-client-radio">Организация</label>
           </div>
            <div className="registration">
                <form action="#" className="registrationForm">
                    <span>Email</span>
                    <input type="text" />
                    <span>Пароль</span>
                    <input type="text" />
                    <button type="submit">Войти</button>
                </form>
                <hr/>
                <small><Link to={'/register'}>Зарегистрировать тестировочный аккаунт</Link></small>
            </div>
        </div>
    </main>
    )
}