package main

import (
	"contactapp/user"
	"fmt"
)

func main() {
	var choice int
Here:
	for {
		fmt.Println()
		fmt.Println("----Contact App--------")
		fmt.Println("Enter your choice")
		fmt.Println("1. Create Admin")
		fmt.Println("2. Sign in as Admin")
		fmt.Println("3. Sign in as User")
		fmt.Println("0. Exit")
		fmt.Scan(&choice)

		if choice == -1 {
			break
		} else if choice == 1 {
			user.CreateAdmin()
		} else if choice == 2 {
			// Use the system as Admin
			var admin *user.User
			admin = user.GetAdmin()
			if admin != nil {
				var choice int

				for {
					fmt.Println()
					fmt.Println("----You are signed in as Admin------")
					fmt.Println("Enter a choice")
					fmt.Println("1. Create user")
					fmt.Println("2. Read all users")
					fmt.Println("3. Read user by Id")
					fmt.Println("4. Update user by Id")
					fmt.Println("5. Delete user by Id")
					fmt.Println("0. Exit")
					fmt.Scan(&choice)
					if choice == 0 {
						break
					} else if choice == 1 {
						admin.CreateUser()
					} else if choice == 2 {
						admin.ReadAllUsers()
					} else if choice == 3 {
						var userId int
						fmt.Println("Enter userID: ")
						fmt.Scan(&userId)

						user, msg := admin.ReadUserById(userId)
						if user != nil {
							user.PrintUser()
						} else {
							fmt.Println(msg)
						}

					} else if choice == 4 {
						var userId int
						fmt.Println("Enter userID: ")
						fmt.Scan(&userId)

						user, msg := admin.ReadUserById(userId)
						if user != nil {
							admin.UpdateUserById(userId)
						} else {
							fmt.Println(msg)
						}
					} else if choice == 5 {
						var userId int
						fmt.Println("Enter userID: ")
						fmt.Scan(&userId)

						user, msg := admin.ReadUserById(userId)
						if user != nil {
							admin.DeleteUserById(userId)
						} else {
							fmt.Println(msg)
						}
					} else {
						fmt.Println("Invalid choice")
					}
				}
			} else {
				fmt.Println("Admin does not exist. Please create one")
			}
		} else if choice == 3 {
			var userId int
			var currentUser *user.User
			fmt.Println("Enter user ID to sign in to a user")
			fmt.Scan(&userId)

			currentUser, err := user.GetUserById(userId)
			if currentUser == nil {
				fmt.Println(err)
				goto Here
			}

			fmt.Println("----You are signed in as the below User------")
			currentUser.PrintUser()

			var choice int
			for {
				fmt.Println()
				fmt.Println("Enter a choice")
				fmt.Println("1. Create contact")
				fmt.Println("2. Read All contacts")
				fmt.Println("3. Read contact by id")
				fmt.Println("4. Update contact by id")
				fmt.Println("5. Delete contact by id")
				fmt.Println("6. Create contact Detail for a contact")
				fmt.Println("7. Read all Contact Details for a contact")
				fmt.Println("8. Read contact Details by Id for a contact")
				fmt.Println("9. Update contact Details by Id for a contact")
				fmt.Println("10. Delete contact Detail by Id for a contact")
				fmt.Println("0. Exit")
				fmt.Scan(&choice)

				if choice == 0 {
					break
				} else if choice == 1 {
					res := currentUser.CreateContact()
					fmt.Println(res)
				} else if choice == 2 {
					currentUser.ReadAllContacts()
				} else if choice == 3 {
					var contactId int
					fmt.Println("Enter contact id: ")
					fmt.Scan(&contactId)
					resultContact := currentUser.ReadContactById(contactId)
					if resultContact != nil {
						resultContact.PrintContact()
					} else {
						fmt.Println("Contact does not exist")
					}
				} else if choice == 4 {
					var contactId int
					fmt.Println("Enter contact id: ")
					fmt.Scan(&contactId)
					currentUser.UpdateContactById(contactId)
				} else if choice == 5 {
					var contactId int
					fmt.Println("Enter contact id: ")
					fmt.Scan(&contactId)
					currentUser.DeleteContactById(contactId)
				} else if choice == 6 {
					currentUser.CreateContactDetail()
				} else if choice == 7 {
					currentUser.ReadAllContactDetails()
				} else if choice == 8 {
					foundContactDetail := currentUser.ReadContactDetailById()
					if foundContactDetail != nil {
						foundContactDetail.PrintContactDetail()
					} else {
						fmt.Println("Contact Detail does not exist")
					}
				} else if choice == 9 {
					currentUser.UpdateContactDetailById()
				} else if choice == 10 {
					currentUser.DeleteContactDetailById()
				} else {
					fmt.Println("Invalid choice")
				}
			}

		}
	}
}
