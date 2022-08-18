import React from 'react'
<<<<<<< HEAD
import styles from './InformationCard.module.scss'
=======
import st from './InformationCard.module.css'
>>>>>>> 9059103f366f48ad18d61169c9c7ef12c9b63b01
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
