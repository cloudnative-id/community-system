package storage

type Storage interface {
	WriteMeetup() error
	GetMeetup() 
	DeleteMeetup()
}
