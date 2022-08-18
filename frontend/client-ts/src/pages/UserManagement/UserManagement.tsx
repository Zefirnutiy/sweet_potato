import "./UserManagement.module.scss"
import styles from './UserManagement.module.scss'

export const UserManagement = () =>{
    return (
        <div id={styles.main}>
            <div id={styles["wrapper-users"]}>
                <div className={styles.container}>
                    <div className={styles.title}>Пользователи</div>
                    {/* button-add компонент */}
                    <button className={styles["button-add"]}><i className={styles.fa +' '+ styles['fa-plus']}></i></button>
                </div>
                <div className={styles.nav}>
                    <button className={styles["button-go-back"]}><i className={styles['fa-solid'] +' '+ styles['fa-circle-arrow-left']}></i></button>
                    АиВТ/261
                </div>
                {/* search компонент */}
                <div className={styles.search}><input type="seach" placeholder='Пользователь'/> <i className={styles.fa +' '+ styles['fa-magnifying-glass']}></i></div>
                <div id={styles["list-users"]}>
                    <div className={styles["card-user"]}>
                        <div className={styles["user-name"]}>Данилов Вячеслав В.</div>
                        <div className={styles.created}>Создал: Гордеева Ульяна М. <br/> Создан: 16.08.2022</div>
                    </div>
                    <div className={styles["card-user"]}>
                        <div className={styles["user-name"]}>Данилов Вячеслав В.</div>
                        <div className={styles.created}>Создал: Гордеева Ульяна М. <br/> Создан: 16.08.2022</div>
                    </div>
                </div>
            </div>
            <div id={styles["wrapper-user-controller"]}>
                
            </div>
        </div>     
    )
}