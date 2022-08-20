import "./UserManagement.module.scss"
import st from './UserManagement.module.scss'
import { CardUsers } from "../../components/cards/Cards"
import { DepartManage } from "../../components/common/List/List"

export const UserManagement = () =>{
    return (
        <div id={st["main"]}>
                <DepartManage title="Пользователи" placeholder='Пользователь'>
                    <div className={st["nav"]}>
                        АиВТ/261
                    </div>
                    <CardUsers userName='Данилов Вячеслав В.' autorCreate='Гордеева Ульяна М.' dateCreate='16.08.2022' />
                    <CardUsers userName='Данилов Вячеслав В.' autorCreate='Гордеева Ульяна М.' dateCreate='16.08.2022' />
                </DepartManage>
        </div>     
    )
}