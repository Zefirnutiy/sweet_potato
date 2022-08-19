/* eslint-disable @typescript-eslint/no-unused-expressions */
import React from "react";
import st from './Search.module.scss';

type PropsType = {
    placehold: any
    event?: () => void

}

export const Search: React.FC<PropsType> = ({placehold, event = () => {return}}) => {
    return(
    <div className={st["search"]}>
            <input type="seach" placeholder={placehold} onInput={event}/>
            <i className='fa fa-magnifying-glass'></i>
    </div>
    )
}

