/* eslint-disable jsx-a11y/anchor-is-valid */
import { Link } from "react-router-dom"
import "../styles/css/login.css"

export const LoginPage = () => {
    return (
        <main>
        <div class="container">
            <div class="information">
                <h1>Добро пожаловать в систему тестирования</h1>
                <p>
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
                </p> 
            </div>
        </div>
        <div class="registration-container">
           <div class="switch">
                <input type="radio" id="user" name="switch" checked/><label for="user" class="switch-client-radio">Пользователь</label>
                <input type="radio" id="organization" name="switch"/><label for="organization" class="switch-client-radio">Организация</label>
           </div>
            <div class="registration">
                <form action="#" class="registrationForm">
                    <span>Email</span>
                    <input type="text"/> 
                    <span>Пароль</span>
                    <input type="text"/>
                    <button type="submit">Войти</button>
                </form>
                <hr/>
                <small><Link to={'/register'}>Зарегистрировать тестировочный аккаунт</Link></small>
            </div>
        </div>
    </main>
    )
}