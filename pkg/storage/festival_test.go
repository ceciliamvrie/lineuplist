package storage_test

import (
	"testing"

	"../../../lineuplist"
	"./postgres"
)

var f lineuplist.Festival

func TestFestival(t *testing.T) {
	postgres.MigrateDown("file://postgres/migrations/", dsn)
	postgres.MigrateUp("file://postgres/migrations/", dsn)
	f = lineuplist.Festival{
		Name: "Austin City Limits",
		Lineup: []lineuplist.ArtistPreview{
			{Name: "Paul McCartney"},
		},
	}
}

func TestFestivalCreate(t *testing.T) {
	storedF, err := storage.Festival.Create(f)
	if err != nil {
		t.Errorf("Failed creating festival: %s", err)
	}
	if storedF.Name != f.Name {
		t.Errorf("Failed creating festival: Festival name %s does not match %s", storedF.Name, f.Name)
	}
	if storedF.Lineup[0].Name != f.Lineup[0].Name {
		t.Errorf("Failed creating festival: Festival lineup %s does not match %s", storedF.Lineup[0].Name, f.Lineup[0].Name)
	}
}

func TestFestivalRead(t *testing.T) {
	storedF, err := storage.Festival.Read(f.Name)
	if err != nil {
		t.Errorf("Failed reading festival: %s", err)
	}
	if storedF.Name != f.Name {
		t.Errorf("Failed reading festival: Expected %s but got %s", f.Name, storedF.Name)
	}
	if storedF.Lineup[0].Name != f.Lineup[0].Name {
		t.Errorf("Failed reading festival: Festival lineup %s does not match %s", storedF.Lineup[0].Name, f.Lineup[0].Name)
	}
}

func TestFestivalList(t *testing.T) {
	storedFF, err := storage.Festival.List()
	if err != nil {
		t.Errorf("Failed listing festivals: %s", err)
	}
	if len(storedFF) == 0 {
		t.Errorf("Failed listing festivals:/n There should be at least one festival in stored Festivals")
	}
}

func TestFestivalUpdate(t *testing.T) {
	nf := lineuplist.Festival{
		Name: "Austin City Limits",
		Lineup: []lineuplist.ArtistPreview{
			{Name: "Paul McCartney"},
			{Name: "John Lennon"},
		},
	}
	updatedF, err := storage.Festival.Update(nf)
	if err != nil {
		t.Errorf("Failed updating festivals: %s", err)
	}
	if len(updatedF.Lineup) != 2 {
		t.Errorf("Failed updating festivals:/n Expected updated to be %#v, bit got %#v", nf, updatedF)
	}
}

func TestFestivalDelete(t *testing.T) {
	err := storage.Festival.Delete(f.Name)
	if err != nil {
		t.Errorf("Failed deleating festival: %s", err)
	}
}
