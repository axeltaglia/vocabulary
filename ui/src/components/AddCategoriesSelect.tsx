import React, {useEffect} from "react";
import {Autocomplete, Chip} from "@mui/material";
import TextField from "@mui/material/TextField";

export default function AddCategoriesSelect() {

    const top100Films = [
        { title: 'The Shawshank Redemption', year: 1994 },
        { title: 'The Godfather', year: 1972 },
        { title: 'The Godfather: Part II', year: 1974 },
        { title: 'The Dark Knight', year: 2008 },
        { title: '12 Angry Men', year: 1957 },
        { title: "Schindler's List", year: 1993 },
        { title: 'Pulp Fiction', year: 1994 },
        { title: 'The Lord of the Rings: The Return of the King', year: 2003}
    ]
    return (
        <Autocomplete
            multiple
            id="categories"
            options={top100Films.map((option) => option.title)}
            freeSolo
            renderTags={(value: readonly string[], getTagProps) =>
                value.map((option: string, index: number) => (
                    <Chip variant="outlined" label={option} {...getTagProps({ index })} />
                ))
            }
            renderInput={(params) => (
                <TextField
                    {...params}
                    variant="standard"
                    label="Categories"
                />
            )}
        />
    )
}