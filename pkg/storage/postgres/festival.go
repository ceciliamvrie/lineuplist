package postgres

import (
	"fmt"

	"../../../../lineuplist"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

/* NewFestivalStorage returns a FestivalStorage postgres implementation */
func NewFestivalStorage(dsn string) *FestivalStorage {
	return &FestivalStorage{sqlx.MustConnect("postgres", dsn)}
}

/* FestivalStorage implements lineuplist.FestivalStorage */
type FestivalStorage struct {
	*sqlx.DB
}

/* Create creates and writes a new festival to db */
func (db *FestivalStorage) Create(f lineuplist.Festival) (lineuplist.Festival, error) {
	id, err := uuid.NewV4()
	f.ID = id.String()
	if err != nil {
		return lineuplist.Festival{}, fmt.Errorf("failed creating festival id: %s", err)
	}
	_, err = db.Exec("INSERT INTO festival(id, name) VALUES($1, $2)", f.ID, f.Name)
	if err != nil {
		return lineuplist.Festival{}, fmt.Errorf("failed inserting festival %#v", f.ID, f.Name)
	}
	return f, nil
}

/* Read returns the first stored festival in db given provided name */
func (db *FestivalStorage) Read(name string) (lineuplist.Festival, error) {
	f := lineuplist.Festival{}
	err := db.Get(&f, "SELECT * FROM festival WHERE name = $1", name)
	if err != nil {
		return lineuplist.Festival{}, fmt.Errorf("festival with name: %s not found: %s", name, err)
	}
	as, err := db.readArtists(f.Name)
	if err != nil {
		fmt.Errorf("failed reading festival artists from db: %s", err)
	}
	f.Lineup = as
	return f, nil
}

/* Update updates a festival and returns it if exists in db */
func (db *FestivalStorage) Update(f lineuplist.Festival) (lineuplist.Festival, error) {
	q := "INSERT INTO festival(id, name) VALUES($1, $2) ON CONFLICT DO NOTHING"
	_, err := db.Exec(q, f.ID, f.Name)
	if err != nil {
		return lineuplist.Festival{}, fmt.Errorf("failed updating festival %s", err)
	}
	db.Get(&f, "SELECT * FROM festival WHERE name = $1", f.Name)
	return f, nil
}

/* Delete removes a festival from db if exists */
func (db *FestivalStorage) Delete(name string) error {
	_, err := db.Exec("DELETE FROM festivals WHERE name = $1", name)
	if err != nil {
		return fmt.Errorf("failed deleating festival from db: %s", err)
	}
	return nil
}

/* List returns a list of all festivals in db */
func (db *FestivalStorage) List() ([]lineuplist.Festival, error) {
	fs := []lineuplist.Festival{}
	err := db.Get(&fs, "SELECT * FROM festival")
	if err != nil {
		return []lineuplist.Festival{}, fmt.Errorf("failed reading festivals from db: %s", err)
	}
	for i, f := range fs {
		a, err := db.readArtists(f.Name)
		if err != nil {
			fmt.Errorf("failed reading festival artists from db:%s", err)
		}
		fs[i].Lineup = a
	}
	return fs, nil
}

/* ReadArtists returns list of artists from associated festival */
func (db *FestivalStorage) readArtists(festivalName string) ([]lineuplist.ArtistPreview, error) {
	as := []lineuplist.ArtistPreview{}
	q := `SELECT * FROM festival
		WHERE id = (
			SELECT id FROM festival_artist	
			WHERE artist_id = (
				SELECT id FROM artist
				WHERE name = $1
			)
		)`

	err := db.Get(&as, q, festivalName)
	if err != nil {
		return []lineuplist.ArtistPreview{}, err
	}
	return as, nil
}
