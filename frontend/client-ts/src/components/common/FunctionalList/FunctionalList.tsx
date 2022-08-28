import React from "react";
import { AddButton } from "../../buttons/Buttons";
import { Loader } from "../Loader/Loader";
import { Search } from "../Search/Search";
import st from './FunctionalList.module.scss'

interface PropsType {
    placeholder?: string;
    title?: string;
    children: React.ReactNode
    search?: boolean
    load?: boolean
}

export const FunctionalList: React.FC<PropsType> = ({placeholder, title, children, search, load}) => {
    return(
            <div id={st['wrapper']}>
                <div className={st["container"]}>
                    <div className={st["title"]}>{title}</div>
                    <AddButton/>
                    <AddButton/>
                </div>
                {search && <Search placehold={placeholder}/> }
                <div id={st["list"]}>
                    {load ? <Loader/> : children}
                </div>
            </div>
    )
}