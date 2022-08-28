import React, { FC } from "react";
import { Link } from "react-router-dom";
import st from './Buttons.module.scss';

interface PropTypesPlus {
    event?: () => void
}


export const AddButton: FC<PropTypesPlus> = ({event = () => {console.log("Кнопка '+' была нажата")}}) => {
    return(
        <button className={st["button-add"]} onClick={event}><i className='fa fa-plus'></i></button>
    )
}