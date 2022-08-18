import React from 'react';
import styles from './Modal.module.scss'
type PropTypes = {
    onClose: () => void
    headerText: string;
    children?: React.ReactNode
}

export const Modal: React.FC<PropTypes> = ({onClose, headerText, children}) =>{
    return (
        <>
        <div className={styles.modalContent}>
            <strong>{headerText}</strong>
            {children}
        </div>
        <div className={styles.modal +' '+ styles.open} onClick={onClose} />
        </>     
      
        
    )
}