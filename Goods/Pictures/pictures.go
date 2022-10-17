package Pictures

import (
	"fmt"
	"time"
)

type Picture struct {
	id      uint
	author  string
	date    time.Time
	height  uint
	width   uint
	article byte
	info    string
}

func (p *Picture) GetPictures() []*Picture {
	return nil
}

func (p *Picture) GetPicturesByOwner(pictureOwner string) []*Picture {
	return nil
}

func (p *Picture) AddPictures(pictureId uint, pictureAuthor string, pictureDate time.Time, pictureHeight uint, pictureWidth uint, pictureArticle byte, pictureInfo string) {
	newPicture := Picture{
		id:      pictureId,
		author:  pictureAuthor,
		date:    pictureDate,
		height:  pictureHeight,
		width:   pictureWidth,
		article: pictureArticle,
		info:    pictureInfo,
	}

	fmt.Println(newPicture)

}

func (p *Picture) DeletePictures(pictureOwner string, pictureId uint) {

}
