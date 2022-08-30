import st from "./Path.module.scss"
interface PropsTypePath{
    firstAddress: string
    secondAddress: string
    showForPath: (isDepartament: boolean) => void
}

export const Path: React.FC<PropsTypePath> = ({firstAddress, secondAddress, showForPath}) => {
      if(firstAddress && !secondAddress)
        return <div className={st["path"]} >
            <div onClick={() => showForPath(true)}>/{firstAddress}</div>
        </div>
      
      if(secondAddress)
        return <div className={st["path"]}>
                <div onClick={() => showForPath(true)}>
                    /{firstAddress}
                </div>
                <div onClick={() => showForPath(false)}>
                    /{secondAddress}
                </div>
            </div>

      return null
}
