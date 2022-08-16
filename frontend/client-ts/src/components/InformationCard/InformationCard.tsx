import React from 'react'
import cs from './InformationCard.module.css'
console.log(cs)
console.log('hello')
type PropTypes = {
  title: string
  message: string
}
export const InformationCard: React.FC<PropTypes> = ({title, message}) => {
  return (
    <div className={cs.cardInformation}>
        <div className={cs.cardTitle}><strong>{title}</strong></div>
        <div className={cs.cardHorizontalLine}></div>
        <div className={cs.cardText}>{message}</div>
    </div>
  )
}
