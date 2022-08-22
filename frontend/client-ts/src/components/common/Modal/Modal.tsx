import React from 'react';
import st from './Modal.module.scss'
import { FC } from "react";


// type PropTypes = {
//     onClose: () => void
//     headerText: string;
//     children?: React.ReactNode
// }

// export const Modal: FC<PropTypes> = ({onClose, headerText, children}) =>{
//     return (
//         <>
//         <div className={st["modal"]} onClick={onClose} />
//         <div className={st["modal-content"]}>
//             <strong>{headerText}</strong>
//             {children}
//         </div>
//         </>     
//     )
// }


type PropTypes = {
    onClose?: () => void
    cross?: boolean
    children: React.ReactNode
}

export const Modal: FC<PropTypes> = ({onClose, children, cross=false}) =>{
    return (
        <>
        <div className={st["modal"]} onClick={onClose} />
        <div className={st["modal-content"]}>
            {cross && <button className={st['cross']} onClick={onClose}><i className="fa fa-xmark"></i></button>}
            {children}
        </div>
        </>     
    )
}