/* eslint-disable jsx-a11y/anchor-is-valid */
import { Link } from "react-router-dom"
import "../styles/css/registration.css"

export const RegisterPage = () => {
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
            <div className="registration">
                <h3>Регистрация</h3>
                <form action="/" className="registrationForm">
                    <span>Название организации</span>
                    <input type="text"/>
                    <span>Email</span>
                    <input type="text"/> 
                    <span>Пароль</span>
                    <input type="text"/>
                    <span>Повторите пароль</span>
                    <input type="text"/>
                    <div className="contract">
                        <p><input type="checkbox"/></p>
                        <p>
                        <small>
                            Отправляя сведения через электронную форму, вы соглашаетесь с условиями <a href="/">Лицензионного договора</a>, 
                            даете согласие на обработку персональных данных, их сбор, хранение и передачу третьим лицам предоставленной вами 
                            информации на условиях <a href="/">Политики обработки персональных данных</a>.
                        </small>
                        </p>
                    </div>
                    <button type="submit">Зарегистрировать аккаунт</button>
                </form>
                <hr/>
                <small><Link to={'/login'}>Войти в аккаунт</Link></small>
            </div>
        </div>
    </main>
    )
}