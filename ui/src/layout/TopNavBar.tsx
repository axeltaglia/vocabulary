import Toolbar from "@mui/material/Toolbar"
import AppBar from "@mui/material/AppBar"
import * as React from "react"
import {Typography} from "@mui/material"

function TopNavBar() {
    return <AppBar position="fixed" >
        <Toolbar>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                Vocabulary
            </Typography>
        </Toolbar>
    </AppBar>
}

export default TopNavBar

