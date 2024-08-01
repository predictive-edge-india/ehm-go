package models

import (
	"context"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GeoJson struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func (point GeoJson) GormDataType() string {
	return "geography(Point,4326)"
}

func (point GeoJson) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", point.Coordinates[0], point.Coordinates[1])},
	}
}

func (point *GeoJson) Scan(v interface{}) error {
	data := GeoJson{}
	err := json.Unmarshal([]byte(v.(string)), &data)
	*point = data
	return err
}
