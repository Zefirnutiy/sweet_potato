import { FC } from "react";
import st from './AddFolderForm.module.scss';

interface PropTypesPlus {
    event?: () => void
    title: string
    onClose: () => void
}


export const AddFolderForm: FC<PropTypesPlus> = ({event = () => {}, title, onClose}) => {
    function click(){
        event()
        onClose()
    }
    return(
        <div className={st["wrapper-add-folder"]}>
          <strong>{title} :</strong> <input type="text" /><button onClick={click}>Готово</button>
        </div>
    )
}
