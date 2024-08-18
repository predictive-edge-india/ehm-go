package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FindDeviceById(id uuid.UUID) models.Device {
	var device models.Device
	Database.Where("id = ?", id).Find(&device)
	return device
}

func FindUnassignedDeviceById(id uuid.UUID) models.Device {
	var device models.Device
	Database.
		Joins("LEFT JOIN asset_devices ON devices.id = asset_devices.device_id").
		Where("asset_devices.device_id IS NULL").
		Where("devices.id = ?", id).
		Find(&device)
	return device
}

func FindDeviceBySerialNo(serialNo string) models.Device {
	var device models.Device
	Database.Where("serial_no = ?", serialNo).Find(&device)
	return device
}
