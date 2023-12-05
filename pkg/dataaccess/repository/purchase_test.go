package repository

import (
	"testing"
	"time"

	"github.com/heronvitor/pkg/entities"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestPurchaseRepository_SavePurchase(t *testing.T) {
	t.Run("should create purchase", func(t *testing.T) {
		db := createDB(t)
		defer func() { db.Close() }()

		repository := PurchaseRepository{DB: db}

		goErr := repository.SavePurchase(entities.Purchase{
			ID:              "5e08af1b-b9c3-4049-9a27-b000c86340d1",
			Description:     "purchase description",
			Amount:          5.4,
			TransactionDate: time.Date(2000, 1, 2, 0, 0, 3, 4, time.UTC),
		})

		assert.NoError(t, goErr)

		// check db
		fields := make([]string, 4, 4)

		assert.NoError(t, db.QueryRow("SELECT * FROM purchase").
			Scan(&fields[0], &fields[1], &fields[2], &fields[3]))

		wantFields := []string{"5e08af1b-b9c3-4049-9a27-b000c86340d1", "purchase description", "5.40", "2000-01-02T00:00:00Z"}
		assert.Equal(t, wantFields, fields)
	})
}

func TestPurchaseRepository_GetPurchaseByID(t *testing.T) {
	t.Run("should return null if not found", func(t *testing.T) {
		db := createDB(t)
		defer func() { db.Close() }()

		repository := PurchaseRepository{DB: db}

		gotPurchase, goErr := repository.GetPurchaseByID("f0006fef-233b-48c2-a066-4da2062b2f56")

		assert.NoError(t, goErr)
		assert.Nil(t, gotPurchase)
	})

	t.Run("should get purchase", func(t *testing.T) {
		db := createDB(t)
		defer func() { db.Close() }()

		// prepare db
		query := "INSERT INTO purchase (id, description, amount, transaction_date) VALUES ($1,$2,$3,$4)"
		_, err := db.Exec(query, "f0006fef-233b-48c2-a066-4da2062b2f56", "desc", "5.35", "2032-01-02T00:00:00Z")

		assert.NoError(t, err)

		// actual test
		wantPurchase := &entities.Purchase{
			ID:              "f0006fef-233b-48c2-a066-4da2062b2f56",
			Description:     "desc",
			Amount:          5.35,
			TransactionDate: time.Date(2032, time.January, 2, 0, 0, 0, 0, time.UTC),
		}

		repository := PurchaseRepository{DB: db}

		gotPurchase, goErr := repository.GetPurchaseByID("f0006fef-233b-48c2-a066-4da2062b2f56")

		assert.NoError(t, goErr)
		assert.Equal(t, wantPurchase, gotPurchase)
	})
}
