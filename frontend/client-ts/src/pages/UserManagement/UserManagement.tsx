import "./UserManagement.module.scss"
import st from './UserManagement.module.scss'
import { Search } from '../../components/Search/Search'
import { PlusButton } from "../../components/StupidButton/PlusButton"
import { CardUsers } from "../../components/CardUsers/CardUsers"

export const UserManagement = () =>{
    return (
        <div id={st["main"]}>
            <div id={st["wrapper-users"]}>
                <div className={st["container"]}>
                    <div className={st["title"]}>Пользователи</div>
                    {/* button-add компонент */}
                    <PlusButton />
                    <button className={st["button-go-back"]}><i className="fa-solid fa-circle-arrow-left"></i></button>
                </div>
                {/* search компонент */}
                <Search placehold='Пользователь'/>
                <div className={st["nav"]}>
                    АиВТ/261
                </div>
                <div id={st["list-users"]}>
                <CardUsers userName='Данилов Вячеслав В.' autorCreate='Гордеева Ульяна М.' dateCreate='16.08.2022' />
                <CardUsers userName='Данилов Вячеслав В.' autorCreate='Гордеева Ульяна М.' dateCreate='16.08.2022' />
                </div>
            </div>
            <div id={st["wrapper-user-controller"]}>
            </div>
        </div>     
    )
}