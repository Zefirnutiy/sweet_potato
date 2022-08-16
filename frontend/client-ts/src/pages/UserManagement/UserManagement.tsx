import "./UserManagement.css"

export const UserManagement = () =>{
    return (
        <div id='main'>
            <div id="wrapper-users">
                <div className="container">
                    <div className='title'>Пользователи</div>
                    {/* button-add компонент */}
                    <button className="button-add"><i className="fa fa-plus"></i></button>
                </div>
                <div className="nav">
                    <button className="button-go-back"><i className="fa-solid fa-circle-arrow-left"></i></button>
                    АиВТ/261
                </div>
                {/* search компонент */}
                <div className="search"><input type="seach" placeholder='Пользователь'/> <i className="fa fa-magnifying-glass"></i></div>
                <div id="list-users">
                    <div className="card-user">
                        <div className="user-name">Данилов Вячеслав В.</div>
                        <div className="created">Создал: Гордеева Ульяна М. <br/> Создан: 16.08.2022</div>
                    </div>
                    <div className="card-user">
                        <div className="user-name">Данилов Вячеслав В.</div>
                        <div className="created">Создал: Гордеева Ульяна М. <br/> Создан: 16.08.2022</div>
                    </div>
                </div>
            </div>
            <div id="wrapper-user-controller">
                
            </div>
        </div>     
    )
}