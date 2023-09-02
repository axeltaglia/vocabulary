import {createContext, useContext} from 'react'
import {VocabularyContextProps} from './types'

export const VocabularyContext = createContext<VocabularyContextProps>({} as VocabularyContextProps)

export const useVocabulary = () => {
    return useContext(VocabularyContext)
}