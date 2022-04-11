package handlers

import (
	"encoding/json"

	"github.com/nats-io/stan.go"
	"github.com/siraj18/zeroLvl/internal/models"
	"github.com/siraj18/zeroLvl/internal/ports"
	"github.com/sirupsen/logrus"
)

type stanHandler struct {
	stanCon      stan.Conn
	modelService ports.ModelsService
	logger       *logrus.Logger
}

func NewStanHandler(stanCon stan.Conn, modelsSrv ports.ModelsService) *stanHandler {
	return &stanHandler{
		stanCon:      stanCon,
		logger:       logrus.New(),
		modelService: modelsSrv,
	}
}

//TODO исправить обработку ошибок

func (s *stanHandler) Subscribe(channel, durableName string) {
	_, err := s.stanCon.Subscribe(channel, func(m *stan.Msg) {
		model := &models.Model{}
		err := json.Unmarshal(m.Data, &model)

		if err != nil {
			s.logger.Error("error unmarshaling data:", err.Error())
			return
		}

		err = s.modelService.AddModelToDb(model)

		if err != nil {
			s.logger.Error("error writing to database:", err.Error())
			return
		}

		s.modelService.SetModelToCache(model)

	}, stan.DurableName(durableName))

	if err != nil {
		s.logger.Fatal("error when subscribe to the channel:", err.Error())
	}

}
