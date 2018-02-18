package store

import (
	"encoding/hex"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Mkdir("./tmp", os.ModeDir)

	result := m.Run()

	os.Exit(result)
}

// TestStorages ensures that values are stored
func TestStoreGet(t *testing.T) {
	value := make([]byte, 4)
	r := rand.New(rand.NewSource(7))
	r.Read(value)
	for _, s := range []Storage{
		NewMemoryStore(),
		NewDirStore("./tmp"),
	} {
		name, err := s.Store(value)
		if err != nil {
			t.Errorf("ERR: %T.Store(...) => _, %s", s, err)
		}
		got, err := s.Get(name)
		if err != nil {
			t.Errorf("ERR: %T.Get(...) => _, %s", s, err)
		}
		if string(got) != string(value) {
			t.Errorf(
				"%T.Get(%s) => %s, _ want %s",
				s,
				name,
				hex.EncodeToString(got),
				hex.EncodeToString(value),
			)
		}
	}
}

func TestCanPass(t *testing.T) {
	return
	t.Errorf("can fail")
}


