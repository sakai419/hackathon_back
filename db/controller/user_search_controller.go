package controller

import (
	"database/sql"
	"db/dao"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func SearchUserControll(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	queryParams := r.URL.Query()

	name, err := usecase.CheckName(queryParams)

	if err != nil {
		log.Printf("fail: usecase.CheckName, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := dao.SelectUserByName(db, name)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users, err := usecase.GetUsers(rows)
	if err != nil {
		log.Printf("fail: usecase.GetUsers, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)

	if err != nil {
		log.Printf("fail: w.Write, %v\n", err)
		return
	}
}