import React from "react";
import st from './PlusButton.module.scss';

export const PlusButton = () => {
    return(
        <button className={st["button-add"]}><i className='fa fa-plus'></i></button>
)
}