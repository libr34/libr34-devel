package models

import (
	"encoding/xml"
)

type Posts struct {
	XMLName xml.Name `xml:"posts"`
	Count   int      `xml:"count,attr"`
	Offset  int      `xml:"offset,attr"`
	PostS   []Post   `xml:"post"`
}
type Post struct {
	Height        int    `xml:"height,attr"`
	Score         int    `xml:"score,attr"`
	FileURL       string `xml:"file_url,attr"`
	ParentID      string `xml:"parent_id,attr"`
	SampleURL     string `xml:"sample_url,attr"`
	SampleWidth   int    `xml:"sample_width,attr"`
	SampleHeight  int    `xml:"sample_height,attr"`
	PreviewURL    string `xml:"preview_url,attr"`
	Rating        string `xml:"rating,attr"`
	Tags          string `xml:"tags,attr"`
	ID            int    `xml:"id,attr"`
	Width         int    `xml:"width,attr"`
	Change        int    `xml:"change,attr"`
	MD5           string `xml:"md5,attr"`
	CreatorID     int    `xml:"creator_id,attr"`
	HasChildren   bool   `xml:"has_children,attr"`
	CreatedAt     string `xml:"created_at,attr"`
	Status        string `xml:"status,attr"`
	Source        string `xml:"source,attr"`
	HasNotes      bool   `xml:"has_notes,attr"`
	HasComments   bool   `xml:"has_comments,attr"`
	PreviewWidth  int    `xml:"preview_width,attr"`
	PreviewHeight int    `xml:"preview_height,attr"`
}
