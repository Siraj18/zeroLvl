package stancon

import (
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func NewStanConnection(clusterId, clientId, natsUrl string) *stan.Conn {

	sc, err := stan.Connect(clusterId, clientId, stan.NatsURL(natsUrl))
	if err != nil {
		logrus.Fatal("failed to connect nats-streaming server:", err.Error())
	}
	return &sc
}
