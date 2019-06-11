package model

type MessageNewItemResponseV1 struct {
	OK bool   `json:"ok"`
	ID string `json:"id"`
}

type MessageRegisterResponseV1 struct {
	OK bool `json:"ok"`
}

type MessageLoginResponseV1 struct {
	OK  bool   `json:"ok"`
	Key string `json:"key"`
	ID  string `json:"id"`
}

type MessageMyItemsResponseV1 struct {
	CreatedByMe  []*Item        `json:"createdByMe"`
	WaitingForMe []*Item        `json:"waitingForMe"`
	Users        MessageUsersV1 `json:"users"`
}

type MessageGetItemResponseV1 struct {
	Item  *Item          `json:"item"`
	Users MessageUsersV1 `json:"users"`
}

type MessageUsersV1 []*MessageUserV1

func (m *MessageUsersV1) Contains(ID string) bool {
	for _, r := range *m {
		if r.ID == ID {
			return true
		}
	}
	return false
}

type MessageUserV1 struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
