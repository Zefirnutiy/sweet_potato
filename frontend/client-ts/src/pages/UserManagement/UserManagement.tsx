import st from "./UserManagement.module.css"

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