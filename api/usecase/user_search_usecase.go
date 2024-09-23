package usecase

import (
	"database/sql"
	"api/model"
	"errors"
	"net/url"
)

func CheckName(queryParams url.Values) (string, error) {
	name := queryParams.Get("name")

	if name == "" {
		return "", errors.New("name is empty")
	}

	return name, nil
}

func GetUsers(rows *sql.Rows) ([]model.UserResForHTTPGet, error) {
	users := make([]model.UserResForHTTPGet, 0)
	for rows.Next() {
		var u model.UserResForHTTPGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}