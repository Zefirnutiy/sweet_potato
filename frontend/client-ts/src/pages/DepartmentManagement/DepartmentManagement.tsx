import { InformationCard } from '../../components/cards/InformationCard/InformationCard'
import st from './DepartmentManagement.module.scss'
import { Search } from '../../components/common/Search/Search'
import { PlusButton } from '../../components/buttons/PlusButton/PlusButton'
import { DepartManage } from '../../components/DepartManageComponoent/DepartManageComponoent'

const testData = [
    {
        "title": "Учителя",
        "message": "30 Пользователей",
        "path": "/userManagement"
    },
    {
        "title": "261",
        "message": "20 Пользователей",
        "path": "/userManagement"
    },
    {
        "title": "369",
        "message": "33 Пользователя",
        "path": "/userManagement"
    },
    {
        "title": "362",
        "message": "6 Пользователей",
        "path": "/userManagement"
    }
]

export const DepartmentManagement = () => {
    return (
        <div id={st["main"]}>
            <DepartManage placeholder='Пользователь или группа' title='Отделение'/>
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