package model

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"time"
)

type Item struct {
	ID                 string    `json:"id"`
	Parent             string    `json:"parent"`
	Title              string    `json:"title"`
	Body               string    `json:"body"`
	CreateDate         time.Time `json:"createdDate"`
	CreatedBy          string    `json:"createdBy"`
	WaitingFor         string    `json:"waitingFor"`
	OrgID              string    `json:"orgID"`
	WaitingForDone     bool      `json:"waitingForDone"`
	WaitingForDoneDate time.Time `json:"waitingForDoneDate"`
	CreatedByDone      bool      `json:"createdByDone"`
	CreatedByDoneDate  time.Time `json:"createdByDoneDate"`
}

func ReadCloserToItem(body io.ReadCloser) (item *Item, err error) {
	var b []byte
	b, err = ioutil.ReadAll(body)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	item = &Item{}
	err = json.Unmarshal(b, item)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}

func (i *Item) GenerateID() {
	i.ID = uuid.NewV4().String()
}
