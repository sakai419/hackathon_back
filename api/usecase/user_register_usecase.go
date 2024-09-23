package usecase

import (
	"crypto/rand"
	"database/sql"
	"api/dao"
	"api/model"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/oklog/ulid"
)

func GetData(body []byte) (model.UserJsonForHTTPPost, error) {
	data := model.UserJsonForHTTPPost{}
	err := json.Unmarshal(body, &data)
	return data, err
}

func CheckUser(data model.UserJsonForHTTPPost) (error) {
	if len(data.Name) == 0 {
		return errors.New("data.Name is empty")
	} else if len(data.Name) > 50 {
		return errors.New("data.Name is too long")
	}

	if data.Age < 20 || data.Age > 80 {
		return errors.New("data.Age is out of range")
	}
	return nil
}

func RegisterUser(db *sql.DB, data model.UserJsonForHTTPPost) (*ulid.ULID, error) {
	tx, err := dao.GetTX(db)
	if err != nil {
		return nil, fmt.Errorf("fail: dao.GetTX, %v", err)
	}

    defer func() {
        if err != nil {
            if rbErr := dao.RollbackTX(tx); rbErr != nil {
                err = fmt.Errorf("rollback failed: %v, original error: %w", rbErr, err)
            }
        }
    }()

	id, err := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("fail: ulid.New, %v", err)
	}

	if err := dao.InsertUser(tx, id.String(), data.Name, data.Age); err != nil {
		return nil, fmt.Errorf("fail: dao.InserUser, %v", err)
	}

	if err := dao.CommitTX(tx); err != nil {
		return nil, fmt.Errorf("fail: dao.CommitTX, %v", err)
	}

	return &id, nil
}

func MakeResponse(id *ulid.ULID) ([]byte, error) {
	return json.Marshal(model.UserResForHTTPPost{
		Id: id.String(),
	})
}