import { DeportationCard, InformationCard } from '../../components/cards/Cards'
import st from './DepartmentManagement.module.scss'
import { DepartManage } from '../../components/common/List/List'
import { PlusButton } from '../../components/buttons/Buttons'

const testData = [
    {
        "title": "Учителя",
        "message": "30 Пользователей",
        "path": "/users"
    },
    {
        "title": "261",
        "message": "20 Пользователей",
        "path": "/users"
    },
    {
        "title": "369",
        "message": "33 Пользователя",
        "path": "/users"
    },
    {
        "title": "362",
        "message": "6 Пользователей",
        "path": "/users"
    }
]

export const DepartmentManagement = () => {
    return (
        <div id={st["main"]}>
            <DepartManage placeholder='Пользователь или группа' title='Отделение'>
                <DeportationCard title='АИВТ' number={12}/>
            </DepartManage>
            <div id={st['wrapper-groups']}>
            <div className={st["container"]}>
                <div className={st["title"]}>Группы</div>
                    {/* button-add компонент */}
                    <PlusButton />
                </div>
                <div id={st['list-groups']}>

                    {/* card-group-teacher Всегда создан по умолчанию */}
                    {testData.map(data => 
                        <InformationCard title={data.title} message={data.message} path={data.path}/>
                    )}
                </div>
            </div>
        </div>
    )
}