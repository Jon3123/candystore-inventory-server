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

//GetKeys - get keys
func (s *Storage) GetKeys() (keys []string) {
	c := s.DB.Keys()
	var val []byte
	var ok = true
	for ok == true { //wait for channel to be fill
		select {
		case val, ok = <-c:
			keys = append(keys, string(val))
		}
	}
	return
}

//GetAllMappings get all mappings
func (s *Storage) GetAllMappings() (values []string) {
	keys := s.GetKeys()
	for _, key := range keys {
		values = append(values, s.GetValue(key))
	}
	return
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
