
import st from './DepartmentManagement.module.scss'
import { useCallback, useEffect, useState } from 'react'
import { Search } from '../../components/common/Search/Search'
import { getDepartamentsAPI, getGroupsAPI, getUsersAPI } from './API'
import { ListInfo } from '../../components/single/List/List'
import { Path } from '../../components/single/Path/Path'

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
    const [pathDepartament, setPathDepartament] = useState("")
    const [pathGroup, setPathGroup] = useState("")

    const showForPath = (isDepartament: boolean) => {
        setShowUser(false)
        setPathGroup("")
        
        if(isDepartament){
            setPathDepartament("")
            setShowGroup(false)
            setShowDepartament(true)
            return
        }
        setShowGroup(true)    
    }

    const showForButtonBack = () => {
        if(showUser){
            setShowUser(false)
            setShowGroup(true)
            setPathGroup("")
            return
        }

        if(showGroup){
            setShowGroup(false)
            setShowDepartament(true)
            setPathDepartament("")
            return
        }

    }

    const getUsers = useCallback(async (groupId: number, groupTitle: string) => {
        getUsersAPI({groupId, setLoading, setUsersData})
        setShowGroup(false) 
        setShowUser(true)
        setPathGroup(groupTitle)
    }, [])


    const getGroups = useCallback(async (departamentId: number, departamentTitle: string) => {
        getGroupsAPI({departamentId, setLoading, setUsersData, setGroupsData})
        setShowDepartament(false) 
        setShowGroup(true)
        setPathDepartament(departamentTitle)
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
                    buttons={
                        showDepartament ? 
                        [{icon: "fa fa-plus", event: ()=>{alert("Нажата кнопка 'добавить'")}}]
                        :
                        [{icon: "fa fa-plus", event: ()=>{alert("Нажата кнопка 'добавить'")}}, 
                        {icon: "fa-solid fa-reply", event: ()=>{showForButtonBack()}}]
                    }
                >   
                    <Path 
                    pathDepartament={pathDepartament} 
                    pathGroup={pathGroup}
                    showForPath={showForPath}
                    />
                    <div id={st['functions']}><Search placehold={"Поиск"} width={"300px"}/></div>
                </ListInfo>
                <div>
                    Тут будет инфа о пользователе 
                </div>
            </div>
        </div>
    )
}