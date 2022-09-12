import { TwoCellsCard, UserCard } from "../../cards/Cards";
import { FunctionalList } from "../../common/FunctionalList/FunctionalList";

interface IButton {
    icon: string
    event: () => void
}

interface PropsTypesList{
    buttons: IButton[]
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
    buttons,
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
        { showDepartament ? <FunctionalList buttons={buttons} placeholder='Пользователь или группа' title='Отделения' load={loading}>
                    {children}
                      { departamentsData.map(data => 
                        <div onClick={e => getGroups(data.Id, data.Title)}>
                            <TwoCellsCard 
                                title={data.Title} 
                                message={"Всего групп " + data.GroupCount} 
                                key={data.Id} 
                            />
                        </div>
                    )}
                </FunctionalList> : null}

                { showGroup ? <FunctionalList buttons={buttons} placeholder='Пользователь или группа' title='Группы' load={loading}>
                    {children}
                      { groupsData.map(data => 
                        <div onClick={e => getUsers(data.Id, data.Title)}>
                            <TwoCellsCard 
                                title={data.Title} 
                                message={data.TitleSingular} 
                                key={data.Id} 
                            />
                        </div>
                    )}
                </FunctionalList> : null}

                {showUser ? <FunctionalList buttons={buttons} placeholder='Пользователь или группа' title='Пользователи' load={loading}>
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
    
