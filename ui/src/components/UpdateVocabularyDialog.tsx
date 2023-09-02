import React, {useEffect} from "react";
import {
    Autocomplete,
    Box,
    Button, Chip,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle, Stack,
} from "@mui/material";
import {useVocabulary} from "../contexts/VocabularyContext/VocabularyContext";
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import * as yup from "yup";
import {Controller, useForm} from "react-hook-form";
import {yupResolver} from "@hookform/resolvers/yup";
import TextField from "@mui/material/TextField";

import {Vocabulary, VocabularyWithCategoriesRequest} from "../contexts/VocabularyContext/types";

export default function UpdateVocabularyDialog() {
    let {
        state: { categories, updateVocabularyDialogVisible, vocabularyIdToUpdate},
        closeUpdateVocabularyDialog,
        updateVocabularyWithCategories,
        getVocabularyById,
    } = useVocabulary()

    const schema = yup.object().shape({
        words: yup.string().required('Words are required'),
        translation: yup.string(),
        usedInPhrase: yup.string(),
        explanation: yup.string(),
        categories: yup.array(),
    });

    const { control, handleSubmit, formState: { errors }, reset } = useForm({
        resolver: yupResolver(schema),
    })
    let vocabularyToUpdate: Vocabulary | undefined
    if(vocabularyIdToUpdate) {
        vocabularyToUpdate = getVocabularyById(vocabularyIdToUpdate)
    }

    useEffect(() => {
        reset({
            words: vocabularyToUpdate?.words,
            translation: vocabularyToUpdate?.translation,
            usedInPhrase: vocabularyToUpdate?.usedInPhrase,
            explanation: vocabularyToUpdate?.explanation,
            categories: vocabularyToUpdate?.categories?.map((c) => c.name)
        });
    }, [vocabularyToUpdate, reset])


    const handleCancelButtonClick = () => {
        closeUpdateVocabularyDialog()
    };

    const handleUpdateVocabularyButtonClick = (data: any) => {
        const vocabularyWithCategoriesRequest: VocabularyWithCategoriesRequest = {
            vocabulary: {
                id: vocabularyToUpdate?.id,
                words: data.words,
                translation: data.translation,
                usedInPhrase: data.usedInPhrase,
                explanation: data.explanation
            },
            categories: data.categories
        }

        updateVocabularyWithCategories(vocabularyWithCategoriesRequest)
            .then(() => {
                closeUpdateVocabularyDialog()
            })
    }

    const theme = useTheme();
    const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

    return (
        <Dialog
            open={updateVocabularyDialogVisible}
            onClose={handleCancelButtonClick}
            fullScreen={fullScreen}
            maxWidth={'md'}
            aria-labelledby="responsive-dialog-title"
        >
            <DialogTitle>Edit vocabulary</DialogTitle>
            <DialogContent>
                <Box component="form" noValidate sx={{ mt: 1 }}>
                    <Stack spacing={3} sx={{ width: 800 }}>

                        <Controller
                            name="words"
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <TextField
                                    required
                                    id="words"
                                    name="words"
                                    label="Words"
                                    value={value}
                                    onChange={onChange}
                                    fullWidth
                                    variant="standard"
                                    error={!!errors.words}
                                    helperText={errors.words?.message}
                                />
                            )}
                        />

                        <Controller
                            name="translation"
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <TextField
                                    id="translation"
                                    name="translation"
                                    label="Translation"
                                    value={value}
                                    onChange={onChange}
                                    fullWidth
                                    variant="standard"
                                    error={!!errors.translation}
                                    helperText={errors.translation?.message}
                                />
                            )}
                        />

                        <Controller
                            name="usedInPhrase"
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <TextField
                                    id="usedInPhrase"
                                    name="usedInPhrase"
                                    label="Used in phrase"
                                    value={value}
                                    onChange={onChange}
                                    fullWidth
                                    variant="standard"
                                    error={!!errors.usedInPhrase}
                                    helperText={errors.usedInPhrase?.message}
                                />
                            )}
                        />

                        <Controller
                            name="explanation"
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <TextField
                                    id="explanation"
                                    name="explanation"
                                    label="Explanation"
                                    value={value}
                                    onChange={onChange}
                                    fullWidth
                                    variant="standard"
                                    error={!!errors.explanation}
                                    helperText={errors.explanation?.message}
                                />
                            )}
                        />

                        <Controller
                            name="categories"
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <Autocomplete
                                    multiple
                                    freeSolo
                                    options={categories.map((category) => category.name)}
                                    value={value}
                                    onChange={(event, newValue) => {
                                        onChange(newValue);
                                    }}
                                    renderTags={(value: readonly string[], getTagProps) => {
                                        return value.map((option: string, index: number) => (
                                            <Chip variant="outlined" label={option} {...getTagProps({ index })} />
                                        ))
                                    }}
                                    renderInput={(params) => (
                                        <TextField
                                            {...params}
                                            label="Select options"
                                            variant="standard"
                                            error={!!errors.categories}
                                            helperText={errors.categories?.message}
                                        />
                                    )}
                                />
                            )}
                        />
                    </Stack>
                </Box>
            </DialogContent>
            <DialogActions>
                <Button onClick={handleCancelButtonClick}>Cancel</Button>
                <Button onClick={handleSubmit(handleUpdateVocabularyButtonClick)}>Submit</Button>
            </DialogActions>
        </Dialog>
    )
}