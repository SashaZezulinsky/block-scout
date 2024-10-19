package db

import (
  "database/sql"
  "log"
  "github.com/SashaZezulinsky/block-scout/internal/models"

  _ "github.com/mattn/go-sqlite3"
)

type Database struct {
  db *sql.DB
}

func NewDatabase(filepath string) (*Database, error) {
  db, err := sql.Open("sqlite3", filepath)
  if err != nil {
    return nil, err
  }

  query := `CREATE TABLE IF NOT EXISTS transfers (id INTEGER PRIMARY KEY, sender TEXT, recipient TEXT, value TEXT)`
  _, err = db.Exec(query)
  if err != nil {
    return nil, err
  }

  return &Database{db: db}, nil
}

func (d *Database) SaveTransfer(transfer models.Transfer) error {
  query := `INSERT INTO transfers (sender, recipient, value) VALUES (?, ?, ?)`
  _, err := d.db.Exec(query, transfer.Sender, transfer.Recipient, transfer.Value)
  if err != nil {
    log.Printf("Failed to save transfer: %v", err)
  }
  return err
}
