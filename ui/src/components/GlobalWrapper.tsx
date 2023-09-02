import React, {ReactNode} from 'react'
import {Alert, Snackbar} from "@mui/material"
import {useGlobal} from "../contexts/GlobalContext/GlobalContext"

type WrapperProps = {
    children: ReactNode
}

function GlobalWrapper({ children }: WrapperProps): JSX.Element {
    const {alertMsgData, closeAlertMsg} = useGlobal()
    const handleOnClose = () => {
        closeAlertMsg()
    }

    return <>
        {children}
        <Snackbar open={alertMsgData.visible} autoHideDuration={6000} onClose={handleOnClose}>
            <Alert onClose={handleOnClose} severity={alertMsgData.type} sx={{ width: '100%' }}>
                {alertMsgData.message}
            </Alert>
        </Snackbar>
    </>
}

export default GlobalWrapper