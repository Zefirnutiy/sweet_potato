import { FC } from "react";
import st from './AddFolder.module.scss';

interface PropTypesPlus {
    event?: () => void
}


export const AddFolder: FC<PropTypesPlus> = ({event = () => {}}) => {
    return(
        <div className={st["wrapper-add-folder"]}>
            <input type="text" /><button>Готово</button>
        </div>
    )
}