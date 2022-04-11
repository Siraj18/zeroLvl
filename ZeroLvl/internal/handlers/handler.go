package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/siraj18/zeroLvl/internal/ports"
	"github.com/siraj18/zeroLvl/internal/repositories/modelsrepo"
	"github.com/sirupsen/logrus"
)

type handler struct {
	router    *chi.Mux
	logger    *logrus.Logger
	modelsSrv ports.ModelsService
}

func NewHandler(modelSrv ports.ModelsService) *handler {
	return &handler{
		router:    chi.NewRouter(),
		logger:    logrus.New(),
		modelsSrv: modelSrv,
	}
}

func (h *handler) test(w http.ResponseWriter, r *http.Request) {
	models, err := h.modelsSrv.GetModelsFromDb()
	if err != nil {
		fmt.Println(err.Error())
	}

	json.NewEncoder(w).Encode(models)

}

func (h *handler) getById(w http.ResponseWriter, r *http.Request) {
	// Для того чтобы мог делать запросы с js
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	id := chi.URLParam(r, "id")

	model, err := h.modelsSrv.GetModelFromCache(id)
	if err != nil {

		if err == modelsrepo.ErrorNotFound {
			w.WriteHeader(http.StatusNotFound)
			h.logger.Error(err.Error())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error(err.Error())
		return
	}

	json.NewEncoder(w).Encode(model)
}

func (handler *handler) InitRoutes() *chi.Mux {

	handler.router.Get("/getById/{id}", handler.getById)
	handler.router.Get("/test/", handler.test)

	return handler.router
}
