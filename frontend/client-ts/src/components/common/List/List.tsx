import React from "react";
import { PlusButton } from "../../buttons/Buttons";
import { Search } from "../Search/Search";
import st from './List.module.scss'

type PropsType = {
    placeholder?: string;
    title?: string;
    children: React.ReactNode
}

export const DepartManage: React.FC<PropsType> = ({placeholder, title, children }) => {
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
                    {children}
                </div>
            </div>
    )
}