import React, {useEffect} from "react"
import {Box, Button, Typography} from "@mui/material"
import Toolbar from "@mui/material/Toolbar"
import TopNavBar from "../layout/TopNavBar";
import {useVocabulary} from "../contexts/VocabularyContext/VocabularyContext";
import {AddCircle} from "@mui/icons-material";
import CategoryList from "../components/CategoryList";

function Categories() {
    const {
        state: {categories},
        getCategories,
        openCreateVocabularyDialog
    } = useVocabulary()

    useEffect(() => {
        getCategories()
            .then(() => {
            })
    }, [])



    const handleNewCategoryButtonClick = () => {
        openCreateVocabularyDialog()
    }

    return <>
        <Box sx={{ display: 'flex' }}>
            <TopNavBar />
            <Box component="main" sx={{ flexGrow: 1, bgcolor: 'background.default', p: 3 }}>
                <Toolbar />
                <Box
                    sx={{paddingBottom: "25px", textAlign: "right"}}
                >
                    <Button variant="contained" color="success" startIcon={<AddCircle />}
                            onClick={handleNewCategoryButtonClick}>NEW CATEGORY</Button>
                </Box>
                { categories.length > 0 ?
                    <CategoryList />
                    :
                    <Typography variant="h6" textAlign={"center"} component="div" sx={{ flexGrow: 1 }}>
                        Categories haven't been created yet
                    </Typography>
                }
            </Box>
        </Box>
    </>
}

export default Categories