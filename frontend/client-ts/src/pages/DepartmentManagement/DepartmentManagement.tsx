
import st from './DepartmentManagement.module.scss'
import { useCallback, useContext, useEffect, useState } from 'react'
import { Search } from '../../components/common/Search/Search'
import { getDepartamentsAPI, getGroupsAPI, getUsersAPI } from './API'
import { ListInfo } from '../../components/single/List/List'
import { Path } from '../../components/single/Path/Path'
import { Modal } from '../../components/common/Modal/Modal'
import { AddFolderForm } from '../../modals/AddFolderForm/AddFolderForm'
import { AddUserForm } from '../../modals/AddUserForm/AddUserForm'
import { AuthContext } from '../../context/authContext'


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
    const [openAddOrganization, setAddOrganization] = useState(false)
    const [openAddGroup, setAddGroup] = useState(false)
    const [openAddUser, setAddUser] = useState(false)

    const token = useContext(AuthContext).token

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
        getUsersAPI({groupId, setLoading, setUsersData, token})
        setShowGroup(false) 
        setShowUser(true)
        setPathGroup(groupTitle)
    }, [token])


    const getGroups = useCallback(async (departamentId: number, departamentTitle: string) => {
        getGroupsAPI({departamentId, setLoading, setUsersData, setGroupsData, token})
        setShowDepartament(false) 
        setShowGroup(true)
        setPathDepartament(departamentTitle)
    }, [token])


    const getDepartaments = useCallback(async () => {
        getDepartamentsAPI({setLoading, setDepartamentsData, token})
    }, [token])

    useEffect(() => {
        getDepartaments()
    }, [getDepartaments])

    console.log(departamentsData)

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
                        [{icon: "fa fa-plus", event: ()=>{setAddOrganization(true)}}]
                        :  showGroup ?
                        [{icon: "fa fa-plus", event: ()=>{setAddGroup(true)}}, 
                        {icon: "fa-solid fa-reply", event: ()=>{showForButtonBack()}}] 
                        : showUser ? 
                        [{icon: "fa fa-plus", event: ()=>{setAddUser(true)}}, 
                        {icon: "fa-solid fa-reply", event: ()=>{showForButtonBack()}}] 
                        : []
                    }
                >   
                    <Path 
                    firstAddress={pathDepartament} 
                    secondAddress={pathGroup}
                    showForPath={showForPath}
                    />
                    <div id={st['functions']}><Search placehold={"Поиск"} width={"300px"}/></div>
                </ListInfo>
                <div>
                    Тут будет инфа о пользователе 
                </div>
            </div>
            {openAddOrganization &&
                <Modal onClose={() => setAddOrganization(false)} closingBackground={true}>
                    <AddFolderForm onClose={() => setAddOrganization(false)} title={"Название отделения"}/>
                </Modal>}
            {openAddGroup &&
            <Modal onClose={() => setAddGroup(false)} closingBackground={true} >
                <AddFolderForm onClose={() => setAddGroup(false)} title={"Название группы"}/>
            </Modal>}

            {openAddUser &&
            <Modal onClose={() => setAddUser(false)} closingBackground={true}>
                <AddUserForm event={()=>{}}/>
            </Modal>}
        </div>
    )
}