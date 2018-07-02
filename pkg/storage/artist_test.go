package storage_test

import (
	"testing"

	"../../../lineuplist"
	"./postgres"
)

var a lineuplist.Artist

func TestArtist(t *testing.T) {
	postgres.MigrateDown("file://postgres/migrations/", dsn)
	postgres.MigrateUp("file://postgres/migrations/", dsn)
	a = lineuplist.Artist{
		Name: "John Lemon",
	}
}

func TestArtistCreate(t *testing.T) {
	storedA, err := storage.Artist.Create(a)
	if err != nil {
		t.Errorf("Failed creating artist: %s", err)
	}
	if storedA.Name != a.Name {
		t.Errorf("Failed creating artist: Artist name %s does not match %s", storedA.Name, a.Name)
	}
}

func TestArtistRead(t *testing.T) {
	storedA, err := storage.Artist.Read(a.Name)
	if err != nil {
		t.Errorf("Failed reading artist: %s", err)
	}
	if storedA.Name != a.Name {
		t.Errorf("Failed reading artist: Expected %s but got %s", a.Name, storedA.Name)
	}
}

func TestArtistList(t *testing.T) {
	storedAF, err := storage.Artist.List()
	if err != nil {
		t.Errorf("Failed listing artists: %s", err)
	}
	if len(storedAF) == 0 {
		t.Errorf("Failed listing artists:/n There should be at least one artist in stored artists")
	}
}

func TestArtistUpdate(t *testing.T) {
	na := lineuplist.Artist{
		Name: "John Lemon",
		Festivals: []lineuplist.FestivalPreview{
			{Name: "Austin City Limits"},
			{Name: "Coachella"},
		},
	}
	updatedA, err := storage.Artist.Update(na)
	if err != nil {
		t.Errorf("Failed updating artists: %s", err)
	}
	if len(updatedA.Festivals) != 2 {
		t.Errorf("Failed updating artists:/n Expected updated to be %#v, bit got %#v", na, updatedA)
	}
}

func TestArtistDelete(t *testing.T) {
	err := storage.Artist.Delete(a.Name)
	if err != nil {
		t.Errorf("Failed deleating Artist: %s", err)
	}
}
