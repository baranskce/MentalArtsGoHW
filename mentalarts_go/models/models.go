package models

// Musician represents an individual artist or band member.
type Musician struct {
	ID        string "json:\"id\""
	FirstName string "json:\"first_name\""
	LastName  string "json:\"last_name\""
	Age       uint   "json:\"age\""
	Genre     string "json:\"genre\""
}

// Album represents a music album.
type Album struct {
	ID          uint     "json:\"id\" gorm:\"primaryKey\""
	Title       string   "json:\"title\""
	MusicianID  string   "json:\"musician_id\""
	ReleaseYear int      "json:\"release_year\""
	Musician    Musician "json:\"musician\" gorm:\"foreignKey:MusicianID\""
}
