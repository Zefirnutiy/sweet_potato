import React from 'react'
import './InformationCard.css'

type PropTypes = {
  title: string
  message: string
}
export const InformationCard: React.FC<PropTypes> = ({title, message}) => {
  return (
    <div className="cardInformation">
        <div className="cardTitle"><strong>{title}</strong></div>
        <div className="cardHorizontalLine"></div>
        <div className="cardText">{message}</div>
    </div>
  )
}
