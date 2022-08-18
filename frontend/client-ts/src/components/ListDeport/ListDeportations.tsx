import React from "react";
import styles from './ListDeportations.module.css'

type TypeProps ={
    Department?: string;
    AmountGroup?: string;
}

export const ListDeportations: React.FC<TypeProps> = ({Department, AmountGroup}) => {
    return(
        <>
            <div id={styles['list-deportations']}>
                <div className={styles['card-deportation']}>
                    <div className={styles['title-deportation']}>{Department}</div>
                    <div className={styles['number-groups']}>{AmountGroup}</div>
                </div>
            </div>
        </>
    )
}