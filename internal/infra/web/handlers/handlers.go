package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
	"github.com/kameikay/url-shortener/internal/dtos"
	"github.com/kameikay/url-shortener/internal/infra/repository"
	"github.com/kameikay/url-shortener/internal/usecases"
	"github.com/kameikay/url-shortener/utils"
)

type Handler struct {
	repository repository.RepositoryInterface
}

func NewHandler(repository repository.RepositoryInterface) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) ReturnCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    http.StatusText(http.StatusMethodNotAllowed),
			Success:    false,
		})
		return
	}

	var input dtos.CreateCodeInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid input",
			Success:    false,
		})
		return
	}

	generateCodeUseCase := usecases.NewGenerateCodeUseCase(h.repository)
	code, err := generateCodeUseCase.Execute(r.Context(), input)
	if err != nil {
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Success:    false,
		})
		return
	}

	response := struct {
		Code string `json:"code"`
	}{
		Code: code,
	}

	utils.JsonResponse(w, utils.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Success:    true,
		Data:       response,
	})
}

func (h *Handler) RedirectToUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    http.StatusText(http.StatusMethodNotAllowed),
			Success:    false,
		})
		return
	}

	code := chi.URLParam(r, "code")
	if code == "" {
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "code is required",
			Data:       nil,
			Success:    false,
		})
		return
	}

	getUrlUseCase := usecases.NewGetUrlUseCase(h.repository)
	url, err := getUrlUseCase.Execute(r.Context(), code)
	if err != nil {
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Success:    false,
		})
		return
	}

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
