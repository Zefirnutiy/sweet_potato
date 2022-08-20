import { DeportationCard, InformationCard } from '../../components/cards/Cards'
import st from './DepartmentManagement.module.scss'
import { DepartManage } from '../../components/common/List/List'
import { PlusButton } from '../../components/buttons/Buttons'
import axios from 'axios'
import { useCallback, useEffect, useState } from 'react'
import { Loader } from '../../components/common/Loader/Loader'


type GroupsList = {
    id: number
    title: string
    message: string
}

export const DepartmentManagement = () => {


    const [groupsData, setGroupsData] = useState<GroupsList[] | never[]>([])
    const [loading, setLoading] = useState(false)

    const getGroups = useCallback(async () => {
        
        setLoading(true)
        await axios.get(
            `/api/groups/`, 
            ).then((response) => {
                setLoading(false)
                setGroupsData(response.data.groups)
            })  
            .catch(e => console.log(e))
    }, [])

    useEffect(() => {
        getGroups()
    }, [getGroups])



    return (
        <div id={st["main"]}>
            <DepartManage placeholder='Пользователь или группа' title='Отделение'>
                <DeportationCard title='АИВТ' number={12}/>
            </DepartManage>
            <div id={st['wrapper-groups']}>
            <div className={st["container"]}>
                <div className={st["title"]}>Группы</div>
                    {/* button-add компонент */}
                    <PlusButton />
                </div>
                <div id={st['list-groups']}>

                    {/* card-group-teacher Всегда создан по умолчанию */}
                    {loading ? <Loader/> : groupsData.map(data => 
                        <InformationCard title={data.title} message={data.message} path={"#"}/>
                    )}
                </div>
            </div>
        </div>
    )
}

