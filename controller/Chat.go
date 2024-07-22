package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/anakilang-ai/backend/models"
	"github.com/anakilang-ai/backend/modules"
	"github.com/anakilang-ai/backend/utils"
	"github.com/go-resty/resty/v2"
)

func Chat(respw http.ResponseWriter, req *http.Request, tokenmodel string) {
	var chat models.AIRequest

	err := json.NewDecoder(req.Body).Decode(&chat)
	if err != nil {
		utils.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	if chat.Prompt == "" {
		utils.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}

	client := resty.New()

	// Hugging Face API URL dan token
	apiUrl := modules.GetEnv("HUGGINGFACE_API_KEY")
	apiToken := "Bearer " + tokenmodel

	var response *resty.Response
	var retryCount int
	maxRetries := 5
	retryDelay := 20 * time.Second

	parsedURL, err := url.Parse(apiUrl)

	if err != nil {
		utils.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing URL model hugging face"+err.Error())
		return
	}

	segments := strings.Split(parsedURL.Path, "/")

	modelName := strings.Join(segments[2:], "/")

	// Request ke Hugging Face API
	for retryCount < maxRetries {
		response, err = client.R().
			SetHeader("Authorization", apiToken).
			SetHeader("Content-Type", "application/json").
			SetBody(`{"inputs": "` + chat.Prompt + `"}`).
			Post(apiUrl)

		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}

		if response.StatusCode() == http.StatusOK {
			break
		} else {
			var errorResponse map[string]interface{}
			err = json.Unmarshal(response.Body(), &errorResponse)
			if err == nil && errorResponse["error"] == "Model "+modelName+" is currently loading" {
				retryCount++
				time.Sleep(retryDelay)
				continue
			}
			utils.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
			return
		}
	}

	if response.StatusCode() != 200 {
		utils.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Errorr", "error from Hugging Face API "+string(response.Body()))
		return
	}

	var data []map[string]interface{}

	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		utils.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing response body "+err.Error())
		return
	}

	if len(data) > 0 {
		generatedText, ok := data[0]["generated_text"].(string)
		if !ok {
			utils.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error extracting generated text")
			return
		}
		utils.WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText})
	} else {
		utils.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: response")
	}
}
