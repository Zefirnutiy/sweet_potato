import "./UserManagement.module.scss"
import st from './UserManagement.module.scss'
import { UserCard } from "../../components/cards/Cards"
import { FunctionalList } from "../../components/common/FunctionalList/FunctionalList"

export const UserManagement = () =>{
    return (
        <div id={st["main"]}>
                <FunctionalList title="Пользователи" placeholder='Пользователь'>
                    <div className={st["nav"]}>
                        АиВТ/261
                    </div>
                    <UserCard userName='Данилов Вячеслав В.' autorCreate='Гордеева Ульяна М.' dateCreate='16.08.2022' />
                    <UserCard userName='Данилов Вячеслав В.' autorCreate='Гордеева Ульяна М.' dateCreate='16.08.2022' />
                </FunctionalList>
        </div>     
    )
}