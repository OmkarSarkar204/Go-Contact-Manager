package main

import "fmt"

// Contact struct to store contact details
type Contact struct {
    Name  string
    Email string
    Phone string
}

// addContact adds a new contact to the slice
func addContact(contacts []Contact, name, email, phone string) []Contact {
    if name == "" {
        fmt.Println("Error: Name cannot be empty")
        return contacts
    }
    newContact := Contact{
        Name:  name,
        Email: email,
        Phone: phone,
    }
    return append(contacts, newContact)
}

// listContacts prints all contacts
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
    for {
        fmt.Println("\nContact Manager:")
        fmt.Println("1. Add Contact")
        fmt.Println("2. List Contacts")
        fmt.Println("3. Find Contact")
        fmt.Println("4. Delete Contact")
        fmt.Println("5. Exit")
        fmt.Print("Choose an option (1-5): ")

        var choice int
        _, err := fmt.Scan(&choice)
        if err != nil || choice < 1 || choice > 5 {
            fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
            continue
        }

        switch choice {
        case 1:
            var name, email, phone string
            fmt.Print("Enter name (single word): ")
            fmt.Scan(&name)
            fmt.Print("Enter email: ")
            fmt.Scan(&email)
            fmt.Print("Enter phone: ")
            fmt.Scan(&phone)
            contacts = addContact(contacts, name, email, phone)
            fmt.Println("Contact added successfully!")

        case 2:
            fmt.Println("\nAll Contacts:")
            listContacts(contacts)

        case 3:
            var name string
            fmt.Print("Enter name to find: ")
            fmt.Scan(&name)
            found := findContact(contacts, name)
            if found != nil {
                fmt.Printf("Found: %s, %s, %s\n", found.Name, found.Email, found.Phone)
            } else {
                fmt.Println("Contact not found")
            }

        case 4:
            var name string
            fmt.Print("Enter name to delete: ")
            fmt.Scan(&name)
            beforeLen := len(contacts)
            contacts = deleteContact(contacts, name)
            if len(contacts) < beforeLen {
                fmt.Println("Contact deleted successfully!")
            } else {
                fmt.Println("Contact not found")
            }

        case 5:
            fmt.Println("Goodbye!")
            return

        default:
            fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
        }
    }
}