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
	o.Router.HandleFunc("/createVocabulary", o.makeHttpFunc(createVocabulary))
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
		vocabularyRepository := txRepositoryFactory.GetVocabularyRepository()
		vocabularyEntity := VocabularyEntity.New(vocabularyRepository)
		if err := f(w, r, vocabularyEntity); err != nil {
			txRepositoryFactory.RollbackTransaction()
			if err = writeJson(w, http.StatusInternalServerError, err); err != nil {
				logger.GetLogger().LogError("endpoint error", err)
				return
			}
			return
		}
		txRepositoryFactory.CommitTransaction()
	}
}

func createVocabulary(w http.ResponseWriter, r *http.Request, vocabularyEntity VocabularyEntity.Entity) error {
	var request CreateVocabulary
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	vocabulary := VocabularyEntity.Vocabulary{
		Words:        &request.Words,
		Translation:  &request.Translation,
		UsedInPhrase: &request.UseInPhrase,
		Explanation:  &request.Explanation,
	}

	categories := []string{"hola", "mundo"}

	voc, err := vocabularyEntity.CreateWithCategories(&vocabulary, categories)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, voc)
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
