package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FindAssetClassById(id uuid.UUID) models.AssetClass {
	var assetClass models.AssetClass
	Database.Where("id = ?", id).Find(&assetClass)
	return assetClass
}
