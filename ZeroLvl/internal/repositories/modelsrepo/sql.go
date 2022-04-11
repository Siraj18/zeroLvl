package modelsrepo

const initSchema = `
				CREATE TABLE IF NOT EXISTS models 
				(
					order_uid  	       TEXT PRIMARY KEY,
					track_number       TEXT,
					entry              TEXT,
					locale             TEXT,
					internal_signature TEXT,
					customer_id        TEXT,
					delivery_service   TEXT,
					shard_key          TEXT,
					sm_id              BIGINT,
					date_created       TIMESTAMP,
					oof_shard          TEXT
				);

				CREATE TABLE IF NOT EXISTS deliveries 
				(
					order_uid TEXT PRIMARY KEY,
					name  	  TEXT,
					phone     TEXT,
					zip       TEXT,
					city      TEXT,
					address   TEXT,
					region    TEXT,
					email     TEXT
				);

				CREATE TABLE IF NOT EXISTS payments 
				(
					order_uid     TEXT PRIMARY KEY,
					transaction   TEXT,
					request_id    TEXT,
					currency      TEXT,
					provider      TEXT,
					amount        BIGINT,
					payment_dt    BIGINT,
					bank          TEXT,
					delivery_cost BIGINT,
					goods_total    BIGINT,
					custom_fee    BIGINT
				);

				CREATE TABLE IF NOT EXISTS items 
				(
					item_id      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
					order_uid    TEXT,
					chrt_id      BIGINT,
					track_number TEXT,
					price        BIGINT,
					rid          TEXT,
					name         TEXT,
					sale         BIGINT,
					size         TEXT,
					total_price  BIGINT,
					nm_id        BIGINT,
					brand        TEXT,
					status       BIGINT
				);
`

const sqlAddToModels = `INSERT INTO models VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

const sqlAddToDeliveries = `INSERT INTO deliveries VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

const sqlAddToPayments = `INSERT INTO payments VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

const sqlAddToItems = `INSERT INTO items(order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
					   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`

const sqlGetModel = `
				SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard
				FROM models
				`

const sqlGetDelivery = `
				SELECT name, phone, zip, city, address, region, email
				FROM deliveries
				WHERE order_uid = $1;
`

const sqlGetPayment = `
				SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
				FROM payments
				WHERE order_uid = $1;
`

const sqlGetItems = `
				SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
				FROM items
				WHERE order_uid = $1;
`
