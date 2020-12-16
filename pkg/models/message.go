package models

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	Sender int64 `json:"sender"`
	Recipient int64 `json:"recipient"`
	Contents Content `json:"content"`
}

type ResponseMessage struct {
	ID int64
	Timestamp string
}

//To group different types of content under a single interface
type Content interface {
	getType() string
}

type Text struct {
	Type string `json:"type"`
	TextStr string `json:"text"`
}

type Image struct {
	Type string `json:"type"`
	Url string `json:"url"`
	Height int `json:"height"`
	Width int `json:"width"`
}

type Video struct {
	Type string `json:"type"`
	Url string `json:"url"`
	Source string `json:"source"`
}

//MessengerMain processes dsata and delegates & handle message actions
func MessengerMain(token string, messengerPayload string) (int64, string) {
	var newMessage Message
	var err error
	err = json.Unmarshal([]byte(messengerPayload), &newMessage)
	if err != nil {
		log.Errorf("could not unmarshall into struct: %v\n", err)
		return 0, ""
	}
	log.Debugf("Unmarshaled messege payload to %+v", newMessage)
	id, timeStamp := SendMessage(newMessage)
	return id, timeStamp
}

func SendMessage(message Message) (int64, string) {
	/*
		TODO:
			1. Insert message fields (SenderID, RecipientID, Content) into 'CommunicationTable' using sql INSERT query commands
			2. Add AUTOINCREMENT messagesID, and timestamp columns into table
			3. Generate formatted timestamp
			4. Insert timestamp into table
			5. return messgageID and timestamp
	 */
	return 0, ""
}

func (txt Text) getType() string {
	return txt.Type
}

func (img Image) getType() string {
	return img.Type
}

func (vid Video) getType() string {
	return vid.Type
}