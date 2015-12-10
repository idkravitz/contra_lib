package db

import (
	"time"
	"gopkg.in/mgo.v2"
	"github.com/streadway/amqp"
)


const (
	DATABASE = "tram"
	SESSION_TTL = time.Hour * 24 * 2;
)

func GetCol(s *mgo.Session, colName string) *mgo.Collection {
	return s.DB(DATABASE).C(colName)
}

func RabbitInitConnect(socket string) (*amqp.Connection, error) {
	con, err := amqp.Dial(socket)
    if err != nil {
    	return con, err
    }

    ch, err2 := con.Channel()
    if err2 != nil {
    	return con, err
    }
    if _, err := ch.QueueDeclare("execution_queue", true, false, false, false, nil); err != nil {
    	return con, err
    }
    if err := ch.ExchangeDeclare("workers", "direct", true, false, false, false, nil); err != nil {
    	return con, err
    }
    if err := ch.QueueBind("execution_queue", "task", "workers", false, nil); err != nil {
    	return con, err
    }
    return con, nil
}

func MongoInitConnect(socket string) (*mgo.Session, error) {
	session, err := mgo.Dial(socket)
	if err != nil {
        return session, err
    }

	session.SetSafe(&mgo.Safe{WMode: "majority"})
	cUsers := GetCol(session, "users")
	cSessions := GetCol(session, "sessions")
    cUsers.EnsureIndex(mgo.Index{ Key: []string{"username"}, Unique: true})
    cSessions.EnsureIndex(mgo.Index{ Key: []string{"sid"}, Unique: true})
    cSessions.EnsureIndex(mgo.Index{ Key: []string{"username"}, Unique: true})
    cSessions.EnsureIndex(mgo.Index{ Key: []string{"createdAt"}, ExpireAfter: SESSION_TTL})
    dataFiles := GetCol(session, "data.files")
    dataFiles.EnsureIndex(mgo.Index{ Key: []string{"metadata.owner_username"}})
    controlFiles := GetCol(session, "control.files")
    controlFiles.EnsureIndex(mgo.Index{ Key: []string{"metadata.owner_username"}})

    return session, nil
}