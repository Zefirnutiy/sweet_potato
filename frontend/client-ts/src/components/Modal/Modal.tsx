import './Modal.css'
type PropTypes = {
    openModal: boolean
    setActiveModal: (open: boolean) => void
    headerText: string;
    children?: React.ReactNode
}

export const Modal: React.FC<PropTypes> = ({openModal, setActiveModal, headerText, children}) =>{
    return (
       
       <div className={openModal ? "modal open" : "modal"} onClick={() => setActiveModal(false)}>
         <div className="modalContent" onClick={(e)=>{e.stopPropagation()}}>
            <strong>{headerText}</strong>
            {children}
        </div>
       </div>
        
    )
}