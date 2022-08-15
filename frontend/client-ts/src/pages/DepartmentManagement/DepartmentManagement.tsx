import { InformationCard } from '../../components/InformationCard/InformationCard'
import './DepartmentManagement.css'
export const DepartmentManagement = () => {
    return (
        <div id="main">
            <div id="wrapper-deportations">
                <div className="container">
                    <div className='title'>Отделения</div>
                    {/* button-add компонент */}
                    <button className="button-add"><i className="fa fa-plus"></i></button>
                </div>
                {/* search компонент */}
                <div className="search"><input type="seach" placeholder='Пользователь или группа'/> <i className="fa fa-magnifying-glass"></i></div>
                <div id="list-deportations">
                    <div className="card-deportation">
                        <div className="title-deportation">АиВТ</div>
                        <div className="number-groups">40 групп</div>
                    </div>
                    <div className="card-deportation">
                        <div className="title-deportation">Не АиВТ</div>
                        <div className="number-groups">40000 групп</div>
                    </div>
                </div>
            </div>
            <div id="wrapper-groups">
            <div className="container">
                <div className='title'>Группы</div>
                    {/* button-add компонент */}
                    <button className="button-add"><i className="fa fa-plus"></i></button>
                </div>
                <div id="list-groups">
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