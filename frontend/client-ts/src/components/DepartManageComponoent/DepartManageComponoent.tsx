import React from "react";
import { PlusButton } from "../buttons/PlusButton/PlusButton";
import { Search } from "../common/Search/Search";
import st from './DepartManageComponoent.module.scss'

type PropsType = {
    placeholder?: string;
    title?: string;
    userName?: string 
    autorCreate?: string
    dateCreate?: string
}

export const DepartManage: React.FC<PropsType> = ({placeholder, title, userName, autorCreate, dateCreate}) => {
    return(
            <div id={st['wrapper-deportations']}>
                <div className={st["container"]}>
                    <div className={st["title"]}>{title}</div>
                    {/* button-add компонент */}
                    <PlusButton />
                </div>
                {/* search компонент */}
                <Search placehold={placeholder}/>
                <div id={st["list-deportations"]}>
                    <div className={st["card-deportation"]}>
                        <div className={st["title-deportation"]}>АиВТ</div>
                        <div className={st["number-groups"]}>40 групп</div>
                    </div>
                    <div className={st['card-deportation']}>
                        <div className={st['title-deportation']}>Не АиВТ</div>
                        <div className={st['number-groups']}>40000 групп</div>
                    </div>
                </div>
            </div>
    )
}