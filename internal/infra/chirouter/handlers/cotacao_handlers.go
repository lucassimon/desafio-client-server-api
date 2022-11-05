package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/lucassimon/desafio-client-server-api/internal/dto"
	"github.com/lucassimon/desafio-client-server-api/internal/entity"
	"github.com/lucassimon/desafio-client-server-api/internal/infra/database"
)

type DollarExchangeHandler struct {
	DollarExchangeDB database.DollarExchangeInterface
}

func NewDollarExchangeHandler(db database.DollarExchangeInterface) *DollarExchangeHandler {
	return &DollarExchangeHandler{
		DollarExchangeDB: db,
	}
}

func (h *DollarExchangeHandler) GetCotacao(w http.ResponseWriter, r *http.Request) {
	// Make request
	// sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	dollar_real, err := h.makeRequest(ctx)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	// Save in database
	// timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
	usd_brl_et := entity.NewUsdBrl(dollar_real)
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	err = h.DollarExchangeDB.Create(ctx, usd_brl_et)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
	// Return the bid
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dollar_real)
}

func (h *DollarExchangeHandler) prepareRequest(ctx context.Context) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		"https://economia.awesomeapi.com.br/json/last/USD-BRL",
		nil,
	)

	return req, err
}

func (h *DollarExchangeHandler) request(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, err := client.Do(req)

	return response, err
}

func (h *DollarExchangeHandler) makeRequest(ctx context.Context) (*dto.CreateDollarInput, error) {
	req, err := h.prepareRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := h.request(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var c dto.CreateDollarInput
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
