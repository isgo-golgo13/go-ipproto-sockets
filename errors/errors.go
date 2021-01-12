package errors

import (
	log "github.com/Sirupsen/logrus"
)

func CheckError(err error) {
	if err != nil {
		log.Errorf("%v", err)
	}
}
