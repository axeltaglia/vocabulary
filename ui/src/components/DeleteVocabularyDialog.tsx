import React from "react";
import {useTheme} from "@mui/material/styles";
import useMediaQuery from "@mui/material/useMediaQuery";
import {useVocabulary} from "../contexts/VocabularyContext/VocabularyContext";
import {Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle} from "@mui/material";

import {Vocabulary} from "../contexts/VocabularyContext/types";

export default function DeleteVocabularyDialog() {
    const {
        state: { deleteVocabularyDialogVisible, vocabularyToDelete},
        closeDeleteVocabularyDialog,
        deleteVocabulary,
    } = useVocabulary()

    const handleCancelButtonClick = () => {
        closeDeleteVocabularyDialog()
    };

    const handleDeleteVocabularyButtonClick = () => {
        const vocabulary = vocabularyToDelete as Vocabulary
        deleteVocabulary(vocabulary)
            .then(() => {
                closeDeleteVocabularyDialog()
            })
    };

    const theme = useTheme();
    const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

    return (
        <Dialog
            open={deleteVocabularyDialogVisible}
            onClose={handleCancelButtonClick}
            fullScreen={fullScreen}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
        >
            <DialogTitle id="alert-dialog-title">Confirm Deletion</DialogTitle>
            <DialogContent>
                <DialogContentText id="alert-dialog-description">
                    Are you sure you want to delete this vocabulary item?
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={handleCancelButtonClick} color="primary">
                    Cancel
                </Button>
                <Button onClick={handleDeleteVocabularyButtonClick} color="primary" autoFocus>
                    Confirm
                </Button>
            </DialogActions>
        </Dialog>
    )
}