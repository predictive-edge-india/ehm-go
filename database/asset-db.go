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

func FindAssetByIdWithAssetClass(id uuid.UUID) models.Asset {
	var asset models.Asset
	Database.Preload("AssetClass").Where("id = ?", id).Find(&asset)
	return asset
}
