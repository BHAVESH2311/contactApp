package contactdetails

import (
	"contactapp/validators"
	"fmt"
)

type ContactDetail struct {
	contactDetailId   int
	typeContactDetail string
	contactInfo       string
	isActive          bool
}

// Factory
func newContactDetail(contactDetailId int, typeContactDetail string, contactInfo string) *ContactDetail {
	var newContactDetail *ContactDetail
	newContactDetail = &ContactDetail{
		contactDetailId:   contactDetailId,
		typeContactDetail: typeContactDetail,
		contactInfo:       contactInfo,
		isActive:          true,
	}
	return newContactDetail
}

// Create Contact Detail, Calls the factory
func CreateContactDetail(contactDetails *[]*ContactDetail, typeContactDetail string, contactInfo string) {
	newContactDetailId := len(*contactDetails) + 1
	createdContactDetail := newContactDetail(newContactDetailId, typeContactDetail, contactInfo)
	*contactDetails = append(*contactDetails, createdContactDetail)
	fmt.Println("Contact Detail Created Successfully")
}

// Print Contact Detail
func (cd *ContactDetail) PrintContactDetail() {
	fmt.Println("-----------------")
	fmt.Println("ContactDetail Id: ", cd.contactDetailId)
	fmt.Println("ContactDetail Type: ", cd.typeContactDetail)
	fmt.Println("ContactDetail Info: ", cd.contactInfo)
	fmt.Println("Is Active: ", cd.isActive)
	fmt.Println("-----------------")
	fmt.Println()
}

// Read all Contact Details
func ReadAllContactDetails(allContactdetails []*ContactDetail) {
	for i := 0; i < len(allContactdetails); i++ {
		if allContactdetails[i].isActive {
			allContactdetails[i].PrintContactDetail()
		}
	}
}

// Read Contact detail by id
func ReadContactDetailsById(allContactdetails []*ContactDetail, contactDetailId int) *ContactDetail {
	for i := 0; i < len(allContactdetails); i++ {
		if allContactdetails[i].contactDetailId == contactDetailId && allContactdetails[i].isActive {
			return allContactdetails[i]
		}
	}
	return nil
}

// Update Functions
func (cd *ContactDetail) updateTypeOfContactDetail(typeOfContact string) {
	cd.typeContactDetail = typeOfContact
}

func (cd *ContactDetail) updateContactInfo(contactInfo string) {
	cd.contactInfo = contactInfo
}

// Update contact detail by id
func UpdateContactDetailById(allContactdetails []*ContactDetail, contactDetailId int) {
	var contactDetailToUpdate *ContactDetail
	contactDetailToUpdate = ReadContactDetailsById(allContactdetails, contactDetailId)
	if contactDetailToUpdate == nil {
		fmt.Println("Contact Detail does not exist")
		return
	}

	var typeOfContact string
	var contactInfo string
	fmt.Println("Enter updated type of contact (Type 'phone' or 'email'): ")
	fmt.Scan(&typeOfContact)

	fmt.Println("Enter updated contact: ")
	fmt.Scan(&contactInfo)

	isValid, err := validators.ValidateContactDetails(typeOfContact, contactInfo)
	if !isValid {
		fmt.Println(err)
		return
	}

	contactDetailToUpdate.updateTypeOfContactDetail(typeOfContact)
	contactDetailToUpdate.updateContactInfo(contactInfo)
}

// Delete Contact Detail
func DeleteContactDetailById(allContactdetails []*ContactDetail, contactDetailId int) {
	var contactDetailToDelete *ContactDetail
	contactDetailToDelete = ReadContactDetailsById(allContactdetails, contactDetailId)
	if contactDetailToDelete == nil {
		fmt.Println("Contact Detail does not exist")
		return
	}
	contactDetailToDelete.isActive = false
}
