import { TwoCellsCard, UserCard} from '../../components/cards/Cards'
import st from './DepartmentManagement.module.scss'
import { FunctionalList } from '../../components/common/FunctionalList/FunctionalList'
import axios from 'axios'
import { useCallback, useEffect, useState } from 'react'
import { Loader } from '../../components/common/Loader/Loader'
import { Search } from '../../components/common/Search/Search'


type List = {
    id: number
    title: string
    message: string
}

type userData = {
    id: number
    firstName: string
    dateCreate: string
    autorCreate: string


}
export const DepartmentManagement = () => {


    const [groupsData, setGroupsData] = useState<List[] | never[]>([])
    const [departamentsData, setDepartamentsData] = useState<List[] | never[]>([])
    const [usersData, setUsersData] = useState<userData[] | never[]>([])
    const [loading, setLoading] = useState(false)

    const getUsers = useCallback(async (GroupId: number) => {
        setLoading(true)
        await axios.get(
            `/api/users/${GroupId}`, 
            ).then((response) => {
                setLoading(false)
                console.log(response.data.users)
                setUsersData(response.data.users)
            })  
            .catch(e => console.log(e))
    }, [])

    const getGroups = useCallback(async (idDepartament: number) => {
        
        setLoading(true)
        await axios.get(
            `/api/groups/${idDepartament}`, 
            ).then((response) => {
                setLoading(false)
                console.log(response.data.groups)
                setGroupsData(response.data.groups)
            })  
            .catch(e => console.log(e))
    }, [])

    const getDepartaments = useCallback(async () => {
        
        setLoading(true)
        await axios.get(
            `/api/depataments/`, 
            ).then((response) => {
                setLoading(false)
                setDepartamentsData(response.data.groups)
            })  
            .catch(e => console.log(e))
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