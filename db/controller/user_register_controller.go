package controller

import (
	"database/sql"
	"db/usecase"
	"io"
	"log"
	"net/http"
)

func RegisterUserControll(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fail: ioutil.ReadAll(r.Body), %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := usecase.GetData(body)
	if err != nil {
		log.Printf("fail: usecase.GetData, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := usecase.CheckUser(data); err != nil {
		log.Printf("fail: usecase.CheckUser, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := usecase.RegisterUser(db, data)

	if err != nil {
		log.Printf("fail: usecase.RegisterUser, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


	bytes, err := usecase.MakeResponse(id)

	if err != nil {
		log.Printf("fail: usecase.MakeResponse, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(bytes); err != nil {
		log.Printf("fail: w.Write, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}