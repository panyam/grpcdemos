package main

import (
	"time"

	"github.com/lib/pq"
)

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Id        string `gorm:"primaryKey"`
	Version   int    // used for optimistic locking
}

type Artist struct {
	BaseModel

	// Full name of the artist
	Name string

	// Artist's date of
	DateOfBirth time.Time

	Country string
}

// Album showcasing a bunch of artists performing a bunch of songs.
type Album struct {
	BaseModel

	// Album name
	Name string

	Venue string

	ReleaseDate time.Time

	// Songs performed in this album
	SongIds pq.StringArray `gorm:"type:text[]" json:"tokens"`
}

type Label struct {
	BaseModel

	// Name of this record label company (eg, "Virgin Records")
	Name string

	// When this company was established
	EstablishedDate time.Time
}

// A song - eg Jingle Bells
type Song struct {
	BaseModel
	// Name for this song
	Name     string
	Composer string
	Homepage string

	// All performing artists for this song
	Performances []Performance
}

type Performance struct {
	// ID of the song being performed
	SongID string

	// Venue where the song was performed
	Venue string

	// Lead artist performing the song
	Performer Artist

	// When did this happen?
	PerformedOn time.Time
}
