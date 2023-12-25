package files

import (
	"encoding/gob"
	"errors"
	"os"
	"path/filepath"

	"github.com/DEHbNO4b/tgbot-advertisement/internal/lib/e"
	"github.com/DEHbNO4b/tgbot-advertisement/internal/storage"
)

const defaultPerm = 0774

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(p *storage.Page) error {

	if err := os.MkdirAll(s.basePath, defaultPerm); err != nil {
		return e.Wrap("unable to save page", err)
	}

	fPath := filepath.Join(s.basePath, "store")
	file, err := os.Create(fPath)
	if err != nil {
		return e.Wrap("unable to create file in save method", err)
	}
	defer file.Close()

	if err := gob.NewEncoder(file).Encode(p); err != nil {
		return e.Wrap("unable to encode in gob", err)
	}
	return nil
}

func (s Storage) Get() (storage.Page, error) {
	files, err := os.ReadDir(s.basePath)
	if err != nil {
		return storage.Page{}, e.Wrap("can't get page", err)
	}
	if len(files) == 0 {
		return storage.Page{}, errors.New("there are no storage files")
	}
}
func (s Storage) decodePage() (storage.Page, error) {
	f, err := os.Open(filepath.Join(s.basePath, "store"))
	if err != nil {
		return storage.Page{}, e.Wrap("unable to open storage file", err)
	}
	defer f.Close()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return storage.Page{}, e.Wrap("unable to decode gob", err)
	}

	return p, nil
}
