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

export type Vocabulary = {
    id?: number
    words: string
    translation: string
    usedInPhrase: string
    explanation: string
    categories: Category[]
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
    categoryNames: string[]
}

export type VocabularyContextProps = {
    state: VocabularyState,
    getVocabularies: () => Promise<void>
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