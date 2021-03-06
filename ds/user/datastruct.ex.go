package user

import (
	"sync"
)

const (
	MAX_EMAIL       = 50
	MAX_PHONE_NUM   = 15
	MIN_GENDER      = 1
	MAX_GENDER      = 3
	MAX_NICKNAME    = 8
	MAX_AVATOR      = 120
	MAX_NAME        = 20
	MAX_ID          = 18
	MAX_NATIONALITY = 5
)

var poolEntry = sync.Pool{
	New: func() interface{} {
		return &Entry{}
	},
}

func NewEntry() *Entry {
	return poolEntry.Get().(*Entry)
}

func DeleteEntry(obj *Entry) {
	poolEntry.Put(obj)
}

var poolBasic = sync.Pool{
	New: func() interface{} {
		return &Basic{}
	},
}

func NewBasic() *Basic {
	return poolBasic.Get().(*Basic)
}

func DeleteBasic(obj *Basic) {
	poolBasic.Put(obj)
}

var poolPersonal = sync.Pool{
	New: func() interface{} {
		return &Personal{}
	},
}

func NewPersonal() *Personal {
	return poolPersonal.Get().(*Personal)
}

func DeletePersonal(obj *Personal) {
	poolPersonal.Put(obj)
}

var poolRelation = sync.Pool{
	New: func() interface{} {
		return &Relation{}
	},
}

func NewRelation() *Relation {
	return poolRelation.Get().(*Relation)
}

func DeleteRelation(obj *Relation) {
	poolRelation.Put(obj)
}

var poolContact = sync.Pool{
	New: func() interface{} {
		return &Contact{}
	},
}

func NewContact() *Contact {
	return poolContact.Get().(*Contact)
}

func DeleteContact(obj *Contact) {
	poolContact.Put(obj)
}

var poolLoginLog = sync.Pool{
	New: func() interface{} {
		return LoginLog{}
	},
}

func NewLoginLog() *LoginLog {
	return poolLoginLog.Get().(*LoginLog)
}

func DeleteLoginLog(obj *LoginLog) {
	poolLoginLog.Put(obj)
}

var poolArchiveLog = sync.Pool{
	New: func() interface{} {
		return &ArchiveLog{}
	},
}

func NewArchiveLog() *ArchiveLog {
	return poolArchiveLog.Get().(*ArchiveLog)
}

func DeleteArchiveLog(obj *ArchiveLog) {
	poolArchiveLog.Put(obj)
}
