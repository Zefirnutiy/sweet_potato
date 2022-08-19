import React from "react";
import st from './CardUsers.module.scss';

type PropsType = {
    userName: string;
    autorCreate: string;
    dateCreate: string;

}
export const CardUsers: React.FC<PropsType> = ({userName, autorCreate, dateCreate}) => {
    return(
        <div className={st["card-user"]}>
            <div className={st["user-name"]}>{userName}</div>
            <div className={st["created"]}>Создал: {autorCreate}<br/> Создан: {dateCreate}</div>
        </div>
)
}