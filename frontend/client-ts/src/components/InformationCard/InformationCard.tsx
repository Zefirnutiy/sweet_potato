import React from 'react'
import styles from './InformationCard.module.css'
type PropTypes = {
  title: string
  message: string
}
export const InformationCard: React.FC<PropTypes> = ({title, message}) => {
  return (
    <div className={styles.cardInformation}>
        <div className={styles.cardTitle}><strong>{title}</strong></div>
        <div className={styles.cardHorizontalLine}></div>
        <div className={styles.cardText}>{message}</div>
    </div>
  )
}
