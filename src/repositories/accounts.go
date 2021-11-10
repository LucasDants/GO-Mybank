package repositories

import (
	"myBank/src/models"
	"database/sql"
)

type accounts struct {
	db *sql.DB
}

// Cria um repositorio de usuarios
func NewAccountsRepository(db *sql.DB) *accounts {
	return &accounts{db}
}

func (repository accounts) Create(account models.Account) (uint64, error) {
	statement, err := repository.db.Prepare("insert into accounts (name, password, currency, amount) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(account.Name, account.Password, account.Currency, 0)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository accounts) SearchByName(name string) (models.Account, error) {
	row, err := repository.db.Query("select id, password from accounts where name = ?", name)
	if err != nil {
		return models.Account{}, err
	}
	defer row.Close()

	var account models.Account

	if row.Next() {
		if err = row.Scan(
			&account.ID,
			&account.Password,
		); err != nil {
			return models.Account{}, err
		}
	}
	return account, nil
}

func (repository accounts) SearchByID(ID uint64) (models.Account, error) {
	row, err := repository.db.Query("select id, name, amount, currency, createdAt from accounts where id = ?", ID)
	if err != nil {
		return models.Account{}, err
	}
	defer row.Close()

	var account models.Account

	if row.Next() {
		if err = row.Scan(
			&account.ID,
			&account.Name,
			&account.Amount,
			&account.Currency,
			&account.CreatedAt,
		); err != nil {
			return models.Account{}, err
		}
	}
	return account, nil
}