import * as React from 'react';
import { DataGrid, GridColDef } from '@mui/x-data-grid';
import {useVocabulary} from "../contexts/VocabularyContext/VocabularyContext";
import {IconButton} from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';



export default function VocabularyList() {
    const {
        state: { vocabularies},
        openDeleteVocabularyDialog,
        openUpdateVocabularyDialog,
        getVocabularyCategories
    } = useVocabulary()

    const columns: GridColDef[] = [
        {
            field: 'words',
            headerName: 'Words',
            sortable: false,
            width: 130,
            renderCell: (params) => (
                <div
                    style={{ color: 'blue', cursor: 'pointer' }}
                    onClick={() => handleWordsClick(params.row.id)}
                >
                    {params.row.words}
                </div>
            ),
        },
        {
            field: 'usedInPhrase',
            headerName: 'Use',
            sortable: false,
            flex: 1,
            renderCell: (params) => (
                <div className="custom-cell">
                    {params.row.usedInPhrase.length > 180 ? (
                        <>
                            {params.row.usedInPhrase.slice(0, 180)}...
                        </>
                    ) : (
                        params.row.usedInPhrase
                    )}
                </div>
            ),
        },
        {
            field: 'actions',
            headerName: 'Actions',
            sortable: false,
            width: 150,
            renderHeader: (params) => (
                <div style={{ textAlign: 'center' }}>Actions</div>
            ),
            renderCell: (params) => (
                <div>
                    <IconButton aria-label="edit" size="large" onClick={() => handleEdit(params.row.id)}>
                        <EditIcon fontSize="inherit" />
                    </IconButton>
                    <IconButton aria-label="delete" size="large" onClick={() => handleDelete(params.row.id)}>
                        <DeleteIcon fontSize="inherit" />
                    </IconButton>
                </div>
            ),
        },
    ];

    const handleWordsClick = async (id: number) => {
        await getVocabularyCategories(id)
        openUpdateVocabularyDialog(id)
    }

    const handleEdit = async (id: number) => {
        await getVocabularyCategories(id)
        openUpdateVocabularyDialog(id)
    };

    const handleDelete = (id: number) => {
        openDeleteVocabularyDialog(id)
    };

    return (
        <div style={{ width: '100%' }}>
            <DataGrid
                rows={vocabularies}
                columns={columns}
                checkboxSelection={false}
                disableRowSelectionOnClick
                pageSizeOptions={[5, 10, 20]}
                initialState={{
                    pagination: {
                        paginationModel: { page: 0, pageSize: 20 },
                    },
                }}

            />
        </div>
    );
}
