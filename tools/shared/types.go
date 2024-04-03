package shared

import "time"

// Maps where a file in an archive file is to where it should be moved in the blog
type MediaMap map[string]string

// Represents the frontmatter of a post
type FrontMatter struct {
	Title      string      `yaml:"title,omitempty"`
	Media      []MediaItem `yaml:"media,omitempty"`
	Date       string
	Tags       []string    `yaml:"tags,omitempty"`
	BookmarkOf string      `yaml:"bookmarkOf,omitempty"`
	References []Reference `yaml:"references,omitempty"`
}

type Reference struct {
	URL    string    `yaml:"url"`
	Type   string    `yaml:"type"`
	Name   string    `yaml:"name"`
	Author string    `yaml:"author"`
	Date   time.Time `yaml:"date"`
}

type MediaItem struct {
	URL string `yaml:"url"`
	Alt string `yaml:"alt,omitempty"`
}

type Post struct {
	Path        string
	PostFile    string
	FrontMatter FrontMatter
}
