package lineuplist

import "time"

/* Storage holds both FestivalStorage and ArtistStorage */
type Storage struct {
	Festival FestivalStorage
	Artist   ArtistStorage
}

/* Festival is a representation of a music festival */
type Festival struct {
	ID        string          `json:"-"`
	Name      string          `json:"name"`
	Lineup    []ArtistPreview `json:"lineup"`
	ImgSrc    string          `json:"imgSrc" db:"img_src"`
	StartDate time.Time       `json:"startDate" db:"start_date"`
	EndDate   time.Time       `json:"endDate" db:"end_date"`
	Country   string          `json:"country"`
	State     string          `json:"state"`
	City      string          `json:"city"`
}

type FestivalPreview struct {
	ID        string    `json:"-"`
	Name      string    `json:"name"`
	ImgSrc    string    `json:"imgSrc" db:"img_src"`
	StartDate time.Time `json:"startDate" db:"start_date"`
	EndDate   time.Time `json:"endDate" db:"end_date"`
	Country   string    `json:"country"`
	State     string    `json:"state"`
	City      string    `json:"city"`
}

/* FestivalStorage is an interface for creating, updating, reading, and deleating festivals */
type FestivalStorage interface {
	Create(Festival) (Festival, error)
	Read(name string) (Festival, error)
	Update(f Festival) (Festival, error)
	Delete(name string) error
	List() ([]Festival, error)
}

/* Artist is a representation of a musician or band */
type Artist struct {
	ID          string            `json:"-"`
	Name        string            `json:"name"`
	ImgSrc      string            `json:"imgSrc" db:"img_src"`
	ExternalURL string            `json:"externalURL" db:"external_url"`
	Popularity  int               `json:"popularity"`
	Followers   int               `json:"followers"`
	Genres      []string          `json:"genres"`
	TopTracks   []Track           `json:"topTracks" db:"top_tracks"`
	Albums      []Album           `json:"albums"`
	Related     []Artist          `json:"relatedArtists" db:"related_artist"`
	Festivals   []FestivalPreview `json:"festivals"`
}

type ArtistPreview struct {
	ID          string   `json:"-"`
	Name        string   `json:"name"`
	ImgSrc      string   `json:"imgSrc" db:"img_src"`
	ExternalURL string   `json:"externalURL" db:"external_url"`
	Popularity  int      `json:"popularity"`
	Followers   int      `json:"followers"`
	Genres      []string `json:"genres"`
}

/* ArtistStorage is an interface for creating, updating, reading, and deleating artists */
type ArtistStorage interface {
	Create(Artist) (Artist, error)
	Read(name string) (Artist, error)
	Update(a Artist) (Artist, error)
	Delete(name string) error
	List() ([]Artist, error)
}

/* Track is a representation of an artist's song */
type Track struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	ExternalURL string `json:"externalUrl"`
	Album       `json:"album"`
}

/* Album is a representation of an artist's music album */
type Album struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	ImgSrc      string `json:"imgSrc"`
	ExternalURL string `json:"externalUrl"`
}
