package myutil

import (
	"log"
)

func failOnError(err error, msg string) error {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		return err
	}
	return nil
}
