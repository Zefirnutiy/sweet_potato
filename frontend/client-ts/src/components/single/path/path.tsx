import st from "./Path.module.scss"
interface PropsTypePath{
    pathDepartament: string
    pathGroup: string
    showForPath: (isDepartament: boolean) => void
}

export const Path: React.FC<PropsTypePath> = ({pathDepartament, pathGroup, showForPath}) => {
      if(pathDepartament && !pathGroup)
        return <div className={st["path"]} >
            <div onClick={() => showForPath(true)}>/{pathDepartament}</div>
        </div>
      
      if(pathGroup)
        return <div className={st["path"]}>
                <div onClick={() => showForPath(true)}>
                    /{pathDepartament}
                </div>
                <div onClick={() => showForPath(false)}>
                    /{pathGroup}
                </div>
            </div>

      return null
}
