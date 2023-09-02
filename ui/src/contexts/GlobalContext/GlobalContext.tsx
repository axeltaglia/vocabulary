import {createContext, ReactElement, useContext, useState} from "react"
import {AlertMsgType} from "./types"
import {ChildrenType} from "../types"

const useGlobalContext = () => {
    const [alertMsgData, setAlertMsgData] = useState<AlertMsgType>({message: "", type: "success", visible: false})

    const alertSuccessMsg = (msg: string) => {
        setAlertMsgData({message: msg, type: 'success', visible: true})
    }

    const alertInfoMsg = (msg: string) => {
        setAlertMsgData({message: msg, type: 'info', visible: true})
    }

    const alertWarningMsg = (msg: string) => {
        setAlertMsgData({message: msg, type: 'warning', visible: true})
    }

    const alertErrorMsg = (msg: string) => {
        setAlertMsgData({message: msg, type: 'error', visible: true})
    }

    const closeAlertMsg = () => {
        setAlertMsgData({message: alertMsgData.message, type: alertMsgData.type, visible: false})
    }

    return { alertMsgData, alertSuccessMsg , alertInfoMsg, alertWarningMsg, alertErrorMsg, closeAlertMsg }
}

const GlobalContext = createContext({} as ReturnType<typeof useGlobalContext>)

export default function GlobalProvider({ children }: ChildrenType): ReactElement {
    return <GlobalContext.Provider value={useGlobalContext()}>
        {children}
    </GlobalContext.Provider>
}

export const useGlobal = () => {
    return useContext(GlobalContext)
}
