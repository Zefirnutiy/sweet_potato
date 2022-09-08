import { FC } from "react";
import st from './UserAccountInformation.module.scss'

interface PropTypes{
    
}
export const UserAccountInformation: FC<PropTypes> = () => (
    <div className={st["window-user-account-information"]}>
        <div className={st["user-account-avatar"]}><img src="#" alt="Аватар" /></div>
        <div className={st["user-account-information"]}>
            <div className={st["user-account-name"]}>
                Папов Иван Васильевич
            </div>
            <div className={st["created"]}>Создал: Учитель(имя) <br/> Создан: 2022.04.4</div>
            <div className={st["contacts"]}>Телефон: 89091795631 <br/> Почта: popow2000@mail.ru</div>
        </div>
        <div className={st["user-account-permissions"]}>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["permission"]}>
                <input type="checkbox" /> 
                <label htmlFor="#">Дышать</label>
            </div>
            <div className={st["control-buttons"]}>
                <button>Редактировать</button>
                <button>Удалить</button>
            </div>
        </div>
    </div>
)