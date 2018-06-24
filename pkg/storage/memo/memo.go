package memo

import "../../../../lineuplist"

//
type storage struct {
	artists []lineuplist.Artist
	festivals []lineuplist.Festival
}

func New() *lineuplist.Storage {
	var store storage
	return &lineuplist.Storage{
		Artist: NewArtistStorage(&store),
		Festival: NewFestivalStorage(&store),
	}
}

