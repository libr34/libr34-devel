package models

import (
	"encoding/xml"
)

type Posts struct {
	XMLName xml.Name `xml:"posts" json:"name"`
	Count   int      `xml:"count,attr" json:"count"`
	Offset  int      `xml:"offset,attr" json:"offset"`
	PostS   []Post   `xml:"post" json:"posts"`
}
type Post struct {
	Height        int    `xml:"height,attr" json:"height"`
	Score         int    `xml:"score,attr" json:"score"`
	FileURL       string `xml:"file_url,attr" json:"file_url"`
	ParentID      string `xml:"parent_id,attr" json:"parent_id"`
	SampleURL     string `xml:"sample_url,attr" json:"sample_url"`
	SampleWidth   int    `xml:"sample_width,attr" json:"sample_width"`
	SampleHeight  int    `xml:"sample_height,attr" json:"sample_height"`
	PreviewURL    string `xml:"preview_url,attr" json:"preview_url"`
	Rating        string `xml:"rating,attr" json:"rating"`
	Tags          string `xml:"tags,attr" json:"tags"`
	ID            int    `xml:"id,attr" json:"id"`
	Width         int    `xml:"width,attr" json:"width"`
	Change        int    `xml:"change,attr" json:"change"`
	MD5           string `xml:"md5,attr" json:"md5"`
	CreatorID     int    `xml:"creator_id,attr" json:"creator_id"`
	HasChildren   bool   `xml:"has_children,attr" json:"has_children"`
	CreatedAt     string `xml:"created_at,attr" json:"created_at"`
	Status        string `xml:"status,attr" json:"status"`
	Source        string `xml:"source,attr" json:"source"`
	HasNotes      bool   `xml:"has_notes,attr" json:"has_notes"`
	HasComments   bool   `xml:"has_comments,attr" json:"has_comments"`
	PreviewWidth  int    `xml:"preview_width,attr" json:"preview_width"`
	PreviewHeight int    `xml:"preview_height,attr" json:"preview_height"`
}
