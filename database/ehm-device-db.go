package database

import (
	"github.com/google/uuid"
	"github.com/iisc/demo-go/models"
)

func FindAllEhmDevices() []models.EhmDevice {
	var devices []models.EhmDevice
	Database.Find(&devices, "deleted_at IS NULL")
	return devices
}

func FindEhmDeviceById(ehmDeviceId uuid.UUID) models.EhmDevice {
	var ehmDevice models.EhmDevice
	Database.First(&ehmDevice, "id = ?", ehmDeviceId)
	return ehmDevice
}

func FindEhmDeviceBySerialNo(serialNo string) models.EhmDevice {
	var ehmDevice models.EhmDevice
	Database.First(&ehmDevice, "serial_no = ?", serialNo)
	return ehmDevice
}

func FindOrCreateEhmDevice(serialNo string) (models.EhmDevice, error) {
	ehmDevice := FindEhmDeviceBySerialNo(serialNo)
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
