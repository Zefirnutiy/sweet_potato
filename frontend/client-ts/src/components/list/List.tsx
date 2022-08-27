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
}


export const List: React.FC<PropsTypesList> = ({
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
    
