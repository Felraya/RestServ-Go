package rest

import (
	"encoding/json"
	"fmt"
	. "internal/entities"
	"internal/persistence"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var studentDAO = persistence.GetStudentDAO()

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func getAllStudents")
	list := studentDAO.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func getStudentById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func getStudentById")
	vars := mux.Vars(r)
	searchId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("Erreur : %s", err)
	}
	foundStudent := studentDAO.Find(searchId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundStudent)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func deleteStudent")
	vars := mux.Vars(r)
	searchId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("Erreur : %s", err)
	}
	isDeleted := studentDAO.Delete(searchId)
	if !isDeleted {
		fmt.Printf("Le student n'a pas pu être supprimé")
		return
	}
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func createStudent")
	var stu Student

	err := json.NewDecoder(r.Body).Decode(&stu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	alreadyExist := studentDAO.Exists(stu.Id)
	if alreadyExist {
		fmt.Printf("Le student existe déja")
		return
	}

	isCreated := studentDAO.Create(&stu)
	if !isCreated {
		fmt.Printf("Erreur de création du student")
		return
	}
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func updateStudent")
	var stu Student

	err := json.NewDecoder(r.Body).Decode(&stu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	alreadyExist := studentDAO.Exists(stu.Id)
	if !alreadyExist {
		fmt.Printf("Le student n'existe pas")
		return
	}

	isUpdated := studentDAO.Update(&stu)
	if !isUpdated {
		fmt.Printf("Erreur pendant l'update")
		return
	}
}

func StudentHandlers(router *mux.Router) {
	router.HandleFunc("/students", getAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", getStudentById).Methods("GET")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	router.HandleFunc("/students", createStudent).Methods("POST")
	router.HandleFunc("/students", updateStudent).Methods("PUT")
}
