package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FindAssetById(id uuid.UUID) models.Asset {
	var asset models.Asset
	Database.Where("id = ?", id).Find(&asset)
	return asset
}
