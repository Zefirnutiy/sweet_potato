import { InformationCard } from '../../components/InformationCard/InformationCard'
import styles from './DepartmentManagement.module.scss'
export const DepartmentManagement = () => {
    return (
        <div id={styles.main}>
            <div id={styles['wrapper-deportations']}>
                <div className={styles.container}>
                    <div className={styles.title}>Отделения</div>
                    {/* button-add компонент */}
<<<<<<< HEAD
                </div>
                {/* search компонент */}
                <div className={styles.search}><input type="seach" placeholder='Пользователь или группа'/> <i className={styles.fa +' '+ styles['fa-magnifying-glass']}></i></div>
                <div id={styles['list-deportations']}>
                    <div className={styles['card-deportation']}>
                        <div className={styles['title-deportation']}>АиВТ</div>
                        <div className={styles['number-groups']}>40 групп</div>
=======
                    <button className={cs.button_add}><i className='fa fa-plus'></i></button>
                </div>
                {/* search компонент */}
                <div className={cs.search}><input type="seach" placeholder='Пользователь или группа'/> <i className='fa fa-magnifying-glass'></i></div>
                <div id={cs.list_deportations}>
                    <div className={cs.card_deportation}>
                        <div className={cs.title_deportation}>АиВТ</div>
                        <div className={cs.number_groups}>40 групп</div>
>>>>>>> 9059103f366f48ad18d61169c9c7ef12c9b63b01
                    </div>
                    <div className={styles['card-deportation']}>
                        <div className={styles['title-deportation']}>Не АиВТ</div>
                        <div className={styles['number-groups']}>40000 групп</div>
                    </div>
                </div>
            </div>
            <div id={styles['wrapper-groups']}>
            <div className={styles.container}>
                <div className={styles.title}>Группы</div>
                    {/* button-add компонент */}
<<<<<<< HEAD
                    <button className={styles['button-add']}><i className={styles.fa + ' ' + styles.fa_plus}></i></button>
=======
                    <button className={cs.button_add}><i className='fa fa-plus'></i></button>
>>>>>>> 9059103f366f48ad18d61169c9c7ef12c9b63b01
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