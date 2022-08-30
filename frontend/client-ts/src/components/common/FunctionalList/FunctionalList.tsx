import React from "react";
import { SquareButton } from "../../buttons/Buttons";
import { Loader } from "../Loader/Loader";
import { Search } from "../Search/Search";
import st from './FunctionalList.module.scss'

interface IButton {
    icon: string
    event: () => void
}

interface PropsType {
    placeholder?: string;
    title?: string;
    children: React.ReactNode
    search?: boolean
    load?: boolean
    buttons?: IButton[]
}

export const FunctionalList: React.FC<PropsType> = ({placeholder, title, children, buttons, search, load}) => {
    return(
            <div id={st['wrapper']}>
                <div className={st["container"]}>
                    <div className={st["title"]}>{title}</div>
                    <div className={st["buttons"]}>
                        {buttons?.map(button => <SquareButton icon={button.icon} event={button.event}/> )}
                    </div>
                </div>
                {search && <Search placehold={placeholder}/> }
                <div id={st["list"]}>
                    {load ? <Loader/> : children}
                </div>
            </div>
    )
}