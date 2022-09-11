import axios from "axios"

interface IGetRequest{
    setLoading: (value: React.SetStateAction<boolean>) => void
    token: string | null
}

interface IGetUsers extends IGetRequest{
    groupId: number
    setUsersData: (value: React.SetStateAction<never[] | any[]>) => void
}

interface IGetGroups extends IGetRequest{
    departamentId: number
    setUsersData: (value: React.SetStateAction<never[] | any[]>) => void
    setGroupsData: (value: React.SetStateAction<never[] | any[]>) => void
}

interface IGetDepartament extends IGetRequest{
    setDepartamentsData: (value: React.SetStateAction<never[] | any[]>) => void
}



export const getUsersAPI = async ({groupId, setLoading, setUsersData, token}: IGetUsers) => {
    setLoading(true)
    await axios.get(
        `/api/users/${groupId}`, 
        {headers: {'Authorization': `Bearer ${token}`}},
        ).then((response) => {
            setLoading(false)
            console.log(response.data.users)
            setUsersData(response.data.users)
        })  
        .catch(e => console.log(e))
}

export const getGroupsAPI = async ({departamentId, setLoading, setGroupsData, setUsersData, token}: IGetGroups) => {
    setUsersData([])
    setLoading(true)
    await axios.get(
        `/api/groups/${departamentId}`, 
        {headers: {'Authorization': `Bearer ${token}`}},
        ).then((response) => {
            setLoading(false)
            console.log(response.data.groups)
            setGroupsData(response.data.groups)
        })  
        .catch(e => console.log(e))
}

export const getDepartamentsAPI = async ({setLoading, setDepartamentsData, token} : IGetDepartament) => {
    setLoading(true)
    await axios.get(
        `http://localhost:8080/api/department/getDepartments`, 
        {headers: {'Authorization': `Bearer ${token}`}},
        ).then((response) => {
            setLoading(false)
            setDepartamentsData(response.data.groups)
        })  
        .catch(e => console.log(e))
}