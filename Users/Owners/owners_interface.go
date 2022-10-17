package Owners

type OwnersI interface {
	GetOwners() []*Owner
	GetOwnerById(OwnerId uint) *Owner
	AddOwners(ownerId uint, ownerFirstName string, ownerLastName string, ownerMail string, ownerPhoneNumber string, ownerActivity bool)
	DeleteOwners(ownerId uint)
}
