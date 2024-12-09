package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"net/http"
	"vocabulary/entities"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/logger"
)

type ApiServer struct {
	ListenPort          string
	TxRepositoryHandler entities.TxRepositoryHandler
	Router              *mux.Router
}

func NewApiServer(listenPort string, txRepositoryHandler entities.TxRepositoryHandler) ApiServer {
	router := mux.NewRouter()

	return ApiServer{
		ListenPort:          listenPort,
		TxRepositoryHandler: txRepositoryHandler,
		Router:              router,
	}
}

func (o *ApiServer) HandleEndpoints() {
	o.Router.HandleFunc("/test", test)
	o.Router.HandleFunc("/createVocabularyWithCategories", interceptCors(o.makeHttpFunc(createVocabulary)))
}

func interceptCors(httpFunc func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Allow preflight requests
		if r.Method == "OPTIONS" {
			return
		}

		// Call the next handler
		httpFunc(w, r)
	}
}

type CreateVocabulary struct {
	Words       string `json:"words"`
	Translation string `json:"translation"`
	UseInPhrase string `json:"useInPhrase"`
	Explanation string `json:"explanation"`
}

func (o *ApiServer) makeHttpFunc(f func(http.ResponseWriter, *http.Request, VocabularyEntity.Entity) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		txRepositoryFactory := o.TxRepositoryHandler.GetTxRepositoryFactory()
		vocabularyRepository := txRepositoryFactory.CreateVocabularyRepository()
		vocabularyEntity := VocabularyEntity.New(vocabularyRepository)
		if err := f(w, r, vocabularyEntity); err != nil {
			txRepositoryFactory.RollbackTransaction()
			if err = writeJson(w, http.StatusInternalServerError, err); err != nil {
				logger.GetLogger().LogError("endpoint error", err)
			}
			return
		}
		txRepositoryFactory.CommitTransaction()
	}
}

type CreateVocabularyWithCategories struct {
	Id           *uint      `json:"id"`
	Words        *string    `json:"words"`
	Translation  *string    `json:"translation"`
	UsedInPhrase *string    `json:"usedInPhrase"`
	Explanation  *string    `json:"explanation"`
	Categories   []Category `json:"categories"`
}

type Category struct {
	Id   *uint   `json:"id"`
	Name *string `json:"name"`
}

type Vocabulary struct {
	Id           *uint      `json:"id"`
	Words        *string    `json:"words" validate:"required,min=1"`
	Translation  *string    `json:"translation"`
	UsedInPhrase *string    `json:"usedInPhrase"`
	Explanation  *string    `json:"explanation"`
	Categories   []Category `json:"categories"`
}

type CreateVocabularyWithCategoriesRequest struct {
	Vocabulary    Vocabulary `json:"vocabulary"`
	CategoryNames []string   `json:"categoryNames"`
}

func createVocabulary(w http.ResponseWriter, r *http.Request, vocabularyEntity VocabularyEntity.Entity) error {
	var request CreateVocabularyWithCategoriesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	vocabulary := VocabularyEntity.Vocabulary{
		Words:        request.Vocabulary.Words,
		Translation:  request.Vocabulary.Translation,
		UsedInPhrase: request.Vocabulary.UsedInPhrase,
		Explanation:  request.Vocabulary.Explanation,
	}

	voc, err := vocabularyEntity.CreateWithCategories(&vocabulary, request.CategoryNames)
	if err != nil {
		return err
	}

	createVocabularyWithCategories := CreateVocabularyWithCategories{}

	createVocabularyWithCategories.Id = voc.Id
	createVocabularyWithCategories.Words = voc.Words
	createVocabularyWithCategories.Translation = voc.Translation
	createVocabularyWithCategories.UsedInPhrase = voc.UsedInPhrase
	createVocabularyWithCategories.Explanation = voc.Explanation
	for _, categoryEntity := range voc.Categories {
		createVocabularyWithCategories.Categories = append(createVocabularyWithCategories.Categories, Category{
			Id:   categoryEntity.Id,
			Name: categoryEntity.Name,
		})
	}

	return writeJson(w, http.StatusOK, createVocabularyWithCategories)
}

func (o *ApiServer) ListenAndServe() error {
	if err := http.ListenAndServe(o.ListenPort, o.Router); err != nil {
		logger.GetLogger().LogError("couldn't up the server", err)
	}

	return nil
}

func test(w http.ResponseWriter, _ *http.Request) {
	if err := writeJson(w, http.StatusOK, `{"hello": "world"}`); err != nil {
		logger.GetLogger().LogError("couldn't encode in json", err)
	}
}

func writeJson(w http.ResponseWriter, status int, obj interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		return err
	}
	return nil
}
