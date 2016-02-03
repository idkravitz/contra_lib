package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	TASK_STATUS_PENDING = "pending"
	TASK_STATUS_DONE    = "done"
)

type User struct {
	Username string
	Email    string
	Password []byte
	Dor      time.Time
	Banned   bool
}

type Session struct {
	Sid       string
	Username  string
	CreatedAt time.Time
}

type FileDescription struct {
	Filename       string
	Owner_Username string
}

type FileShortMeta struct {
	Md5        string
	Id         string
	Filename   string
	Size       int
	UploadDate time.Time
}

type TaskMsg struct {
	TaskId     bson.ObjectId
	DataFid    string
	ControlFid string
}

type Task struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Output     string
	Status     string
	Owner      string
	DataFid    string
	OutputFid  string `json:"output_fid" bson:"output_fid,omitempty"`
	ControlFid string
}
