package repository

import (
	"crm_admin/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ClientsPG struct {
	db *sqlx.DB
}

func (c ClientsPG) GetClients() ([]domain.Client, error) {
	query := fmt.Sprintf(`	
			SELECT * FROM clients
	`)

	var clients []domain.Client
	err := c.db.Select(&clients, query)

	return clients, err
}

func (c ClientsPG) GetClientByID(clientID int) (domain.Client, error) {
	query := fmt.Sprintf(`	
			SELECT * FROM clients WHERE client_id = $1
	`)

	var client domain.Client
	err := c.db.Get(&client, query, clientID)

	return client, err
}

func (c ClientsPG) CreateClient(client domain.Client) (int, error) {
	query := fmt.Sprintf(`
			INSERT INTO clients (client_name)
      VALUES ($1)
			RETURNING client_id
	`)

	var id int
	err := c.db.QueryRow(query, client.Name).Scan(&id)

	return id, err
}

func (c ClientsPG) UpdateClient(client domain.Client) error {
	query := fmt.Sprintf(`
		 UPDATE clients
		    SET client_name = $1
		  WHERE client_id = $2
	`)

	_, err := c.db.Exec(query, client.Name, client.ID)
	return err
}

func NewClientsPG(db *sqlx.DB) *ClientsPG {
	return &ClientsPG{db: db}
}
