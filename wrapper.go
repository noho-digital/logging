package logging

import (
	"fmt"
	"github.com/noho-digital/insurews/pkg/errors"
)

func ErrorOnlyWrapper(logger Logger, while string, args ...interface{}) func(error, string, ...interface{}) error {
	return errorWrapper(false, logger, while, args...)
}

func ErrorWrapper(logger Logger, while string, args ...interface{}) func(error, string, ...interface{}) error {
	return errorWrapper(true, logger, while, args...)
}

func errorWrapper(infoAnnounce bool, logger Logger, while string, args ...interface{}) func(error, string, ...interface{}) error {
	if len(args) > 0 {
		while = fmt.Sprintf(while, args...)
	}
	if logger != nil && infoAnnounce {
		logger.Info(while)
	}
	errorWhile := "error " + while
	return func(err error, msg string, args ...interface{}) error {
		if while != "" {
			if msg != "" || len(args) > 0 {
				if len(args) > 0 {
					err = errors.Wrapf(err, msg, args...)
				} else {
					err = errors.Wrap(err, msg)
				}
			}
			err = errors.Wrap(err, errorWhile)
		}
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
}
