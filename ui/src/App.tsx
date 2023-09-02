import React from "react";
import {
    CssBaseline,
    ThemeProvider,
} from "@mui/material";
import { createTheme } from "@mui/material/styles";
import WebRouter from "./routes/WebRouter";
import GlobalWrapper from "./components/GlobalWrapper";
import GlobalProvider from "./contexts/GlobalContext/GlobalContext";
import VocabularyProvider from "./contexts/VocabularyContext/VocabularyProvider";

function App() {
    const theme = createTheme();

    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <GlobalProvider>
                <GlobalWrapper>
                    <VocabularyProvider>
                        <WebRouter />
                    </VocabularyProvider>
                </GlobalWrapper>
            </GlobalProvider>
        </ThemeProvider>
    );
}

export default App;