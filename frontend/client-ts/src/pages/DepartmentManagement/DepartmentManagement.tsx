import { TwoCellsCard, UserCard} from '../../components/cards/Cards'
import st from './DepartmentManagement.module.scss'
import { FunctionalList } from '../../components/common/FunctionalList/FunctionalList'
import { useCallback, useEffect, useState } from 'react'
import { Loader } from '../../components/common/Loader/Loader'
import { Search } from '../../components/common/Search/Search'
import { getDepartamentsAPI, getGroupsAPI, getUsersAPI } from './API'


interface List {
    id: number
    title: string
    message: string
}

interface userData {
    id: number
    firstName: string
    dateCreate: string
    autorCreate: string
}

interface groupList extends List{
    departamentId: number
}

export const DepartmentManagement = () => {


    const [groupsData, setGroupsData] = useState<groupList[] | never[]>([])
    const [departamentsData, setDepartamentsData] = useState<List[] | never[]>([])
    const [usersData, setUsersData] = useState<userData[] | never[]>([])
    const [loading, setLoading] = useState(false)


    const getUsers = useCallback(async (groupId: number) => {
        getUsersAPI({groupId, setLoading, setUsersData})
    }, [])


    const getGroups = useCallback(async (departamentId: number) => {
        getGroupsAPI({departamentId, setLoading, setUsersData, setGroupsData})
    }, [])


    const getDepartaments = useCallback(async () => {
        getDepartamentsAPI({setLoading, setDepartamentsData})
    }, [])

    useEffect(() => {
        getDepartaments()
    }, [getDepartaments])



    return (
        <div id={st["main"]}>
            <div id={st['functions']}><Search placehold={"Поиск"} width={"300px"}/></div>
            <div id={st['control']}>
                <FunctionalList placeholder='Пользователь или группа' title='Отделение'>
                    {departamentsData.map(data => 
                        <div onClick={e => getGroups(data.id)}>
                            <TwoCellsCard 
                                title={data.title} 
                                message={data.message} 
                                key={data.id} 
                            />
                        </div>
                    )}
                </FunctionalList>
                <FunctionalList title='Группы'>
                {loading ? <Loader/> 
                        : groupsData.map(data => 
                            <div onClick={e => getUsers(data.id)}>
                                <TwoCellsCard title={data.title} key={data.id} message={data.message}/>
                            </div>)
                        }
                </FunctionalList>
                <FunctionalList title='Пользователи'>
                {loading ? <Loader/> 
                        : usersData.map(data => <UserCard userName={data.firstName} key={data.id} dateCreate={"2022.04.14"} autorCreate="Учитель"/>)
                        }
                </FunctionalList>
                <div>
                    Тут будет инфа о пользователе 
                </div>
            </div>
        </div>
    )
}