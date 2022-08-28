import { TwoCellsCard, UserCard } from "../../cards/Cards";
import { FunctionalList } from "../../common/FunctionalList/FunctionalList";



interface PropsTypesList{
    showDepartament: boolean
    departamentsData: any[]
    getGroups: (departamentId: number, departamentTitle: string) => Promise<void>
    showGroup: boolean
    groupsData: any[]
    getUsers: (groupId: number, groupTitle: string) => Promise<void>
    showUser: boolean
    usersData: any[]
    loading: boolean
    children: React.ReactNode
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
    loading,
    children
}) => (
    <>
        { showDepartament ? <FunctionalList placeholder='Пользователь или группа' title='Отделения' load={loading}>
                    {children}
                      { departamentsData.map(data => 
                        <div onClick={e => getGroups(data.id, data.title)}>
                            <TwoCellsCard 
                                title={data.title} 
                                message={data.message} 
                                key={data.id} 
                            />
                        </div>
                    )}
                </FunctionalList> : null}

                { showGroup ? <FunctionalList placeholder='Пользователь или группа' title='Группы' load={loading}>
                    {children}
                      { groupsData.map(data => 
                        <div onClick={e => getUsers(data.id, data.title)}>
                            <TwoCellsCard 
                                title={data.title} 
                                message={data.message} 
                                key={data.id} 
                            />
                        </div>
                    )}
                </FunctionalList> : null}

                {showUser ? <FunctionalList placeholder='Пользователь или группа' title='Пользователи' load={loading}>
                    {children}
                      { usersData.map(data => 
                        <div>
                            <UserCard 
                                userName={data.firstName} 
                                dateCreate={data.message} 
                                autorCreate={data.autorCreate}
                                key={data.id} 
                            />
                        </div>
                    )}
                </FunctionalList> : null}
    </>
)
    
