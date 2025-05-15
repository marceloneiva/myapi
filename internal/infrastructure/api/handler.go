package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/marceloneiva/myapi/internal/aplication/usecase"
	"github.com/marceloneiva/myapi/internal/domain/entity"
)

type Handler struct {
	usecase *usecase.ConvertCurrencyUseCase
}

func NewHandler(uc *usecase.ConvertCurrencyUseCase) *Handler {
	return &Handler{usecase: uc}
}

func (h *Handler) ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	from := entity.Currency(r.URL.Query().Get("from"))
	to := entity.Currency(r.URL.Query().Get("to"))
	amountStr := r.URL.Query().Get("amount")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Valor inválido", http.StatusBadRequest)
		return
	}

	result, err := h.usecase.Execute(from, to, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{"result": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
