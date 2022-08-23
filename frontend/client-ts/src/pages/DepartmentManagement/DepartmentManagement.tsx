import { DeportationCard, InformationCard } from '../../components/cards/Cards'
import st from './DepartmentManagement.module.scss'
import { FunctionalList } from '../../components/common/FunctionalList/FunctionalList'
import { PlusButton } from '../../components/buttons/Buttons'
import axios from 'axios'
import { useCallback, useEffect, useState } from 'react'
import { Loader } from '../../components/common/Loader/Loader'


type GroupsList = {
    id: number
    title: string
    message: string
}

type DepartamentsList = {
    id: number
    title: string
    number: number
}

export const DepartmentManagement = () => {


    const [groupsData, setGroupsData] = useState<GroupsList[] | never[]>([])
    const [departamentsData, setDepartamentsData] = useState<DepartamentsList[] | never[]>([])
    const [loading, setLoading] = useState(false)


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
            <FunctionalList placeholder='Пользователь или группа' title='Отделение'>
                {departamentsData.map(data => 
                    <div onClick={e => getGroups(data.id)}>
                        <DeportationCard 
                            title={data.title} 
                            number={data.number} 
                            key={data.id} 
                        />
                    </div>
                )}
            </FunctionalList>
            <div id={st['wrapper-groups']}>
                <div className={st["container"]}>
                    <div className={st["title"]}>Группы</div>
                    <PlusButton />
                </div>
                <div id={st['list-groups']}>
                    {loading ? <Loader/> 
                    : groupsData.map(data => <InformationCard title={data.title} key={data.id} message={data.message} path={"#"}/>)
                    }
                </div>
            </div>
        </div>
    )
}