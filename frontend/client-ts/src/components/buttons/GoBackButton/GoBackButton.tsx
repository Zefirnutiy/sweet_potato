import React, { FC } from "react";
import { Link } from "react-router-dom";
import st from './GoBackButton.module.scss';

type PropTypes = {
    path?: string
  }

export const GoBackButton: FC<PropTypes> = ({path = '#'}) => {
    return(
        <Link to={path}>
            <button className={st["button-go-back"]}><i className='fa-solid fa-circle-arrow-left'></i></button>
        </Link>
    )
}