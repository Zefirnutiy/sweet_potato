import "./UserManagement.module.scss"
import st from './UserManagement.module.scss'
import { Search } from '../../components/common/Search/Search'
import { PlusButton } from "../../components/buttons/PlusButton/PlusButton"
import { CardUsers } from "../../components/cards/CardUsers/CardUsers"
import { GoBackButton } from "../../components/buttons/GoBackButton/GoBackButton"

export const UserManagement = () =>{
    return (
        <div id={st["main"]}>
            <div id={st["wrapper-users"]}>
                <div className={st["container"]}>
                    <div className={st["title"]}>Пользователи</div>
                    {/* button-add компонент */}
                    <PlusButton />
                    <GoBackButton path="/departmentManagement"/>
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