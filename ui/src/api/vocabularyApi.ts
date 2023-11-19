import {api} from "./config/axiosConfig"
import {Category, Vocabulary, VocabularyWithCategoriesRequest} from "../contexts/VocabularyContext/types";

export const VocabularyApi = {
    getVocabularies: function() {
        return api.post<Vocabulary[]>("/getVocabularies");
    },
    updateVocabularyWithCategories: function(payload: VocabularyWithCategoriesRequest) {
        return api.post<VocabularyWithCategoriesRequest>("/updateVocabularyWithCategories", payload);
    },
    deleteVocabulary: function(id: number | undefined) {
        return api.post("/deleteVocabulary/" + id)
    },
    createVocabularyWithCategories: function(payload: VocabularyWithCategoriesRequest) {
        return api.post<VocabularyWithCategoriesRequest>("/createVocabularyWithCategories", payload);
    },
    getCategories: function() {
        return api.post<Category[]>("/getCategories");
    },
    getVocabularyCategories: function(id: number | undefined) {
        return api.post<Category[]>("/getVocabularyCategories/" + id)
    },
}