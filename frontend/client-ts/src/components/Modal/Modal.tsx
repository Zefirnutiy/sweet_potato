import React from 'react';
import st from './Modal.module.css'
type PropTypes = {
    onClose: () => void
    headerText: string;
    children?: React.ReactNode
}

export const Modal: React.FC<PropTypes> = ({onClose, headerText, children}) =>{
    return (
        <>
        <div className={st.modal} onClick={onClose} />
        <div className={st.modalContent}>
            <strong>{headerText}</strong>
            {children}
        </div>
        </>     
    )
}