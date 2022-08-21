import React from "react";
import { Link } from "react-router-dom";
import st from './CardUsers.module.scss';

type PropsTypesCardsUsers = {
    userName: string;
    autorCreate: string;
    dateCreate: string;

}

type PropsTypesCardsInformation = {
    title: string
    message: string
    path?: string
}


type PropsTypesCardsDeportation = {
    title: string
    number: number
}



export const CardUsers: React.FC<PropsTypesCardsUsers> = ({userName, autorCreate, dateCreate}) => (
        <div className={st["card-user"]}>
            <div className={st["user-name"]}>{userName}</div>
            <div className={st["created"]}>Создал: {autorCreate}<br/> Создан: {dateCreate}</div>
        </div>
)


export const InformationCard: React.FC<PropsTypesCardsInformation> = ({title, message, path = "#"}) => (
    <Link to={path}>
     <div className={st["card-information"]}>
        <div className={st["card-title"]}><strong>{title}</strong></div>
        <div className={st["card-horizontal-line"]}></div>
        <div className={st["card-text"]}>{message}</div>
    </div>
    </Link>
  )


export const DeportationCard: React.FC<PropsTypesCardsDeportation> = ({title, number}) => (
    <div className={st['card-deportation']}>
            <div className={st['title-deportation']}>{title}</div>
            <div className={st['number-groups']}>{number} групп</div>
    </div>
)




