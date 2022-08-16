import React from 'react';
import cs from'./Modal.module.css'
type PropTypes = {
    onClose: () => void
    headerText: string;
    children?: React.ReactNode
}

export const Modal: React.FC<PropTypes> = ({onClose, headerText, children}) =>{
    return (
        <>
        <div className={cs.modalContent}>
            <strong>{headerText}</strong>
            {children}
        </div>
        <div className={cs.modal + ' ' + cs.open} onClick={onClose} />
        </>     
      
        
    )
}