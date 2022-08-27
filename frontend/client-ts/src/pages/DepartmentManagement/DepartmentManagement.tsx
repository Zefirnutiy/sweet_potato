
import st from './DepartmentManagement.module.scss'
import { useCallback, useEffect, useState } from 'react'
import { Search } from '../../components/common/Search/Search'
import { getDepartamentsAPI, getGroupsAPI, getUsersAPI } from './API'
import { ListInfo } from '../../components/list/List'


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
    const [showDepartament, setShowDepartament] = useState(true)
    const [showGroup, setShowGroup] = useState(false)
    const [showUser, setShowUser] = useState(false)



    const getUsers = useCallback(async (groupId: number) => {
        getUsersAPI({groupId, setLoading, setUsersData})
        setShowGroup(false) 
        setShowUser(true)
    }, [])


    const getGroups = useCallback(async (departamentId: number) => {
        getGroupsAPI({departamentId, setLoading, setUsersData, setGroupsData})
        setShowDepartament(false) 
        setShowGroup(true)
    }, [])


    const getDepartaments = useCallback(async () => {
        getDepartamentsAPI({setLoading, setDepartamentsData})
    }, [])

    useEffect(() => {
        getDepartaments()
    }, [getDepartaments])



    return (
        <div id={st["main"]}>
            <div id={st['control']}>
                <ListInfo 
                    showDepartament={showDepartament}
                    departamentsData={departamentsData}
                    getGroups={getGroups}
                    showGroup={showGroup}
                    groupsData={groupsData}
                    getUsers={getUsers}
                    showUser={showUser}
                    usersData={usersData}
                    loading={loading}
                />
                <div>
                    Тут будет инфа о пользователе 
                </div>
            </div>
        </div>
    )
}