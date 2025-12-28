package db

import (
	"CRUD-service/pkg/entities"
	"encoding/json"
	"os"
)

type DefaultDB struct {
	data []entities.DefaultEntity
}

func (db DefaultDB) Load() error {
	fileName := "default.json"
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &db.data)
	if err != nil {
		return err
	}

	return nil
}
