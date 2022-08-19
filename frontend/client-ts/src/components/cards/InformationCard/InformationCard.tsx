import {FC} from 'react'
import { Link } from 'react-router-dom'
import st from './InformationCard.module.scss'
        
type PropTypes = {
  title: string
  message: string
  path?: string
}

export const InformationCard: FC<PropTypes> = ({title, message, path = "#"}) => {
  return (
    <Link to={path}>
     <div className={st["card-information"]}>
        <div className={st["card-title"]}><strong>{title}</strong></div>
        <div className={st["card-horizontal-line"]}></div>
        <div className={st["card-text"]}>{message}</div>
    </div>
    </Link>
  )
}
