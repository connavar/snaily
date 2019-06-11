package controller

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/model"
)

var NotifyChangeHome func(userID string)
var NotifyChangeItems func(ID string)

func InsertItem(conn *pgx.Conn, item *model.Item) (err error) {
	_, err = conn.Exec("insert into madast.items values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		item.ID,
		item.Parent,
		item.Title,
		item.Body,
		item.CreateDate,
		item.CreatedBy,
		item.WaitingFor,
		item.OrgID,
		item.WaitingForDone,
		item.WaitingForDoneDate,
		item.CreatedByDone,
		item.CreatedByDoneDate)
	if err != nil {
		logrus.Errorln(err)
	}
	NotifyChangeHome(item.CreatedBy)
	NotifyChangeHome(item.WaitingFor)
	return
}

func GetItem(conn *pgx.Conn, ID string) (item *model.Item, err error) {
	rows, err := conn.Query("select * from madast.items where id=$1", ID)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	out := make(chan *model.Item)
	go rowsToItemChan(rows, out)
	for r := range out {
		item = r
		return
	}
	return

}

func GetItemsByParent(conn *pgx.Conn, ID string) (out chan *model.Item, err error) {
	rows, err := conn.Query("select * from madast.items where parent_id=$1", ID)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	out = make(chan *model.Item)
	go rowsToItemChan(rows, out)
	return
}

func GetItemsByCreatedBy(conn *pgx.Conn, ID string) (out chan *model.Item, err error) {
	rows, err := conn.Query("select * from madast.items where created_by=$1", ID)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	out = make(chan *model.Item)
	go rowsToItemChan(rows, out)
	return
}

func GetItemsByWaitingFor(conn *pgx.Conn, ID string) (out chan *model.Item, err error) {
	rows, err := conn.Query("select * from madast.items where waiting_for=$1", ID)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	out = make(chan *model.Item)
	go rowsToItemChan(rows, out)
	return
}

func rowsToItemChan(rows *pgx.Rows, out chan *model.Item) {
	var err error
	for rows.Next() {
		item := model.Item{}
		err = rows.Scan(
			&item.ID,
			&item.Parent,
			&item.Title,
			&item.Body,
			&item.CreateDate,
			&item.CreatedBy,
			&item.WaitingFor,
			&item.OrgID,
			&item.WaitingForDone,
			&item.WaitingForDoneDate,
			&item.CreatedByDone,
			&item.CreatedByDoneDate)
		if err != nil {
			logrus.Errorln(err)
		}
		out <- &item
	}
	close(out)
	return
}
