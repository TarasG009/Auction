package Sculptures

import "time"

type Sculptures interface {
	GetSculptures() []*Sculpture
	GetSculpturesByOwner(sculptureOwner string) []*Sculpture
	AddSculptures(Id uint, Author string, Date time.Time, Height uint, Width uint, Article byte, Info string)
	DeleteSculptures(sculptureOwner string, sculptureId uint)
}
