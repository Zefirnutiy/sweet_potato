import { FC } from "react";
import st from './AddUserForm.module.scss';

interface PropTypes {
    event: () => void
}


export const AddUserForm: FC<PropTypes> = ({event}) => {
    return(
        <div className={st["wrapper-add-user"]}>
            <div className={st["welcome"]}>
                <h3>Выберите возможности</h3>
                <div id={st["blanks"]}>Заготовки 
                    <select>
                        <option value="1">Нет</option>
                        <option value="2">Учитель</option>
                        <option value="3">Ученик</option>
                        <option value="4">Администратор</option>
                        <option value="5">Своё</option>
                    </select>
                </div>
                <div id={st["list-features"]}>
                    <label className={st["feature"]}><input type={"checkbox"}></input> Дышать</label>
                    <label className={st["feature"]}><input type={"checkbox"}></input> Дышать</label>
                    <label className={st["feature"]}><input type={"checkbox"}></input> Дышать</label>
                    <label className={st["feature"]}><input type={"checkbox"}></input> Дышать</label>
                    <label className={st["feature"]}><input type={"checkbox"}></input> Дышать</label>
                    <label className={st["feature"]}><input type={"checkbox"}></input> Дышать</label>
                </div>
            </div>
            <div className={st["registration"]}>
                <h3>Добавить пользователя</h3>
                <div className={st["registration-form"]}>
                    <div>Фамилия</div>
                    <input type="text"/>
                    <div>Имя</div>
                    <input type="text"/>
                    <div>Отчество</div>
                    <input type="text"/>
                    <div>Эл. почта</div>
                    <input type="text"/>
                    <button className={st["button-done"]} onClick={event}>Готово</button>
                </div>
            </div>
        </div>
    )
}