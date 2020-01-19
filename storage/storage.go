package storage

import (
	"github.com/prologic/bitcask"
)

//Storage storage struct
type Storage struct {
	Directoryname string
	DB            *bitcask.Bitcask
}

//NewStorage create new storage
func NewStorage(directory string) *Storage {
	db, _ := bitcask.Open("data/" + directory)
	s := &Storage{
		Directoryname: directory,
		DB:            db,
	}

	return s
}

//AddValue - add value
func (s *Storage) AddValue(key string, pair string) {
	s.DB.Put([]byte(key), []byte(pair))
}

//GetValue - get value
func (s *Storage) GetValue(key string) (res string) {
	value, _ := s.DB.Get([]byte(key))
	return string(value)
}

//Close - close
func (s *Storage) Close() {
	s.DB.Close()
}
