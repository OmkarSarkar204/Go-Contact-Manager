package main

import "fmt"

// Contact struct to store contact details
type Contact struct {
    Name  string
    Email string
    Phone string
}

// addContact adds a new contact to the slice and returns the updated slice
func addContact(contacts []Contact, name, email, phone string) []Contact {
    newContact := Contact{
        Name:  name,
        Email: email,
        Phone: phone,
    }
    return append(contacts, newContact)
}

// listContacts prints all contacts in the slice
func listContacts(contacts []Contact) {
    if len(contacts) == 0 {
        fmt.Println("No contacts found.")
        return
    }
    for i, contact := range contacts {
        fmt.Printf("Contact %d: %s, %s, %s\n", i, contact.Name, contact.Email, contact.Phone)
    }
}

// findContact searches for a contact by name
func findContact(contacts []Contact, name string) *Contact {
    for _, contact := range contacts {
        if contact.Email == "" {
            continue // Skip contacts with empty email
        }
        if contact.Name == name {
            return &contact // Return pointer to contact
        }
    }
    return nil // Not found
}

// deleteContact removes a contact by name
func deleteContact(contacts []Contact, name string) []Contact {
    for i, contact := range contacts {
        if contact.Name == name {
            return append(contacts[:i], contacts[i+1:]...)
        }
    }
    return contacts // Return unchanged if not found
}

func main() {
    var contacts []Contact
    // Add contacts
    contacts = addContact(contacts, "Alice", "alice@example.com", "123-456-7890")
    contacts = addContact(contacts, "Bob", "", "987-654-3210")
    contacts = addContact(contacts, "Charlie", "charlie@example.com", "555-555-5555")

    // List all contacts
    fmt.Println("All Contacts:")
    listContacts(contacts)

    // Find a contact
    fmt.Println("\nFinding Alice:")
    found := findContact(contacts, "Alice")
    if found != nil {
        fmt.Printf("Found: %s, %s, %s\n", found.Name, found.Email, found.Phone)
    } else {
        fmt.Println("Contact not found")
    }

    // Delete a contact
    fmt.Println("\nDeleting Bob:")
    contacts = deleteContact(contacts, "Bob")
    listContacts(contacts)
}