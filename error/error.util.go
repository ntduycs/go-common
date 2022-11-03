package error

import (
	"github.com/go-errors/errors"
	"log"
)

func PrintStacktrace(err error) {
	var wrappedErr *errors.Error

	if errors.As(err, &wrappedErr) {
		log.Println(err.(*errors.Error).ErrorStack())
	} else {
		log.Println(wrap(err).ErrorStack())
	}
}

func wrap(err error) *errors.Error {
	if err != nil {
		return errors.Wrap(err, 2)
	} else {
		return errors.Errorf("unexpected error")
	}
}

func RecoverOnPanic() {
	if err := recover(); err != nil {
		log.Printf("recovered from panic. %v", err)
		log.Println(errors.Wrap(err, 2).ErrorStack())
	}
}
