package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FindPowerLimitByAssetIdAndLimitType(assetId uuid.UUID, limitType string) models.PowerLimit {
	var powerLimit models.PowerLimit
	Database.Where("asset_id = ? AND type = ?", assetId, limitType).Find(&powerLimit)
	return powerLimit
}
