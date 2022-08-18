import { InformationCard } from '../../components/InformationCard/InformationCard'
import { ListDeportations } from '../../components/ListDeport/ListDeportations'
import { Search } from '../../components/Search/Search'
import styles from './DepartmentManagement.module.css'
export const DepartmentManagement = () => {
    return (
        <div id={styles.main}>
            <div id={styles['wrapper-deportations']}>
                <div className={styles.container}>
                    <div className={styles.title}>Отделения</div>
                    {/* button-add компонент */}
                    <button className={styles['button-add']}><i className={styles.fa + ' ' + styles.fa_plus}></i></button>
                </div>
                {/* search компонент */}
                <Search placehold='Пользователь или группа' />
                <ListDeportations Department='АиВТ' AmountGroup='40 Групп' />
                <ListDeportations Department='Не АиВТ' AmountGroup='40000 Групп' />
            </div>
            <div id={styles['wrapper-groups']}>
            <div className={styles.container}>
                <div className={styles.title}>Группы</div>
                    {/* button-add компонент */}
                    <button className={styles['button-add']}><i className={styles.fa + ' ' + styles.fa_plus}></i></button>
                </div>
                <div id={styles['list-groups']}>
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