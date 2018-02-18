package store

import (
	"golang.org/x/crypto/sha3"
	"errors"
)

type Storage interface {
	Store ([]byte) (string, error)
	Get (string) ([]byte, error)
}

type MemoryStorage struct {
	values map[string][]byte
}

func (ms MemoryStorage) Store (value []byte) (string, error) {
	hash := sha3.New256()
	name := string(hash.Sum(value))
	ms.values[name] = value
	return name, nil
}

func (ms MemoryStorage) Get (name string) ([]byte, error) {
	value, present := ms.values[name]
	if !present {
		return value, errors.New("%s not found")
	}
	return value, nil
}

func NewMemoryStore () MemoryStorage {
	return MemoryStorage{make(map[string][]byte)}
}

type DirStorage struct {
	path string
}

func (ds DirStorage) Store (value []byte) (string, error) {
	return "", nil
}

func (ds DirStorage) Get (name string) ([]byte, error) {
	return []byte{}, nil
}

func NewDirStore(path string) DirStorage {
	return DirStorage{path}
}
