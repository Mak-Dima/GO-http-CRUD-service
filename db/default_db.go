package db

import (
	"CRUD-service/pkg/entities"
	"CRUD-service/utils"
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type DefaultDB struct {
	data []entities.DefaultEntity
}

func (db *DefaultDB) Load() error {
	projectDir, err := utils.GetProjectRoot()
	fileName := "db/default.json"
	data, err := os.ReadFile(path.Join(projectDir, fileName))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &db.data)
	if err != nil {
		return err
	}

	return nil
}

func (db *DefaultDB) DataToByteSlice() ([]byte, error) {
	fmt.Println(db.data)
	data, err := json.Marshal(db.data)
	if err != nil {
		return data, err
	}

	return data, nil
}
