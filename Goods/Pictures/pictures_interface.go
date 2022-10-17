package Pictures

import "time"

type Pictures interface {
	GetPictures() []*Picture
	GetPicturesByOwner(pictureOwner string) []*Picture
	AddPictures(id uint, author string, date time.Time, height uint, width uint, article byte, info string)
	DeletePictures(pictureOwner string, pictureId uint)
}
