import "./UserManagement.module.scss"
import st from './UserManagement.module.scss'

export const UserManagement = () =>{
    return (
        <div id={st["main"]}>
            <div id={st["wrapper-users"]}>
                <div className={st["container"]}>
                    <div className={st["title"]}>Пользователи</div>
                    {/* button-add компонент */}
                    <button className={st["button-add"]}><i className="fa fa-plus"></i></button>
                    <button className={st["button-go-back"]}><i className="fa-solid fa-circle-arrow-left"></i></button>
                </div>
                {/* search компонент */}
                <div className={st["search"]}><input type="seach" placeholder='Пользователь'/> <i className="fa fa-magnifying-glass"></i></div>
                <div className={st["nav"]}>
                    АиВТ/261
                </div>
                <div id={st["list-users"]}>
                    <div className={st["card-user"]}>
                        <div className={st["user-name"]}>Данилов Вячеслав В.</div>
                        <div className={st["created"]}>Создал: Гордеева Ульяна М. <br/> Создан: 16.08.2022</div>
                    </div>
                    <div className={st["card-user"]}>
                        <div className={st["user-name"]}>Данилов Вячеслав В.</div>
                        <div className={st["created"]}>Создал: Гордеева Ульяна М. <br/> Создан: 16.08.2022</div>
                    </div>
                </div>
            </div>
            <div id={st["wrapper-user-controller"]}>
            </div>
        </div>     
    )
}