package modelsrepo

import (
	"github.com/jmoiron/sqlx"
	"github.com/siraj18/zeroLvl/internal/models"
)

type PostgreRepository struct {
	db *sqlx.DB
}

func NewPostgreRepository(db *sqlx.DB) (*PostgreRepository, error) {
	rep := &PostgreRepository{db}

	if err := rep.init(); err != nil {
		return nil, err
	}

	return rep, nil
}

func (rep *PostgreRepository) GetModels() ([]*models.Model, error) {
	var models []*models.Model

	err := rep.db.Select(&models, sqlGetModel)
	if err != nil {
		return nil, err
	}

	for _, model := range models {
		err = rep.db.Get(&model.Delivery, sqlGetDelivery, model.OrderUID)
		if err != nil {
			return nil, err
		}

		err = rep.db.Get(&model.Payment, sqlGetPayment, model.OrderUID)
		if err != nil {
			return nil, err
		}

		err = rep.db.Select(&model.Items, sqlGetItems, model.OrderUID)

		if err != nil {
			return nil, err
		}
	}

	return models, nil
}

func (rep *PostgreRepository) AddModel(model *models.Model) error {
	tx, err := rep.db.Begin()

	if err != nil {
		return err
	}
	defer tx.Rollback()

	// add to deliveries
	_, err = tx.Exec(sqlAddToDeliveries,
		model.OrderUID,
		model.Delivery.Name,
		model.Delivery.Phone,
		model.Delivery.Zip,
		model.Delivery.City,
		model.Delivery.Address,
		model.Delivery.Region,
		model.Delivery.Email)

	if err != nil {
		return err
	}

	_, err = tx.Exec(sqlAddToPayments,
		model.OrderUID,
		model.Payment.Transaction,
		model.Payment.RequestID,
		model.Payment.Currency,
		model.Payment.Provider,
		model.Payment.Amount,
		model.Payment.PaymentDt,
		model.Payment.Bank,
		model.Payment.DeliveryCost,
		model.Payment.GoodsTotal,
		model.Payment.CustomFee)

	if err != nil {
		return err
	}

	for _, item := range model.Items {
		_, err = tx.Exec(sqlAddToItems,
			model.OrderUID,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status)

		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(sqlAddToModels,
		model.OrderUID,
		model.TrackNumber,
		model.Entry,
		model.Locale,
		model.InternalSignature,
		model.CustomerID,
		model.DeliveryService,
		model.Shardkey,
		model.SmID,
		model.DateCreated,
		model.OofShard)

	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (rep *PostgreRepository) init() error {
	_, err := rep.db.Exec(initSchema)

	if err != nil {
		return err
	}

	return nil
}
