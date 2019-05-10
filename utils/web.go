package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func WriteXToWriter(w http.ResponseWriter, o interface{}) (err error) {
	b, err := json.Marshal(o)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	return
}
