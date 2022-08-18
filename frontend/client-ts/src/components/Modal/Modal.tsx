import React from 'react';
<<<<<<< HEAD
import styles from './Modal.module.scss'
=======
import st from './Modal.module.css'
>>>>>>> 9059103f366f48ad18d61169c9c7ef12c9b63b01
type PropTypes = {
    onClose: () => void
    headerText: string;
    children?: React.ReactNode
}

export const Modal: React.FC<PropTypes> = ({onClose, headerText, children}) =>{
    return (
        <>
<<<<<<< HEAD
        <div className={styles.modalContent}>
            <strong>{headerText}</strong>
            {children}
        </div>
        <div className={styles.modal +' '+ styles.open} onClick={onClose} />
=======
        <div className={st.modal} onClick={onClose} />
        <div className={st.modalContent}>
            <strong>{headerText}</strong>
            {children}
        </div>
>>>>>>> 9059103f366f48ad18d61169c9c7ef12c9b63b01
        </>     
    )
}