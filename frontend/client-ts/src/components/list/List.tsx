import { TwoCellsCard } from "../cards/Cards";
import { FunctionalList } from "../common/FunctionalList/FunctionalList";
import { Search } from "../common/Search/Search";


interface PropsTypesList{
    showDepartament: boolean
    departamentsData: any[]
    getGroups: (departamentId: number) => Promise<void>
    showGroup: boolean
    groupsData: any[]
    getUsers: (groupId: number) => Promise<void>
    showUser: boolean
    usersData: any[]
    loading: boolean
}


export const ListInfo: React.FC<PropsTypesList> = ({
    showDepartament,
    departamentsData,
    getGroups,
    showGroup,
    groupsData,
    getUsers,
    showUser,
    usersData,
    loading
}) => (
    <>
        { showDepartament ? <FunctionalList placeholder='Пользователь или группа' title='Отделение' load={loading}>
                    <div><Search placehold={"Поиск"} width={"160px"}/></div>
                    
                      { departamentsData.map(data => 
                        <div onClick={e => getGroups(data.id)}>
                            <TwoCellsCard 
                                title={data.title} 
                                message={data.message} 
                                key={data.id} 
                            />
                        </div>
                    )}
                </FunctionalList> : null}

                { showGroup ? <FunctionalList placeholder='Пользователь или группа' title='Группы' load={loading}>
                <div><Search placehold={"Поиск"} width={"160px"}/></div>
                      { groupsData.map(data => 
                        <div onClick={e => getUsers(data.id)}>
                            <TwoCellsCard 
                                title={data.title} 
                                message={data.message} 
                                key={data.id} 
                            />
                        </div>
                    )}
                </FunctionalList> : null}

                {showUser ? <FunctionalList placeholder='Пользователь или группа' title='Пользователи' load={loading}>
                <div><Search placehold={"Поиск"} width={"160px"}/></div>
                      { usersData.map(data => 
                        <div>
                            <TwoCellsCard 
                                title={data.firstName} 
                                message={data.dateCreate} 
                                key={data.id} 
                            />
                        </div>
                    )}
                </FunctionalList> : null}
    </>
)
    
