import React, { FC } from "react";
import st from './PlusButton.module.scss';

type PropTypes = {
    event?: () => void
  }

export const PlusButton: FC<PropTypes> = ({event = () => {console.log("Кнопка '+' была нажата")}}) => {
    return(
        <button className={st["button-add"]} onClick={event}><i className='fa fa-plus'></i></button>
    )
}