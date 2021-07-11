package common

import "github.com/sirupsen/logrus"

func CheckErr(err error, trace string) bool {
	if err != nil {
		Log.WithFields(logrus.Fields{
			"trace": trace,
		}).Error(err)
	}

	return err != nil
}
