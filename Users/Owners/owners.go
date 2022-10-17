package Owners

import "fmt"

type Owner struct {
	id          uint
	FirstName   string
	LastName    string
	Mail        string
	PhoneNumber string
	Activity    bool
}

func (o *Owner) GetOwners() []*Owner {
	return nil
}

func (o *Owner) GetOwnerById(OwnerId uint) *Owner {
	return &Owner{
		id:          0,
		FirstName:   "",
		LastName:    "",
		Mail:        "",
		PhoneNumber: "",
		Activity:    false,
	}
}

func (o *Owner) AddOwners(ownerId uint, ownerFirstName string, ownerLastName string, ownerMail string, ownerPhoneNumber string, ownerActivity bool) {
	newOwner := Owner{
		id:          ownerId,
		FirstName:   ownerFirstName,
		LastName:    ownerLastName,
		Mail:        ownerMail,
		PhoneNumber: ownerPhoneNumber,
		Activity:    ownerActivity,
	}
	fmt.Println(newOwner)
}

func (o *Owner) DeleteOwners(ownerId uint) {

}
