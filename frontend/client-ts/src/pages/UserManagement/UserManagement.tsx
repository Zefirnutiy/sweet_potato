<<<<<<< HEAD
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

export const UserManagement = () =>{
    return (
        <div id={st.main}>
            <div id={st.wrapper_users}>
                <div className={st.container}>
                    <div className={st.title}>Пользователи</div>
                    {/* button-add компонент */}
                    <button className={st.button_add}><i className="fa fa-plus"></i></button>
                </div>
                <div className={st.nav}>
                    <button className={st.button_go_back}><i className="fa-solid fa-circle-arrow-left"></i></button>
                    АиВТ/261
                </div>
                {/* search компонент */}
                <div className={st.search}><input type="seach" placeholder='Пользователь'/> <i className="fa fa-magnifying-glass"></i></div>
                <div id={st.list_users}>
                    <div className={st.card_user}>
                        <div className={st.user_name}>Данилов Вячеслав В.</div>
                        <div className={st.created}>Создал: Гордеева Ульяна М. <br/> Создан: 16.08.2022</div>
                    </div>
                    <div className={st.card_user}>
                        <div className={st.user_name}>Данилов Вячеслав В.</div>
                        <div className={st.created}>Создал: Гордеева Ульяна М. <br/> Создан: 16.08.2022</div>
                    </div>
                </div>
            </div>
            <div id={st.wrapper_user_controller}>
            </div>
        </div>     
    )
}