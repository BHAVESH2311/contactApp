package user

import (
	"contactapp/contact"
	contactdetails "contactapp/contact_details"
	"contactapp/validators"
	"fmt"
)

var users []*User
var admin *User

type User struct {
	userId    int
	firstName string
	lastName  string
	isAdmin   bool
	isActive  bool
	contacts  []*contact.Contact
}

// User factory
func (u *User) newUser(firstName, lastName string) *User {
	newUserId := len(users) + 1
	var newUser = &User{
		firstName: firstName,
		lastName:  lastName,
		userId:    newUserId,
		isAdmin:   false,
		isActive:  true,
		contacts:  nil,
	}
	return newUser
}

// Admin (has Id as -1)
func newAdmin(firstName, lastName string) *User {
	var newAdmin = &User{
		firstName: firstName,
		lastName:  lastName,
		userId:    -1,
		isAdmin:   true,
		isActive:  true,
		contacts:  nil,
	}

	return newAdmin
}

func GetAdmin() *User {
	if admin != nil {
		return admin
	}
	return nil
}

// Create user, takes input and calls the user factory
func (u *User) CreateUser() {
	if !u.isAdmin {
		return
	}
	var firstName string
	var lastName string
	fmt.Println("Enter first name for new user: ")
	fmt.Scan(&firstName)
	isValid := validators.ValidateName(firstName)
	if !isValid {
		fmt.Println("Name validation failed")
		return
	}

	fmt.Println("Enter last name for new user: ")
	fmt.Scan(&lastName)
	isValid = validators.ValidateName(lastName)
	if !isValid {
		fmt.Println("Name validation failed")
		return
	}

	var newUser *User = u.newUser(firstName, lastName)

	users = append(users, newUser)

	fmt.Println("User created successfully")
}

// Create Admin takes input and calls the factory
func CreateAdmin() {
	if admin != nil {
		fmt.Println("Admin exists")
		return
	}
	var firstName string
	var lastName string
	fmt.Println("Enter first name for new user: ")
	fmt.Scan(&firstName)
	isValid := validators.ValidateName(firstName)
	if !isValid {
		fmt.Println("Name validation failed")
		return
	}

	fmt.Println("Enter last name for new user: ")
	fmt.Scan(&lastName)
	isValid = validators.ValidateName(lastName)
	if !isValid {
		fmt.Println("Name validation failed")
		return
	}

	var newAdmin *User = newAdmin(firstName, lastName)

	admin = newAdmin
	fmt.Println("Admin created successfully")
}

// Print user
func (u *User) PrintUser() {
	fmt.Println("-----------------")
	fmt.Println("User Id: ", u.userId)
	fmt.Println("First Name: ", u.firstName)
	fmt.Println("Last Name: ", u.lastName)
	fmt.Println("Is Admin: ", u.isAdmin)
	fmt.Println("Is Active: ", u.isActive)
	// fmt.Println("Contacts: ", u.contacts)
	fmt.Println("-----------------")
}

// Read all users
func (u *User) ReadAllUsers() {
	if !u.isAdmin {
		return
	}
	for i := 0; i < len(users); i++ {
		if users[i].isActive {
			users[i].PrintUser()
		}
	}
}

// Read user by id
func (u *User) ReadUserById(id int) (*User, string) {
	if !u.isAdmin && !u.isActive {
		return nil, "Not an Admin"
	}
	for i := 0; i < len(users); i++ {
		if users[i].userId == id && users[i].isActive {
			return users[i], "User Found"
		}
	}
	return nil, "UserId does not exist"
}

// Get user by id for Staff
func GetUserById(uId int) (*User, string) {
	for i := 0; i < len(users); i++ {
		if users[i].userId == uId && users[i].isActive {
			return users[i], "Found"
		}
	}
	return nil, "UserId does not exist"
}

// Update functions
func (u *User) updateFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) updateLastName(lastName string) {
	u.lastName = lastName
}

// Update User
func (u *User) UpdateUserById(userId int) string {
	if !u.isAdmin && !u.isActive {
		return "Thid ID is Not an admin"
	}
	user, msgString := u.ReadUserById(userId)
	if user == nil {
		return msgString
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
			return "Name Validation failed"
		}
		user.updateFirstName(firstName)
	} else if choice == 3 {
		lastName := ""
		fmt.Println("Enter new last name: ")
		fmt.Scan(&lastName)
		if !validators.ValidateName(lastName) {
			return "Name Validation failed"
		}
		user.updateLastName(lastName)
	} else {
		return "Invalid choice"
	}
	return "Successfully updated the student"
}

// Delete User
func (u *User) DeleteUserById(userId int) string {
	if !u.isAdmin && !u.isActive {
		return "Not an admin/ Not Active Admin"
	}
	user, msgString := u.ReadUserById(userId)
	if user == nil {
		return msgString
	}

	user.isActive = false
	return "Deleted successfully"
}

// CRUD for Contacts

// Create contact
func (u *User) CreateContact() string {
	if u.isAdmin || !u.isActive {
		return "Admin cannot Create a Contact/ User is not Active"
	}
	var firstName string
	var lastName string
	fmt.Println("Enter first name for new Contact: ")
	fmt.Scan(&firstName)
	if !validators.ValidateName(firstName) {
		return "Name Validation failed"
	}

	fmt.Println("Enter last name for new Contact: ")
	fmt.Scan(&lastName)
	if !validators.ValidateName(lastName) {
		return "Name Validation failed"
	}

	contact.CreateContact(&u.contacts, firstName, lastName)
	return ("Contact created successfully")
}

// Read all contacts
func (u *User) ReadAllContacts() {
	if u.isAdmin || !u.isActive {
		fmt.Println("Admin cannot Read a Contact/ User is not Active")
	}
	contact.ReadAllContacts(u.contacts)
}

// Read contact by Id
func (u *User) ReadContactById(contactId int) *contact.Contact {
	if u.isAdmin || !u.isActive {
		fmt.Println("Admin cannot Read a Contact/ User is not Active")
	}
	var resultContact *contact.Contact
	resultContact = contact.ReadContactByContactId(u.contacts, contactId)
	return resultContact
}

// Update contact by id
func (u *User) UpdateContactById(contactId int) {
	if u.isAdmin || !u.isActive {
		fmt.Println("Admin cannot Read a Contact/ User is not Active")
		return
	}

	contact.UpdateContactById(u.contacts, contactId)
}

// Delete Contact by id
func (u *User) DeleteContactById(contactId int) {
	if u.isAdmin || !u.isActive {
		fmt.Println("Admin cannot Delete a Contact/ User is not Active")
		return
	}

	contact.DeleteContact(u.contacts, contactId)
}

// Create contact detail
func (u *User) CreateContactDetail() {
	if u.isAdmin || !u.isActive {
		fmt.Println("Admin cannot create a Contact Detail/ User is not Active")
		return
	}
	var typeOfContact string
	var contactInfo string
	var contactId int

	fmt.Println("Enter contact Id to add contact detail for: ")
	fmt.Scan(&contactId)

	fmt.Println("Enter type of contact to create (Type 'phone' or 'email'): ")
	fmt.Scan(&typeOfContact)

	fmt.Println("Enter contact: ")
	fmt.Scan(&contactInfo)

	isValid, err := validators.ValidateContactDetails(typeOfContact, contactInfo)
	if !isValid {
		fmt.Println(err)
		return
	}

	contact.CreateContactDetail(&u.contacts, contactId, typeOfContact, contactInfo)
}

// Read all contact details
func (u *User) ReadAllContactDetails() {
	if u.isAdmin || !u.isActive {
		fmt.Println("Admin cannot create a Contact Detail/ User is not Active")
		return
	}

	var contactId int
	fmt.Println("Enter contact Id to read all contact details for: ")
	fmt.Scan(&contactId)

	contact.ReadAllContactDetails(u.contacts, contactId)
}

// Read contact detail by ID
func (u *User) ReadContactDetailById() *contactdetails.ContactDetail {
	if u.isAdmin || !u.isActive {
		fmt.Println("Admin cannot create a Contact Detail/ User is not Active")
		return nil
	}

	var contactId int
	fmt.Println("Enter contact Id to read all contact details for: ")
	fmt.Scan(&contactId)

	var contactDetail *contactdetails.ContactDetail
	contactDetail = contact.ReadContactDetailById(u.contacts, contactId)

	return contactDetail
}

// Update contact detail by id
func (u *User) UpdateContactDetailById() {
	if u.isAdmin || !u.isActive {
		fmt.Println("Admin cannot update a Contact Detail/ User is not Active")
		return
	}

	var contactId int
	fmt.Println("Enter contact Id to update contact details for: ")
	fmt.Scan(&contactId)
	contact.UpdateContactDetailById(u.contacts, contactId)
}

// Delete contact detail by Id
func (u *User) DeleteContactDetailById() {
	if u.isAdmin || !u.isActive {
		fmt.Println("Admin cannot delete a Contact Detail/ User is not Active")
		return
	}

	var contactId int
	fmt.Println("Enter contact Id to update contact details for: ")
	fmt.Scan(&contactId)
	contact.DeleteContactDetailById(u.contacts, contactId)

}
