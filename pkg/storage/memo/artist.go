package memo

import (
	"../../../../lineuplist"
	"fmt"
)

type ArtistStorage struct {
	*storage
}
func NewArtistStorage(s *storage) *ArtistStorage {
	return &ArtistStorage{s}
}

func (as ArtistStorage) Create(a lineuplist.Artist) (lineuplist.Artist, error) {
	as.storage.artists = append(as.storage.artists, a)
	return a, nil
}

func (as ArtistStorage) Read(name string) (lineuplist.Artist, error) {
	for _, a := range as.storage.artists {
		if a.Name == name {
			return a, nil
		}
	}
	return lineuplist.Artist{}, fmt.Errorf("Artist %s not found", name)
}

func (as ArtistStorage) List() ([]lineuplist.Artist, error) {
	return as.storage.artists, nil
}

func (as ArtistStorage) Update(a lineuplist.Artist) (lineuplist.Artist, error) {
	for _, b := range as.storage.artists {
		if b.Name == a.Name {
			b = a
			return b, nil
		}
	}
	return lineuplist.Artist{}, fmt.Errorf("Artist %s not found", a.Name)
}

func (as ArtistStorage) Delete(name string) error {
	for i, a := range as.storage.artists {
		if a.Name == name {
			as.storage.artists = append(as.storage.artists[:i], as.storage.artists[i+1:]...)
			return nil
		}
	}
	return nil
}
