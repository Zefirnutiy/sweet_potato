import {FC} from 'react'
import st from './InformationCard.module.scss'
        
type PropTypes = {
  title: string
  message: string
  event?: () => void
}

export const InformationCard: FC<PropTypes> = ({title, message, event = () => {alert("Пока пусто")}}) => {
  return (
    <div className={st["card-information"]} onClick={event}>
        <div className={st["card-title"]}><strong>{title}</strong></div>
        <div className={st["card-horizontal-line"]}></div>
        <div className={st["card-text"]}>{message}</div>
    </div>
  )
}
