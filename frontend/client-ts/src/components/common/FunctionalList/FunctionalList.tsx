import React from "react";
import { PlusButton } from "../../buttons/Buttons";
import { Search } from "../Search/Search";
import st from './FunctionalList.module.scss'

type PropsType = {
    placeholder?: string;
    title?: string;
    children: React.ReactNode
}

export const FunctionalList: React.FC<PropsType> = ({placeholder, title, children }) => {
    return(
            <div id={st['wrapper-deportations']}>
                <div className={st["container"]}>
                    <div className={st["title"]}>{title}</div>
                    <PlusButton />
                </div>
                <Search placehold={placeholder}/>
                <div id={st["list-deportations"]}>
                    {children}
                </div>
            </div>
    )
}