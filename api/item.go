package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/bus"
	"github.com/team142/snaily/model"
	"github.com/team142/snaily/utils"
	"io/ioutil"
	"net/http"
)

func handleCreateItem(w http.ResponseWriter, r *http.Request, ID string) {
	item, err := model.ReadCloserToItem(r.Body)
	if err != nil {
		http.Error(w, "Invalid request or body", http.StatusBadRequest)
		logrus.Errorln(err)
		return
	}

	var created bool

	if created, err = bus.CreateItem(item, ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Errorln(err)
		return
	}

	if err := utils.WriteXToWriter(w, model.MessageNewItemResponseV1{OK: created, ID: item.ID}); err != nil {
		logrus.Errorln(err)
	}

}

func handleGetItem(w http.ResponseWriter, r *http.Request, ID string) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	itemReq := model.Item{}
	if err = json.Unmarshal(b, &itemReq); err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bus.GetItem(itemReq.ID)

	if err = utils.WriteXToWriter(w, result); err != nil {
		logrus.Errorln(err)
	}

}

func handleGetMyItems(w http.ResponseWriter, r *http.Request, ID string) {
	result, err := bus.GetMyItems(ID)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	if err = utils.WriteXToWriter(w, result); err != nil {
		logrus.Errorln(err)
	}

}
