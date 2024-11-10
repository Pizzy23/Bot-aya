package db

import (
	"gorm.io/gorm"
)

func Create(db *gorm.DB, table interface{}) error {
	return db.Create(table).Error
}

func Read(db *gorm.DB, table interface{}, condition interface{}) error {
	return db.Where(condition).First(table).Error
}

func GetByCell(db *gorm.DB, table interface{}, condition map[string]interface{}) (bool, error) {
	err := db.Where(condition).First(table).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func GetByID(db *gorm.DB, table interface{}, id int64) (bool, error) {
	err := db.First(table, id).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func Delete(db *gorm.DB, table interface{}, condition interface{}) error {
	return db.Where(condition).Delete(table).Error
}

func GetAll(db *gorm.DB, tableSlice interface{}) error {
	return db.Find(tableSlice).Error
}

func GetAllWithCondition(db *gorm.DB, tableSlice interface{}, condition string, args ...interface{}) error {
	return db.Where(condition, args...).Find(tableSlice).Error
}

func GetPendingSlips(db *gorm.DB, clientID uint) ([]Slip, error) {
	var slips []Slip
	err := db.Where("client_id = ? AND value > 0", clientID).Find(&slips).Error
	if err != nil {
		return nil, err
	}
	return slips, nil
}

func UpdateDebt(db *gorm.DB, clientID uint) error {
	var dataClient DataClient
	err := db.Where("client_id = ?", clientID).First(&dataClient).Error
	if err != nil {
		return err
	}

	var totalDebt float64
	slips, err := GetPendingSlips(db, clientID)
	if err != nil {
		return err
	}

	for _, slip := range slips {
		totalDebt += slip.Value
	}

	// Atualiza a dívida no registro DataClient
	dataClient.Debt = totalDebt
	return db.Save(&dataClient).Error
}
