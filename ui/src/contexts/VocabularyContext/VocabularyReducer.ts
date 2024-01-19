import {Category, Vocabulary, VocabularyCategory, VocabularyState} from './types'

type VocabularyAction =
    | { type: 'SET_LOADING'; payload: boolean }
    | { type: 'SET_VOCABULARIES'; payload: Vocabulary[] }
    | { type: 'ADD_VOCABULARY'; payload: Vocabulary }
    | { type: 'DELETE_VOCABULARY'; payload: Vocabulary }
    | { type: 'UPDATE_VOCABULARY'; payload: Vocabulary }
    | { type: 'SET_CREATE_VOCABULARY_DIALOG_VISIBLE'; payload: boolean }
    | { type: 'SET_UPDATE_VOCABULARY_DIALOG_VISIBLE'; payload: boolean }
    | { type: 'SET_DELETE_VOCABULARY_DIALOG_VISIBLE'; payload: boolean }
    | { type: 'SET_VOCABULARY_ID_TO_UPDATE'; payload: number }
    | { type: 'SET_VOCABULARY_TO_DELETE'; payload: number }
    | { type: 'SET_CATEGORIES'; payload: Category[] }
    | { type: 'MERGE_CATEGORIES'; payload: Category[] }
    | { type: 'UPDATE_VOCABULARY_CATEGORIES'; payload: VocabularyCategory }

export const vocabularyReducer = (state: VocabularyState, action: VocabularyAction): VocabularyState => {
    switch (action.type) {
        case 'SET_VOCABULARIES':
            return {
                ...state,
                vocabularies: action.payload,
            }
        case 'ADD_VOCABULARY':
            return {
                ...state,
                vocabularies: [action.payload, ...state.vocabularies]
            }
        case 'UPDATE_VOCABULARY':
            const editedList = state.vocabularies.map(
                (item) => item.id === action.payload.id? action.payload : item
            );
            return {
                ...state,
                vocabularies: editedList,
            };
        case 'DELETE_VOCABULARY':
            const updatedList = state.vocabularies.filter(
                (item) => item.id !== action.payload.id
            );
            return {
                ...state,
                vocabularies: updatedList,
            };
        case 'SET_LOADING':
            return {
                ...state,
                loading: action.payload,
            }
        case 'SET_CREATE_VOCABULARY_DIALOG_VISIBLE':
            return {
                ...state,
                createVocabularyDialogVisible: action.payload
            }
        case 'SET_UPDATE_VOCABULARY_DIALOG_VISIBLE':
            return {
                ...state,
                updateVocabularyDialogVisible: action.payload
            }
        case 'SET_DELETE_VOCABULARY_DIALOG_VISIBLE':
            return {
                ...state,
                deleteVocabularyDialogVisible: action.payload
            }
        case 'SET_VOCABULARY_ID_TO_UPDATE':
            return {
                ...state,
                vocabularyIdToUpdate: action.payload
            }
        case 'SET_VOCABULARY_TO_DELETE':
            const vocabularyToDelete = state.vocabularies.find((vocabulary) => vocabulary.id === action.payload)
            return {
                ...state,
                vocabularyToDelete: vocabularyToDelete
            }
        case 'SET_CATEGORIES':
            return {
                ...state,
                categories: action.payload,
            }
        case 'MERGE_CATEGORIES':
            return {
                ...state,
                categories: [...state.categories, ...action.payload.filter(category => !state.categories.some(stateCategory => stateCategory.id === category.id))
                ]
            }
        case 'UPDATE_VOCABULARY_CATEGORIES':
            const updatedVocabularyCategoriesList = state.vocabularies.map((item) => {
                if(item.id === action.payload.id) {
                    item.categories = action.payload.categories
                }
                return item
            });

            return {
                ...state,
                vocabularies: updatedVocabularyCategoriesList,
            };
        default:
            throw new Error()
    }
}
