import React from "react";
import styles from './Search.module.css'

type PropsType = {
    placehold: string;
}

export const Search: React.FC<PropsType> = ({placehold}) => {
    return (
    <div className={styles.search}>
        <input type="search" placeholder={placehold}/>
        <i className={styles.fa + ' ' + styles.fa_magnifying_glass}></i>
    </div>

    )
}