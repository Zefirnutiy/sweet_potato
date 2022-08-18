import React from 'react';
import st from './Modal.module.css'
import { FC } from "react";

type PropTypes = {
    onClose: () => void
    headerText: string;
    children?: React.ReactNode
}

export const Modal: FC<PropTypes> = ({onClose, headerText, children}) =>{
    return (
        <>
        <div className={st["modal"]} onClick={onClose} />
        <div className={st["modal-content"]}>
            <strong>{headerText}</strong>
            {children}
        </div>
        </>     
    )
}