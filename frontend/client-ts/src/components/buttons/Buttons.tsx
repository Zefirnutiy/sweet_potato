import React, { FC } from "react";
import st from './Buttons.module.scss';

interface PropTypesPlus {
    icon: string
    event: () => void
}

export const SquareButton: FC<PropTypesPlus> = ({icon, event = () => {}}) => {
    return(
        <button className={st["button-add"]} onClick={event}><i className={icon}></i></button>
    )
}