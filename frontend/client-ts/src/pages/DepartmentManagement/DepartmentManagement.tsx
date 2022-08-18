import { InformationCard } from '../../components/InformationCard/InformationCard'
import cs from './DepartmentManagement.module.css'
export const DepartmentManagement = () => {
    return (
        <div id={cs.main}>
            <div id={cs.wrapper_deportations}>
                <div className={cs.container}>
                    <div className={cs.title}>Отделения</div>
                    {/* button-add компонент */}
                    <button className={cs.button_add}><i className='fa fa-plus'></i></button>
                </div>
                {/* search компонент */}
                <div className={cs.search}><input type="seach" placeholder='Пользователь или группа'/> <i className={cs.fa +' '+ cs.fa_magnifying_glass}></i></div>
                <div id={cs.list_deportations}>
                    <div className={cs.card_deportation}>
                        <div className={cs.title_deportation}>АиВТ</div>
                        <div className={cs.number_groups}>40 групп</div>
                    </div>
                    <div className={cs.card_deportation}>
                        <div className={cs.title_deportation}>Не АиВТ</div>
                        <div className={cs.number_groups}>40000 групп</div>
                    </div>
                </div>
            </div>
            <div id={cs.wrapper_groups}>
            <div className={cs.container}>
                <div className={cs.title}>Группы</div>
                    {/* button-add компонент */}
                    <button className={cs.button_add}><i className='fa fa-plus'></i></button>
                </div>
                <div id={cs.list_groups}>
                    {/* card-group-teacher Всегда создан по умолчанию */}
                    <InformationCard title={'Учителя'} message={'20 Пользователей'}/>
                    <InformationCard title={'261'} message={'30 Пользователей'}/>
                    <InformationCard title={'269'} message={'28 Пользователей'}/>
                    <InformationCard title={'361'} message={'24 Пользователей'}/>
                </div>
            </div>
        </div>
    )
}