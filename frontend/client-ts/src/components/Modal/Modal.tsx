import React from 'react';
import './Modal.css'
type PropTypes = {
    onClose: () => void
    headerText: string;
    children?: React.ReactNode
}

export const Modal: React.FC<PropTypes> = ({onClose, headerText, children}) =>{
    return (
        <>
        <div className="modalContent">
            <strong>{headerText}</strong>
            {children}
        </div>
        <div className="modal open" onClick={onClose} />
        </>     
      
        
    )
}