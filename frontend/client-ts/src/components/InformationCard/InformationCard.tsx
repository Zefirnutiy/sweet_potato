import React from 'react'
import st from './InformationCard.module.css'
type PropTypes = {
  title: string
  message: string
}
export const InformationCard: React.FC<PropTypes> = ({title, message}) => {
  return (
    <div className={st.card_information}>
        <div className={st.card_title}><strong>{title}</strong></div>
        <div className={st.card_horizontal_line}></div>
        <div className={st.card_text}>{message}</div>
    </div>
  )
}
