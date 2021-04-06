package rest

import (
	"encoding/json"
	"fmt"

	. "internal/entities"
	"internal/persistence"
	"net/http"

	"github.com/gorilla/mux"
)

var languageDAO = persistence.GetLanguageDAO()

func getAllLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func getAllLanguages")
	list := languageDAO.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func getLanguageById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func getLanguageById")
	vars := mux.Vars(r)
	searchCode := vars["code"]
	foundLanguage := languageDAO.Find(searchCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundLanguage)
}

func deleteLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func deleteLanguage")
	vars := mux.Vars(r)
	searchCode := vars["code"]
	for _, s := range studentDAO.FindAll() {
		if s.LanguageCode == searchCode {
			fmt.Printf("Suppression impossible puisqu'au moins un Student utilise ce language")
			return
		}
	}
	isDeleted := languageDAO.Delete(searchCode)
	if !isDeleted {
		fmt.Printf("Le language n'a pas pu être supprimé")
		return
	}
}

func createLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func createLanguage")
	var lang Language

	err := json.NewDecoder(r.Body).Decode(&lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	alreadyExist := languageDAO.Exists(lang.Code)
	if alreadyExist {
		fmt.Printf("Le language existe déja")
		return
	}

	isCreated := languageDAO.Create(&lang)
	if !isCreated {
		fmt.Printf("Erreur de création du language")
		return
	}
}

func updateLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func updateLanguage")
	var lang Language

	err := json.NewDecoder(r.Body).Decode(&lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	alreadyExist := languageDAO.Exists(lang.Code)
	if !alreadyExist {
		fmt.Printf("Le language n'existe pas")
		return
	}

	isUpdated := languageDAO.Update(&lang)
	if !isUpdated {
		fmt.Printf("Erreur pendant l'update")
		return
	}
}

func LanguageHandlers(router *mux.Router) {
	router.HandleFunc("/languages", getAllLanguages).Methods("GET")
	router.HandleFunc("/languages/{code}", getLanguageById).Methods("GET")
	router.HandleFunc("/languages/{code}", deleteLanguage).Methods("DELETE")
	router.HandleFunc("/languages", createLanguage).Methods("POST")
	router.HandleFunc("/languages", updateLanguage).Methods("PUT")
}
