package repositories

import (
	"myBank/src/models"
	"database/sql"
	"context"
)

type operations struct {
	db *sql.DB
}

// Cria um repositorio de usuarios
func NewOperationsRepository(db *sql.DB) *operations {
	return &operations{db}
}

func (repository operations) Deposit(operation models.Operation) (uint64, error) {

	ctx := context.Background()
	tx, err := repository.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	statement, err := repository.db.PrepareContext(ctx, "insert into operations(origin, destiny, exchangeCurrency, type) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(operation.Origin, operation.Destiny, operation.ExchangeCurrency, "deposit")
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	statement2, err := repository.db.PrepareContext(ctx, "update accounts set amount = ? + amount where id = ?")
	if err != nil {
		return 0, err
	}
	defer statement2.Close()

	_, err = statement2.Exec(operation.ExchangeCurrency, operation.Destiny)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository operations) Transfer(operation models.Operation) (uint64, error) {
	ctx := context.Background()
	tx, err := repository.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	statement, err := repository.db.PrepareContext(ctx, "insert into operations(origin, destiny, exchangeCurrency, type, originCurrency, destinyCurrency, currencyTransformed) values (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(operation.Origin, operation.Destiny, operation.ExchangeCurrency, "transfer", operation.OriginCurrency, operation.DestinyCurrency, operation.CurrencyTransformed)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	statement2, err := repository.db.PrepareContext(ctx, "update accounts set amount = amount - ? where id = ?")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer statement2.Close()

	_, err = statement2.Exec(operation.ExchangeCurrency, operation.Origin)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	statement3, err := repository.db.PrepareContext(ctx, "update accounts set amount = amount + ? where id = ?")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer statement3.Close()

	_, err = statement3.Exec(operation.CurrencyTransformed, operation.Destiny)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository operations) Withdraw(operation models.Operation) (uint64, error) {
	ctx := context.Background()
	tx, err := repository.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	statement, err := repository.db.PrepareContext(ctx, "insert into operations(origin, destiny, exchangeCurrency, type) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(operation.Origin, operation.Destiny, operation.ExchangeCurrency, "withdraw")
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	statement2, err := repository.db.PrepareContext(ctx, "update accounts set amount = amount - ? where id = ?")
	if err != nil {
		return 0, err
	}
	defer statement2.Close()

	_, err = statement2.Exec(operation.ExchangeCurrency, operation.Origin)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository operations) Search(accountID uint64) ([]models.Operation, error) {
	rows, err := repository.db.Query("select id, origin, destiny, exchangeCurrency, type, createdAt from operations where origin = ? or destiny = ?", accountID, accountID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var operations []models.Operation

	for rows.Next() {
		var operation models.Operation

		if err = rows.Scan(
			&operation.ID,
			&operation.Origin,
			&operation.Destiny,
			&operation.ExchangeCurrency,
			&operation.Type,
			&operation.CreatedAt,
		); err != nil {
			return nil, err
		}
		operations = append(operations, operation)
	}

	return operations, nil
}