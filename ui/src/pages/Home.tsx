import React, {useEffect} from "react"
import {Box, Button, Typography} from "@mui/material"
import Toolbar from "@mui/material/Toolbar"
import TopNavBar from "../layout/TopNavBar";
import {useVocabulary} from "../contexts/VocabularyContext/VocabularyContext";
import VocabularyList from "../components/VocabularyList";
import CreateVocabularyDialog from "../components/CreateVocabularyDialog";
import UpdateVocabularyDialog from "../components/UpdateVocabularyDialog";
import {AddCircle} from "@mui/icons-material";
import DeleteVocabularyDialog from "../components/DeleteVocabularyDialog";

function Home() {
    const {
        state: {vocabularies},
        getVocabularies,
        getCategories,
        openCreateVocabularyDialog
    } = useVocabulary()

    const handleNewVocabularyButtonClick = () => {
        openCreateVocabularyDialog()
    }

    useEffect(() => {
        getVocabularies();
        getCategories();
    }, [])

    return <>
        <Box sx={{ display: 'flex' }}>
            <TopNavBar />
            <Box component="main" sx={{ flexGrow: 1, bgcolor: 'background.default', p: 3 }}>
                <Toolbar />
                <Box
                    sx={{paddingBottom: "25px", textAlign: "right"}}
                >
                <Button variant="contained" color="success" startIcon={<AddCircle />}
                        onClick={handleNewVocabularyButtonClick}>NEW VOCABULARY</Button>
                </Box>
                { vocabularies.length > 0 ?
                        <VocabularyList />
                :
                    <Typography variant="h6" textAlign={"center"} component="div" sx={{ flexGrow: 1 }}>
                        Vocabularies haven't been created yet
                    </Typography>
                }
                <CreateVocabularyDialog />
                <UpdateVocabularyDialog />
                <DeleteVocabularyDialog />
            </Box>
        </Box>
    </>
}

export default Home