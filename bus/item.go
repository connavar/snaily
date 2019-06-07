package bus

import (
	"errors"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/controller"
	"github.com/team142/snaily/db"
	"github.com/team142/snaily/model"
	"sync"
)

func CreateItem(item *model.Item, createdBy string) (created bool, err error) {
	item.GenerateID()
	item.CreatedBy = createdBy

	conn, err := db.Connect(db.DefaultConfig)
	if err != nil {
		logrus.Errorln(err)
		return false, errors.New("Database connection problem")
	}
	defer conn.Close()

	u, err := controller.GetUserByEmail(conn, item.WaitingFor)
	if err != nil {
		logrus.Errorln(err)
		return false, errors.New("Database read problem")
	}
	if u == nil {
		u = model.NewUserFromEmail(item.WaitingFor)
		if err = controller.InsertUser(conn, u); err != nil {
			logrus.Errorln(err)
			return false, errors.New("Database write problem")
		}
	}
	item.WaitingFor = u.ID

	if err := controller.InsertItem(conn, item); err != nil {
		logrus.Errorln(err)
		return false, errors.New("Database write problem")
	}

	created = true
	return

}

func GetMyItems(ID string) (result model.MessageMyItemsResponseV1, err error) {

	var conn *pgx.Conn
	if conn, err = db.Connect(db.DefaultConfig); err != nil {
		logrus.Errorln(err)
		return
	}
	defer conn.Close()

	user, err := controller.GetUser(conn, ID)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	result = model.MessageMyItemsResponseV1{
		CreatedByMe:  make([]*model.Item, 0),
		WaitingForMe: make([]*model.Item, 0),
		Users:        make(model.MessageUsersV1, 0),
	}

	var wgItems sync.WaitGroup

	in := make(chan string, 10)

	go func(in chan string) {
		var conn *pgx.Conn
		if conn, err = db.Connect(db.DefaultConfig); err != nil {
			logrus.Errorln(err)
			return
		}
		defer conn.Close()
		for ID := range in {
			if !result.Users.Contains(ID) {
				u, err := controller.GetUser(conn, ID)
				if err != nil {
					logrus.Errorln(err)
					wgItems.Done()
					continue
				}
				if u == nil {
					logrus.Errorln("Could not find ID: ", ID)
					wgItems.Done()
					continue
				}
				result.Users = append(result.Users, u.GetUserMessage())
			}
			wgItems.Done()
		}

	}(in)

	wgItems.Add(1)
	go func() {
		var conn *pgx.Conn
		if conn, err = db.Connect(db.DefaultConfig); err != nil {
			logrus.Errorln(err)
			return
		}
		defer conn.Close()

		var out chan *model.Item
		if out, err = controller.GetItemsByCreatedBy(conn, user.ID); err != nil {
			logrus.Errorln(err)
			return
		}
		for item := range out {
			wgItems.Add(1)
			in <- item.WaitingFor
			wgItems.Add(1)
			in <- item.CreatedBy
			result.CreatedByMe = append(result.CreatedByMe, item)
		}
		wgItems.Done()
	}()

	wgItems.Add(1)
	go func() {
		var conn *pgx.Conn
		if conn, err = db.Connect(db.DefaultConfig); err != nil {
			logrus.Errorln(err)
			return
		}
		defer conn.Close()

		var out chan *model.Item
		if out, err = controller.GetItemsByWaitingFor(conn, user.ID); err != nil {
			logrus.Errorln(err)
			return
		}
		for item := range out {
			wgItems.Add(1)
			in <- item.WaitingFor
			wgItems.Add(1)
			in <- item.CreatedBy
			result.WaitingForMe = append(result.WaitingForMe, item)
		}
		wgItems.Done()
	}()

	wgItems.Wait()
	close(in)

	return

}
