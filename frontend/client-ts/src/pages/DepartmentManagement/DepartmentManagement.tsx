import { InformationCard } from '../../components/InformationCard/InformationCard'
import st from './DepartmentManagement.module.scss'

export const DepartmentManagement = () => {
    return (
        <div id={st["main"]}>
            <div id={st['wrapper-deportations']}>
                <div className={st["container"]}>
                    <div className={st["title"]}>Отделения</div>
                    {/* button-add компонент */}
                    <button className={st["button-add"]}><i className='fa fa-plus'></i></button>
                </div>
                {/* search компонент */}
                <div className={st["search"]}><input type="seach" placeholder='Пользователь или группа'/> <i className='fa fa-magnifying-glass'></i></div>
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
                    <button className={st["button-add"]}><i className='fa fa-plus'></i></button>
                </div>
                <div id={st['list-groups']}>
                    {/* card-group-teacher Всегда создан по умолчанию */}
                    <InformationCard title={'Учителя'} message={'20 Пользователей'} event={() =>{window.location.replace("/userManagement")}}/>
                    <InformationCard title={'261'} message={'30 Пользователей'} event={() =>{window.location.replace("/userManagement")}}/>
                    <InformationCard title={'269'} message={'28 Пользователей'} event={() =>{window.location.replace("/userManagement")}}/>
                    <InformationCard title={'361'} message={'24 Пользователей'} event={() =>{window.location.replace("/userManagement")}}/>
                </div>
            </div>
        </div>
    )
}