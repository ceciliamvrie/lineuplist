package memo

import (
	"../../../../lineuplist"
	"fmt"
)


type FestivalStorage struct {
	*storage
}
func NewFestivalStorage(s *storage) *FestivalStorage {
	return &FestivalStorage{s}
}

func (fs FestivalStorage) Create(f lineuplist.Festival) (lineuplist.Festival, error) {
	if f.Name != "" {
		fs.storage.festivals = append(fs.storage.festivals, f)
		return f, nil
	}
	return lineuplist.Festival{}, fmt.Errorf("Failed creating festival: %#v", f)
}

func (fs FestivalStorage) Read(name string) (lineuplist.Festival, error) {
	for _, f := range fs.storage.festivals {
		if f.Name == name {
			return f, nil
		}
	}
	return lineuplist.Festival{}, fmt.Errorf("Festival %s not found", name)
}

func (fs FestivalStorage) Update(a lineuplist.Festival) (lineuplist.Festival, error) {
	for _, f := range fs.storage.festivals {
		if f.Name == a.Name {
			f = a
			return f, nil
		}
	}
	return lineuplist.Festival{}, fmt.Errorf("Festival %s not found", a.Name)
}

func (fs FestivalStorage) Delete(name string) error {
	for i, f := range fs.storage.festivals {
		if f.Name == name {
			fs.storage.festivals = append(fs.storage.festivals[:i], fs.storage.festivals[i+1:]...)
			return nil
		}
	}
	return nil
}

func (fs FestivalStorage) List() ([]lineuplist.Festival, error) {
	return fs.storage.festivals, nil
}