package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exist := contactIndexByName[name]; exist {
		fmt.Println("Contact Already exist ", name)
	}
	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	contactList = append(contactList, newContact)
	contactIndexByName[name] = newContact.ID
	nextID++
	println("Contact Added ", name)

}

func findContactbyName(name string) *Contact {
	id, exsist := contactIndexByName[name]
	if !exsist {
		println("Contact Not Found ", name)
		return nil
	}
	println("Contact Found ", name)
	return &contactList[id-1]

}
func updateContact(id int, name, email, phone string) bool {
	idx := id - 1
	if idx < 0 || idx >= len(contactList) {
		return false
	}

	contactList[idx].Name = name
	contactList[idx].Email = email
	contactList[idx].Phone = phone
	println("Contact Updated ", id)
	return true
}
func listContacts() {
	if len(contactList) == 0 {
		fmt.Println("Contact List is Empty")
		return
	}
	fmt.Println("Contact List:")
	for _, contact := range contactList {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	}
}
func deleteContact(id int) bool {
	idx := id - 1
	if idx < 0 || idx >= len(contactList) {
		println("Contact Not Found ", id)
		return false
	}
	contactList = append(contactList[:idx], contactList[idx+1:]...)
	println("Contact Deleted ", id)
	return true
}
///
func main() {
	addContact("mohamed", "[EMAIL_ADDRESS]", "01012345678")
	addContact("kareem", "[EMAIL_ADDRESS]", "01012345679")
	listContacts()
	findContactbyName("mohamed")
	updateContact(1, "mohamed", "[EMAIL_ADDRESS]", "01012345678")
	listContacts()
	deleteContact(1)
	listContacts()

}
