import {Vocabulary, VocabularyCategory, VocabularyState, VocabularyWithCategoriesRequest} from "./types";
import {ChildrenType} from "../types";
import {ReactElement, useReducer} from "react";
import {vocabularyReducer} from "./VocabularyReducer";
import {AxiosResponse} from "axios";
import {VocabularyApi} from "../../api/vocabularyApi";
import {VocabularyContext} from "./VocabularyContext";
import {useGlobal} from "../GlobalContext/GlobalContext";
import {getErrorMessage} from "../../utils";

const INITIAL_STATE: VocabularyState = {
    vocabularies: [],
    createVocabularyDialogVisible: false,
    updateVocabularyDialogVisible: false,
    deleteVocabularyDialogVisible: false,
    vocabularyIdToUpdate: null,
    vocabularyToDelete: null,
    categories: [],
    loading: false
}

export default function VocabularyProvider({children}: ChildrenType): ReactElement {
    const [state, dispatch] = useReducer(vocabularyReducer, INITIAL_STATE)
    const {alertSuccessMsg, alertErrorMsg} = useGlobal()

    const getVocabularies = async () => {
        dispatch({type: 'SET_LOADING', payload: true})
        try {
            const vocabularyData: AxiosResponse = await VocabularyApi.getVocabularies()
            dispatch({type: 'SET_VOCABULARIES', payload: vocabularyData.data})
            dispatch({type: 'SET_LOADING', payload: false})
            return vocabularyData.data
        } catch (e) {
            dispatch({type: 'SET_LOADING', payload: false})
            const msg : string = getErrorMessage(e)
            alertErrorMsg(msg)
        }
    }

    const getCategories = async () => {
        dispatch({type: 'SET_LOADING', payload: true})
        try {
            const categoryData: AxiosResponse = await VocabularyApi.getCategories()
            dispatch({type: 'SET_CATEGORIES', payload: categoryData.data})
            dispatch({type: 'SET_LOADING', payload: false})
            return categoryData.data
        } catch (e) {
            dispatch({type: 'SET_LOADING', payload: false})
            const msg : string = getErrorMessage(e)
            alertErrorMsg(msg)
        }
    }

    const getVocabularyCategories = async (id: number) => {
        try {
            const data: AxiosResponse = await VocabularyApi.getVocabularyCategories(id)
            const vocabularyCategory: VocabularyCategory = {
                id: id,
                categories: data.data
            }
            dispatch({type: 'UPDATE_VOCABULARY_CATEGORIES', payload: vocabularyCategory})
            return data.data
        } catch (e) {
            dispatch({type: 'SET_LOADING', payload: false})
            const msg : string = getErrorMessage(e)
            alertErrorMsg(msg)
        }
    }

    const createVocabulary = async (newVocabulary: Vocabulary) => {
        dispatch({type: 'SET_LOADING', payload: true})
        try {
            const data: AxiosResponse = await VocabularyApi.createVocabulary(newVocabulary)
            dispatch({type: 'ADD_VOCABULARY', payload: data.data})
            dispatch({type: 'SET_LOADING', payload: false})
            alertSuccessMsg("New vocabulary created")
            return data.data
        } catch (e) {
            dispatch({type: 'SET_LOADING', payload: false})
            const msg : string = getErrorMessage(e)
            alertErrorMsg(msg)
        }
    }

    const updateVocabulary = async (vocabulary: Vocabulary) => {
        dispatch({type: 'SET_LOADING', payload: true})
        try {
            const data: AxiosResponse = await VocabularyApi.updateVocabulary(vocabulary)
            dispatch({type: 'UPDATE_VOCABULARY', payload: vocabulary})
            dispatch({type: 'SET_LOADING', payload: false})
            alertSuccessMsg("Vocabulary updated")
            return data.data
        } catch (e) {
            dispatch({type: 'SET_LOADING', payload: false})
            const msg : string = getErrorMessage(e)
            alertErrorMsg(msg)
        }
    }

    const updateVocabularyWithCategories = async (vocabularyWithCategories: VocabularyWithCategoriesRequest) => {
        dispatch({type: 'SET_LOADING', payload: true})
        try {
            const data: AxiosResponse = await VocabularyApi.updateVocabularyWithCategories(vocabularyWithCategories)
            dispatch({type: 'UPDATE_VOCABULARY', payload: data.data})
            dispatch({type: 'SET_LOADING', payload: false})
            alertSuccessMsg("The vocabulary was updated")
            getCategories()
            return data.data
        } catch (e) {
            dispatch({type: 'SET_LOADING', payload: false})
            const msg : string = getErrorMessage(e)
            alertErrorMsg(msg)
        }
    }
    const deleteVocabulary = async (vocabulary: Vocabulary) => {
        dispatch({type: 'SET_LOADING', payload: true})
        try {
            const data: AxiosResponse = await VocabularyApi.deleteVocabulary(vocabulary.id)
            dispatch({type: 'DELETE_VOCABULARY', payload: vocabulary})
            dispatch({type: 'SET_LOADING', payload: false})
            alertSuccessMsg("Vocabulary deleted")
            return data.data
        } catch (e) {
            dispatch({type: 'SET_LOADING', payload: false})
            const msg : string = getErrorMessage(e)
            alertErrorMsg(msg)
        }
    }

    const createVocabularyWithCategories = async (newVocabularyWithCategories: VocabularyWithCategoriesRequest) => {
        dispatch({type: 'SET_LOADING', payload: true})
        try {
            const data: AxiosResponse = await VocabularyApi.createVocabularyWithCategories(newVocabularyWithCategories)
            dispatch({type: 'ADD_VOCABULARY', payload: data.data})
            dispatch({type: 'SET_LOADING', payload: false})
            alertSuccessMsg("New vocabulary created")
            getCategories()
            return data.data
        } catch (e) {
            dispatch({type: 'SET_LOADING', payload: false})
            const msg : string = getErrorMessage(e)
            alertErrorMsg(msg)
        }
    }

    const openCreateVocabularyDialog = () => {
        dispatch({type: 'SET_CREATE_VOCABULARY_DIALOG_VISIBLE', payload: true})
    }

    const closeCreateVocabularyDialog = () => {
        dispatch({type: 'SET_CREATE_VOCABULARY_DIALOG_VISIBLE', payload: false})
    }

    const openUpdateVocabularyDialog = (id: number) => {
        dispatch({type: 'SET_VOCABULARY_ID_TO_UPDATE', payload: id})
        dispatch({type: 'SET_UPDATE_VOCABULARY_DIALOG_VISIBLE', payload: true})
    }

    const closeUpdateVocabularyDialog = () => {
        dispatch({type: 'SET_UPDATE_VOCABULARY_DIALOG_VISIBLE', payload: false})
    }

    const openDeleteVocabularyDialog = (id: number) => {
        dispatch({type: 'SET_VOCABULARY_TO_DELETE', payload: id})
        dispatch({type: 'SET_DELETE_VOCABULARY_DIALOG_VISIBLE', payload: true})
    }

    const closeDeleteVocabularyDialog = () => {
        dispatch({type: 'SET_DELETE_VOCABULARY_DIALOG_VISIBLE', payload: false})
    }

    const getVocabularyById = (id: number): Vocabulary | undefined => {
        return state.vocabularies.find((v) => v.id === id)
    }

    return <VocabularyContext.Provider value={{
        state,
        getVocabularies,
        updateVocabulary,
        createVocabulary,
        deleteVocabulary,
        createVocabularyWithCategories,
        updateVocabularyWithCategories,
        openCreateVocabularyDialog,
        closeCreateVocabularyDialog,
        openUpdateVocabularyDialog,
        closeUpdateVocabularyDialog,
        openDeleteVocabularyDialog,
        closeDeleteVocabularyDialog,
        getCategories,
        getVocabularyCategories,
        getVocabularyById
    }}>
        {children}
    </VocabularyContext.Provider>
}