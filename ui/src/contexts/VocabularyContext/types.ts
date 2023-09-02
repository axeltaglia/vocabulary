export type Vocabulary = {
    id?: number
    words: string
    translation: string
    usedInPhrase: string
    explanation: string
    categories?: Category[]
}
export type Category = {
    id?: number
    name: string
}
export type VocabularyCategory = {
    id?: number
    categories: Category[]
}

export type VocabularyWithCategoriesRequest = {
    vocabulary: Vocabulary
    categories: string[]
}
export type VocabularyState = {
    vocabularies: Vocabulary[]
    createVocabularyDialogVisible: boolean
    updateVocabularyDialogVisible: boolean
    deleteVocabularyDialogVisible: boolean
    vocabularyIdToUpdate: number | null
    vocabularyToDelete: Vocabulary | null | undefined
    categories: Category[]
    loading: boolean
}
export type VocabularyContextProps = {
    state: VocabularyState,
    getVocabularies: () => Promise<void>
    updateVocabulary: (vocabulary: Vocabulary) => Promise<void>
    createVocabulary: (vocabulary: Vocabulary) => Promise<void>
    deleteVocabulary: (vocabulary: Vocabulary) => Promise<void>
    createVocabularyWithCategories: (vocabularyWithCategories: VocabularyWithCategoriesRequest) => Promise<void>
    updateVocabularyWithCategories: (vocabularyWithCategories: VocabularyWithCategoriesRequest) => Promise<void>
    openCreateVocabularyDialog: () => void
    closeCreateVocabularyDialog: () => void
    openUpdateVocabularyDialog: (id: number) => void
    closeUpdateVocabularyDialog: () => void
    openDeleteVocabularyDialog: (id: number) => void
    closeDeleteVocabularyDialog: () => void
    getCategories: () => Promise<void>
    getVocabularyCategories: (id: number) => Promise<void>
    getVocabularyById: (id: number) => Vocabulary | undefined
}