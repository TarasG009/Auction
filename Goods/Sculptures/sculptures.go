package Sculptures

import (
	"fmt"
	"time"
)

type Sculpture struct {
	id      uint
	author  string
	date    time.Time
	height  uint
	width   uint
	article byte
	info    string
}

func (s *Sculpture) GetSculptures() []*Sculpture {
	return nil
}

func (s *Sculpture) GetSculpturesByOwner(sculptureOwner string) []*Sculpture {
	return nil
}

func (s *Sculpture) AddSculptures(sculptureId uint, sculptureAuthor string, sculptureDate time.Time, sculptureHeight uint, sculptureWidth uint, sculptureArticle byte, sculptureInfo string) {
	newSculpture := Sculpture{
		id:      sculptureId,
		author:  sculptureAuthor,
		date:    sculptureDate,
		height:  sculptureHeight,
		width:   sculptureWidth,
		article: sculptureArticle,
		info:    sculptureInfo,
	}
	fmt.Println(newSculpture)

}

func (s *Sculpture) DeleteSculptures(sculptureOwner string, sculptureId uint) {
}
