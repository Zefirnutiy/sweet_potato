import React from "react";
import st from './Search.module.scss';

type PropsType = {
    placehold: string;

}
export const Search: React.FC<PropsType> = ({placehold}) => {
    return(
    <div className={st["search"]}>
        <input type="seach" placeholder={placehold}/>
         <i className='fa fa-magnifying-glass'></i>
         </div>
)
}