import React, { FC } from "react";
import { Link } from "react-router-dom";
import st from './Buttons.module.scss';

interface PropTypes {
    path?: string
  }

export const GoBackButton: FC<PropTypes> = ({path = '#'}) => {
    return(
        <Link to={path}>
            <button className={st["button-go-back"]}><i className='fa-solid fa-circle-arrow-left'></i></button>
        </Link>
    )
}



interface PropTypesPlus {
    event?: () => void
}


export const PlusButton: FC<PropTypesPlus> = ({event = () => {console.log("Кнопка '+' была нажата")}}) => {
    return(
        <button className={st["button-add"]} onClick={event}><i className='fa fa-plus'></i></button>
    )
}