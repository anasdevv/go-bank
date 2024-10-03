package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter , status int, val any) error{
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(val)
}



type APIServer struct {
	listenAddr string
	store Storage
}

type APIFunc func(http.ResponseWriter,*http.Request) error


type APIError struct {
	Error  string
}

func makeHTTPHandlerFunc(f APIFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r * http.Request)  {
		if err := f(w,r); err != nil {
			// handle the error
			WriteJson(w,http.StatusBadRequest,APIError{
				Error:  err.Error(),
			})
		}
	}
}

func NewAPIServer(listenAddr string , store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}	
}

func (s * APIServer) Run(){
	router := mux.NewRouter()
	router.HandleFunc("/account",makeHTTPHandlerFunc(s.handleAccount))
	router.HandleFunc("/account/{id}",makeHTTPHandlerFunc(s.handleAccount))

	log.Println("HTTP server running on port " , s.listenAddr)
	http.ListenAndServe(s.listenAddr,router)
}


func (s * APIServer) handleAccount(w http.ResponseWriter , r *http.Request) error {

	if r.Method == "GET" {
		println("Get req ... ")
		return s.handleGetAccount(w,r)
	}
	if r.Method == "POST" {
		println("POST req ... ")
		return s.handleCreateAccount(w,r)
	}
	if r.Method == "DELETE" {
		println("Delete req ... ")
		return s.handleDeleteAccount(w,r)
	}

	return fmt.Errorf("method not supported %s" , r.Method)
}



func (s * APIServer) handleGetAccount(w http.ResponseWriter , r *http.Request) error {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		return err
	}
	account,err := s.store.GetAccountById(int(idInt))
	if err != nil {
		fmt.Printf("Error in getting account %v \n",err)
		return WriteJson(w,http.StatusNotFound,err)
	}
	return WriteJson(w,http.StatusOK,account)
}
func (s * APIServer) handleCreateAccount(w http.ResponseWriter , r *http.Request) error {
	createAccReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccReq); err != nil {
		return err
	}
	account := NewAccount(createAccReq.FirstName,createAccReq.LastName)
	if err:= s.store.CreateAccount(account); err != nil {
		return err
	}
	WriteJson(w,http.StatusOK , account)
	return nil
}
func (s * APIServer) handleDeleteAccount(w http.ResponseWriter , r *http.Request) error {
	id := mux.Vars(r)["id"]
	println(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		return err
	}
	if err := s.store.DeleteAccount(idInt); err != nil {
		fmt.Printf("Error in deleting account: %v\n", err)
		return err
	}
	return WriteJson(w,http.StatusOK,"Deleted")
}
func (s * APIServer) handleTransfer(w http.ResponseWriter , r *http.Request) error {
	return nil
}


