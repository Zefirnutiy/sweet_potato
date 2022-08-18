import {FC} from 'react'
import st from './InformationCard.module.css'

type PropTypes = {
  title: string
  message: string
}

export const InformationCard: FC<PropTypes> = ({title, message}) => {
  return (
    <div className={st["card-information"]}>
        <div className={st["card-title"]}><strong>{title}</strong></div>
        <div className={st["card-horizontal-line"]}></div>
        <div className={st["card-text"]}>{message}</div>
    </div>
  )
}
