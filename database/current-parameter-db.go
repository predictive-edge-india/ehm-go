package database

import (
	"github.com/iisc/demo-go/models"
	log "github.com/sirupsen/logrus"
)

func CreateCurrentParameter(currentParameter models.CurrentParameter) {
	err := Database.Save(&currentParameter).Error
	if err != nil {
		log.Errorln(err.Error())
	}
}
