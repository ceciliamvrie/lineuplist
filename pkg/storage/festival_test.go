package storage_test

import (
	"testing"
	"../../../lineuplist"
)
var f lineuplist.Festival
func TestFestival(t *testing.T) {
	f = lineuplist.Festival{
		Name: "Austin City Limits",
		Lineup: []lineuplist.ArtistPreview{
			{ Name: "Paul McCartney" },
		},
	}
}

func TestFestivalCreate(t *testing.T) {
	storedF, err := store.Festival.Create(f)
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
	storedF, err := store.Festival.Read(f.Name)
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
    storedFF, err := store.Festival.List()
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
		    { Name: "Paul McCartney" },
		    { Name: "John Lennon" },
	    },
    }
    updatedF, err := store.Festival.Update(nf)
    if err != nil {
	    t.Errorf("Failed updating festivals: %s", err)
    }
    if len(updatedF.Lineup) != 2 {
        t.Errorf("Failed updating festivals:/n Expected updated to be %#v, bit got %#v", nf, updatedF)
    }
}

func TestFestivalDelete(t *testing.T) {
    err := store.Festival.Delete(f.Name)
	if err != nil {
		t.Errorf("Failed deleating festival: %s", err)
	}
}
