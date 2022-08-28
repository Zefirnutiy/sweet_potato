
interface PropsTypePath{
    pathDepartament: string
    pathGroup: string
    showForPath: (isDepartament: boolean) => void
}

export const Path: React.FC<PropsTypePath> = ({pathDepartament, pathGroup, showForPath}) => {
      if(pathDepartament && !pathGroup)
        return <h5 onClick={() => showForPath(true)}>/{pathDepartament}</h5>
      
      if(pathGroup)
        return <h5>
                <div onClick={() => showForPath(true)}>
                    /{pathDepartament}
                </div>
                <div onClick={() => showForPath(false)}>
                    /{pathGroup}
                </div>
            </h5>

      return null
}
