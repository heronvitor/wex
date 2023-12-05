package repository

import (
	"database/sql"

	"github.com/heronvitor/pkg/entities"
	_ "github.com/lib/pq"
)

type PurchaseRepository struct {
	DB *sql.DB
}

func (r PurchaseRepository) GetPurchaseByID(id string) (*entities.Purchase, error) {
	purchase := &entities.Purchase{}
	row := r.DB.QueryRow("SELECT id, description, amount, transaction_date FROM purchase WHERE id = $1", id)

	if err := row.Scan(&purchase.ID, &purchase.Description, &purchase.Amount, &purchase.TransactionDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	purchase.TransactionDate = purchase.TransactionDate.UTC()
	return purchase, nil
}

func (r PurchaseRepository) SavePurchase(purchase entities.Purchase) error {
	query := "INSERT INTO purchase (id, description, amount, transaction_date) VALUES ($1,$2,$3,$4)"
	_, err := r.DB.Exec(query, purchase.ID, purchase.Description, purchase.Amount, purchase.TransactionDate)
	return err
}
