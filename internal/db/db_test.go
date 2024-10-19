package db

import (
    "testing"
    "github.com/SashaZezulinsky/block-scout/internal/models"
)

func TestSaveTransfer(t *testing.T) {
    db, err := NewDatabase(":memory:")
    if err != nil {
        t.Fatalf("Failed to create database: %v", err)
    }

    transfer := models.Transfer{Sender: "0xSender", Recipient: "0xRecipient", Value: "1000"}
    err = db.SaveTransfer(transfer)
    if err != nil {
        t.Errorf("Failed to save transfer: %v", err)
    }
}
