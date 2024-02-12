package database

import (
	"github.com/google/uuid"
	"github.com/iisc/demo-go/models"
)

func FindEhmDevice(serialNo string) models.EhmDevice {
	var ehmDevice models.EhmDevice
	Database.First(&ehmDevice, "serial_no = ?", serialNo)
	return ehmDevice
}

func FindOrCreateEhmDevice(serialNo string) (models.EhmDevice, error) {
	ehmDevice := FindEhmDevice(serialNo)
	if ehmDevice.Id == uuid.Nil {
		return CreateEhmDevice(serialNo)
	}
	return ehmDevice, nil
}

func CreateEhmDevice(serialNo string) (models.EhmDevice, error) {
	ehmDevice := new(models.EhmDevice)
	ehmDevice.SerialNo = serialNo

	err := Database.Save(&ehmDevice).Error
	if err != nil {
		return *ehmDevice, err
	}
	return *ehmDevice, nil
}
