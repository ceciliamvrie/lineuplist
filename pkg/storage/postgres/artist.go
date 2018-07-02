package postgres

import (
	"fmt"
	"log"

	"../../../../lineuplist"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

/* NewArtistStorage returns ArtistStorage postgres implementation */
func NewArtistStorage(dsn string) *ArtistStorage {
	log.Println("dsn: ", dsn)

	return &ArtistStorage{sqlx.MustConnect("postgres", dsn)}
}

/* ArtistStorage implements lineuplist.ArtistStorage */
type ArtistStorage struct {
	*sqlx.DB
}

/* Create creates and writes a new artist to db */
func (db *ArtistStorage) Create(a lineuplist.Artist) (lineuplist.Artist, error) {
	id, err := uuid.NewV4()
	a.ID = id.String()
	if err != nil {
		return lineuplist.Artist{}, fmt.Errorf("failed creating artist id: %s", err)
	}
	_, err = db.Exec("INSERT INTO artist(id, name) VALUES($1, $2)", a.ID, a.Name)
	if err != nil {
		return lineuplist.Artist{}, fmt.Errorf("failed inserting artist %#v: %s", a, err)
	}
	return a, nil
}

/* Read returns the first stored artist in db given provided name */
func (db *ArtistStorage) Read(name string) (lineuplist.Artist, error) {
	a := lineuplist.Artist{}
	err := db.Get(&a, "SELECT * FROM artist WHERE name = $1", name)
	if err != nil {
		return lineuplist.Artist{}, fmt.Errorf("artist with name: %s not found: %s", name, err)
	}
	f, err := db.readFestivals(a.Name)
	if err != nil {
		fmt.Errorf("failed reading artist festivals from db: %s", err)
	}
	a.Festivals = f
	return a, nil
}

/* Update updates an artist and returns it if exists in db */
func (db *ArtistStorage) Update(a lineuplist.Artist) (lineuplist.Artist, error) {
	q := "INSERT INTO artist(id, name) VALUES($1, $2) ON CONFLICT DO NOTHING"
	_, err := db.Exec(q, a.ID, a.Name)
	if err != nil {
		return lineuplist.Artist{}, fmt.Errorf("failed updating artist/n %s", err)
	}
	db.Get(&a, "SELECT * FROM artist WHERE name = $1", a.Name)
	return a, nil
}

/* Delete removes an artist from db if exists */
func (db *ArtistStorage) Delete(name string) error {
	_, err := db.Exec("DELETE FROM artists WHERE name = $1", name)
	if err != nil {
		return fmt.Errorf("failed deleating artist from db: %s", err)
	}
	return nil
}

/* List returns a list of all artists in db */
func (db *ArtistStorage) List() ([]lineuplist.Artist, error) {
	aa := []lineuplist.Artist{}
	err := db.Select(&aa, "SELECT * FROM artist")
	if err != nil {
		return []lineuplist.Artist{}, fmt.Errorf("failed reading artists from db: %s", err)
	}
	for i, a := range aa {
		f, err := db.readFestivals(a.Name)
		if err != nil {
			fmt.Errorf("failed reading artist festivals from db:%s", err)
		}
		aa[i].Festivals = f
	}
	return aa, nil
}

/* ReadFestivals returns list of festivals from associated artist */
func (db *ArtistStorage) readFestivals(artistName string) ([]lineuplist.FestivalPreview, error) {
	ff := []lineuplist.FestivalPreview{}
	q := `SELECT * FROM festival
		WHERE id = (
			SELECT id FROM festival_artist	
			WHERE artist_id = (
				SELECT id FROM artist
				WHERE name = $1
			)
		)`

	err := db.Get(&ff, q, artistName)
	if err != nil {
		return []lineuplist.FestivalPreview{}, err
	}
	return ff, nil
}
