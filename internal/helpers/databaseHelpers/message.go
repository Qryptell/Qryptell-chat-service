package databasehelpers

import (
	"fmt"

	"github.com/Qryptell/Qryptell-chat-service/pkg/message"
	"github.com/Qryptell/Qryptell-chat-service/pkg/scylla"
)

// Inserting text message to database
func InsertTextMessage(msg message.Msg) {
	var err = scylla.Session.Query("INSERT INTO textmsg(id,msg_from,msg_to,msg_time,type,content,message) VALUES (?,?,?,?,?,?,?);", msg.Id, msg.From, msg.To, msg.Time, msg.Type, msg.Content, msg.Message).Exec()
	if err != nil {
		fmt.Println(err)
	}
}
