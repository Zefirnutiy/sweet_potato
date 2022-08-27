import { TwoCellsCard } from "../cards/Cards";
import { FunctionalList } from "../common/FunctionalList/FunctionalList";



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
        { showDepartament ? <FunctionalList placeholder='Пользователь или группа' title='Отделение' load={loading}>
                    {children}
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
                    {children}
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
                    {children}
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
    
