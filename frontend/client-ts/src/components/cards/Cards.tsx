import React from "react";
import { Link } from "react-router-dom";
import st from './CardUsers.module.scss';

interface PropsTypesUsersCard {
    userName: string;
    autorCreate: string;
    dateCreate: string;

}

interface PropsTypesInformationCard {
    title: string
    message: string
    path?: string
}


interface PropsTypesTwoCellsCard {
    title: string
    message: string
}



export const UserCard: React.FC<PropsTypesUsersCard> = ({userName, autorCreate, dateCreate}) => (
        <div className={st["card-user"]}>
            <div className={st["user-name"]}>{userName}</div>
            <div className={st["created"]}>Создал: {autorCreate}<br/> Создан: {dateCreate}</div>
        </div>
)


export const InformationCard: React.FC<PropsTypesInformationCard> = ({title, message, path = "#"}) => (
    <Link to={path}>
     <div className={st["card-information"]}>
        <div className={st["card-title"]}><strong>{title}</strong></div>
        <div className={st["card-horizontal-line"]}></div>
        <div className={st["card-text"]}>{message}</div>
    </div>
    </Link>
  )


export const TwoCellsCard: React.FC<PropsTypesTwoCellsCard> = ({title, message}) => (
    <div className={st['card-deportation']}>
            <div className={st['title-deportation']}>{title}</div>
            <div className={st['number-groups']}>{message}</div>
    </div>
)




