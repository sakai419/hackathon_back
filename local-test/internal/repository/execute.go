package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
	"os"
)

func (r *Repository) ExecuteCCode(ctx context.Context, content string) (*model.ExecuteResult, error) {
	// Get GCC server URL
	gccServerURL := os.Getenv("GCC_SERVER_URL")
	if gccServerURL == "" {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrInvalidInput{
				Message: "GCC_SERVER_URL is not set",
			},
		)
	}

	// Create request body
	reqBody := map[string]string{
		"code": content,
	}
	jsonBody, _ := json.Marshal(reqBody)

	// Post to GCC server
	req, err := http.NewRequest("POST", gccServerURL+"/compile", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create request to GCC server",
				Err:       err,
			},
		)
	}

	// Set headers including Referer
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", os.Getenv("EXECUTE_PATH"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "post to GCC server",
				Err:       err,
			},
		)
	}
	defer resp.Body.Close()


	// Decode response
	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "decode response from GCC server",
				Err:       err,
			},
		)
	}

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "post to GCC server",
				Err:       errors.New("status code is not 200: " + result["message"].(string)),
			},
		)
	}

	// Create ExecuteResult
	ret := &model.ExecuteResult{
		Status:  result["status"].(string),
	}
	if _, ok := result["output"]; ok {
		temp := result["output"].(string)
		ret.Output = &temp
	}
	if _, ok := result["message"]; ok {
		temp := result["message"].(string)
		ret.Message = &temp
	}

	return ret, nil
}