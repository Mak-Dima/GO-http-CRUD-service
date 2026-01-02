package db

import (
	"CRUD-service/pkg/entities"
	"CRUD-service/utils"
	"encoding/json"
	"errors"
	"os"
	"path"
	"strconv"
)

const filePath = "db/datafiles/default.json"

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

func (db *DefaultDB) WriteNewEntity(entity entities.DefaultEntity) error {
	db.conter++
	entity.ID = strconv.Itoa(db.conter)
	db.data = append(db.data, entity)

	data, err := db.DataToByteSlice()
	if err != nil {
		return err
	}

	projectDir, err := utils.GetProjectRoot()
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(projectDir, filePath), data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (db *DefaultDB) UpdateEntity(entity entities.DefaultEntity) error {

	for i, e := range db.data {
		if e.ID == entity.ID {
			db.data[i] = entity

			data, err := db.DataToByteSlice()
			if err != nil {
				return err
			}

			projectDir, err := utils.GetProjectRoot()
			if err != nil {
				return err
			}

			err = os.WriteFile(path.Join(projectDir, filePath), data, 0644)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("Entity not found.")
}

func (db *DefaultDB) DeleteEntity(id string) error {
	for i, e := range db.data {
		if e.ID == id {
			db.data = append(db.data[:i], db.data[i+1:]...)

			data, err := db.DataToByteSlice()
			if err != nil {
				return err
			}

			projectDir, err := utils.GetProjectRoot()
			if err != nil {
				return err
			}

			err = os.WriteFile(path.Join(projectDir, filePath), data, 0644)
			if err != nil {
				return err
			}

			break
		}
	}

	return nil
}
