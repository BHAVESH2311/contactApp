package contact

import (
	contactdetails "contactapp/contact_details"
	"contactapp/validators"
	"fmt"
)

type Contact struct {
	contactID      int
	firstName      string
	lastName       string
	isActive       bool
	contactDetails []*contactdetails.ContactDetail
}

// Contact factory
func NewContact(contactId int, firstName, lastName string) *Contact {
	var newContact *Contact
	newContact = &Contact{
		contactID:      contactId,
		firstName:      firstName,
		lastName:       lastName,
		isActive:       true,
		contactDetails: nil,
	}
	return newContact
}

// Create Contact, calls the factory
func CreateContact(contacts *[]*Contact, firstName, lastName string) {
	newContactId := len(*contacts) + 1
	createdContact := NewContact(newContactId, firstName, lastName)
	*contacts = append(*contacts, createdContact)
}

// Prints the contact
func (c *Contact) PrintContact() {
	fmt.Println("-----------------")
	fmt.Println("Contact Id: ", c.contactID)
	fmt.Println("First Name: ", c.firstName)
	fmt.Println("Last Name: ", c.lastName)
	fmt.Println("Is Active: ", c.isActive)
	fmt.Println("Contact Details: ", c.contactDetails)
	fmt.Println("-----------------")
}

// Read by Contact Id
func ReadContactByContactId(contacts []*Contact, contactId int) *Contact {
	for i := 0; i < len(contacts); i++ {
		if contacts[i].contactID == contactId && contacts[i].isActive {
			return contacts[i]
		}
	}
	return nil
}

// Read All
func ReadAllContacts(contacts []*Contact) {
	for i := 0; i < len(contacts); i++ {
		if contacts[i].isActive {
			contacts[i].PrintContact()
		}
	}
}

// Update functions
func (c *Contact) updateFirstName(firstName string) {
	c.firstName = firstName
}

func (c *Contact) updateLastName(lastName string) {
	c.lastName = lastName
}

func UpdateContactById(contacts []*Contact, contactId int) {
	var contactToUpdate *Contact
	contactToUpdate = ReadContactByContactId(contacts, contactId)
	if contactToUpdate == nil || !contactToUpdate.isActive {
		fmt.Println("Contact does not exist")
		return
	}

	choice := -1
	fmt.Println(" Enter the field you wish to update")
	fmt.Println("2. First Name")
	fmt.Println("3. Last Name")
	fmt.Scan(&choice)

	if choice == 2 {
		firstName := ""
		fmt.Println("Enter new first name: ")
		fmt.Scan(&firstName)

		if !validators.ValidateName(firstName) {
			fmt.Println("Name Validation failed")
		}
		contactToUpdate.updateFirstName(firstName)
	} else if choice == 3 {
		lastName := ""
		fmt.Println("Enter new last name: ")
		fmt.Scan(&lastName)
		if !validators.ValidateName(lastName) {
			fmt.Println("Name Validation failed")
		}
		contactToUpdate.updateLastName(lastName)
	} else {
		fmt.Println("Invalid Choice")
	}
}

// Delete contact
func DeleteContact(contacts []*Contact, contactId int) {
	var contactToDelete *Contact
	contactToDelete = ReadContactByContactId(contacts, contactId)
	if contactToDelete == nil || !contactToDelete.isActive {
		fmt.Println("Contact does not exist")
		return
	}
	contactToDelete.isActive = false
}

// Create contact detail, calls the Contact detail factory
func CreateContactDetail(contacts *[]*Contact, contactId int, typeContactDetail string, contactInfo string) {
	var resultContact *Contact
	resultContact = ReadContactByContactId(*contacts, contactId)
	if resultContact == nil {
		fmt.Println("Contact does not exist")
		return
	}
	if !resultContact.isActive {
		fmt.Println("Contact is not Active")
		return
	}
	contactdetails.CreateContactDetail(&resultContact.contactDetails, typeContactDetail, contactInfo)
}

// Read all contact details for a contact
func ReadAllContactDetails(contacts []*Contact, contactId int) {
	var resultContact *Contact
	resultContact = ReadContactByContactId(contacts, contactId)
	if resultContact == nil {
		fmt.Println("Contact does not exist")
		return
	}
	if !resultContact.isActive {
		fmt.Println("Contact is not Active")
		return
	}
	contactdetails.ReadAllContactDetails(resultContact.contactDetails)
}

// Read contact detail by Id
func ReadContactDetailById(contacts []*Contact, contactId int) *contactdetails.ContactDetail {
	resultContact := ReadContactByContactId(contacts, contactId)
	if resultContact == nil {
		fmt.Println("Contact does not exist")
		return nil
	}
	if !resultContact.isActive {
		fmt.Println("Contact is not Active")
		return nil
	}
	var contactDetailId int
	fmt.Println("Enter contact Detail Id: ")
	fmt.Scan(&contactDetailId)

	var resultContactDetail *contactdetails.ContactDetail
	resultContactDetail = contactdetails.ReadContactDetailsById(resultContact.contactDetails, contactDetailId)

	return resultContactDetail
}

// Update contact Detail By Id
func UpdateContactDetailById(contacts []*Contact, contactId int) {
	resultContact := ReadContactByContactId(contacts, contactId)
	if resultContact == nil {
		fmt.Println("Contact does not exist")
		return
	}
	if !resultContact.isActive {
		fmt.Println("Contact is not Active")
		return
	}
	var contactDetailId int
	fmt.Println("Enter contact Detail Id to Update: ")
	fmt.Scan(&contactDetailId)

	contactdetails.UpdateContactDetailById(resultContact.contactDetails, contactDetailId)

}

// Delete Contact Detail by Id
func DeleteContactDetailById(contacts []*Contact, contactId int) {
	resultContact := ReadContactByContactId(contacts, contactId)
	if resultContact == nil {
		fmt.Println("Contact does not exist")
		return
	}
	if !resultContact.isActive {
		fmt.Println("Contact is not Active")
		return
	}
	var contactDetailId int
	fmt.Println("Enter contact Detail Id to Delete: ")
	fmt.Scan(&contactDetailId)

	contactdetails.DeleteContactDetailById(resultContact.contactDetails, contactDetailId)
}
