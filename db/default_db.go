package db

import (
	"CRUD-service/pkg/entities"
	"CRUD-service/utils"
	"encoding/json"
	"os"
	"path"
	"strconv"
)

type DefaultDB struct {
	conter int
	data   []entities.DefaultEntity
}

func NewDefaultDB() *DefaultDB {
	db := new(DefaultDB)
	err := db.load()
	if err != nil {
		return nil
	}

	return db
}

func (db *DefaultDB) load() error {
	projectDir, err := utils.GetProjectRoot()
	filePath := "db/datafiles/default.json"
	data, err := os.ReadFile(path.Join(projectDir, filePath))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &db.data)
	if err != nil {
		return err
	}

	db.conter, err = strconv.Atoi(db.data[len(db.data)-1].ID)

	return nil
}

func (db *DefaultDB) DataToByteSlice() ([]byte, error) {
	data, err := json.Marshal(db.data)
	if err != nil {
		return data, err
	}

	return data, nil
}
