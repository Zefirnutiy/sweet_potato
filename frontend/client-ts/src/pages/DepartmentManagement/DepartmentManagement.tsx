import { InformationCard } from '../../components/cards/InformationCard/InformationCard'
import st from './DepartmentManagement.module.scss'
import { Search } from '../../components/common/Search/Search'
import { PlusButton } from '../../components/buttons/PlusButton/PlusButton'

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
            <div id={st['wrapper-deportations']}>
                <div className={st["container"]}>
                    <div className={st["title"]}>Отделения</div>
                    {/* button-add компонент */}
                    <PlusButton />
                </div>
                {/* search компонент */}
                <Search placehold='Пользователь или группа'/>
                <div id={st["list-deportations"]}>
                    <div className={st["card-deportation"]}>
                        <div className={st["title-deportation"]}>АиВТ</div>
                        <div className={st["number-groups"]}>40 групп</div>
                    </div>
                    <div className={st['card-deportation']}>
                        <div className={st['title-deportation']}>Не АиВТ</div>
                        <div className={st['number-groups']}>40000 групп</div>
                    </div>
                </div>
            </div>
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