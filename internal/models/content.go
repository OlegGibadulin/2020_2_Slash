package models

type Content struct {
	ContentID        uint64      `json:"content_id"`
	Name             string      `json:"name"`
	OriginalName     string      `json:"original_name"`
	Description      string      `json:"description"`
	ShortDescription string      `json:"short_description"`
	Rating           int         `json:"rating"`
	Year             int         `json:"year"`
	Images           string      `json:"images"`
	Type             string      `json:"type"`
	IsFree           *bool       `json:"is_free"`
	Countries        []*Country  `json:"countries"`
	Genres           []*Genre    `json:"genres"`
	Actors           []*Actor    `json:"actors"`
	Directors        []*Director `json:"directors"`
	IsLiked          *bool       `json:"is_liked,omitempty"`
	IsFavourite      *bool       `json:"is_favourite"`
}

func (c *Content) ReplaceBy(other *Content) {
	if other.Name != "" {
		c.Name = other.Name
	}
	if other.OriginalName != "" {
		c.OriginalName = other.OriginalName
	}
	if other.Description != "" {
		c.Description = other.Description
	}
	if other.ShortDescription != "" {
		c.ShortDescription = other.ShortDescription
	}
	if other.Year != 0 {
		c.Year = other.Year
	}
	if other.Images != "" {
		c.Images = other.Images
	}
	if other.Type == "movie" || other.Type == "tv_show" {
		c.Type = other.Type
	}
	if other.IsFree != nil {
		c.IsFree = other.IsFree
	}
	if len(other.Countries) > 0 {
		c.Countries = other.Countries
	}
	if len(other.Genres) > 0 {
		c.Genres = other.Genres
	}
	if len(other.Actors) > 0 {
		c.Actors = other.Actors
	}
	if len(other.Directors) > 0 {
		c.Directors = other.Directors
	}
}
