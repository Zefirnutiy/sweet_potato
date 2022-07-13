/* eslint-disable jsx-a11y/anchor-is-valid */
import { Link } from "react-router-dom"
import "../styles/css/registration.css"

export const RegisterPage = () => {
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
            <div class="registration">
                <h3>Регистрация</h3>
                <form action="#" class="registrationForm">
                    <span>Название организации</span>
                    <input type="text"/>
                    <span>Email</span>
                    <input type="text"/> 
                    <span>Пароль</span>
                    <input type="text"/>
                    <span>Повторите пароль</span>
                    <input type="text"/>
                    <div class="contract">
                        <p><input type="checkbox"/></p>
                        <p>
                        <small>
                            Отправляя сведения через электронную форму, вы соглашаетесь с условиями <a href="#">Лицензионного договора</a>, 
                            даете согласие на обработку персональных данных, их сбор, хранение и передачу третьим лицам предоставленной вами 
                            информации на условиях <a href="#">Политики обработки персональных данных</a>.
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