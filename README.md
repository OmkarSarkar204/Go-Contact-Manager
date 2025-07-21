# Go Contact Manager

A command-line tool to manage contacts, built with Go. Ideal for small business CRMs or contact management.

## Features
- Interactively add, list, find, and delete contacts using a menu (options 1–5).
- Uses structs to model contact data (name, email, phone(as string)).
- Validates inputs to prevent empty names.
- Skips invalid contacts (e.g., empty emails) during searches.
- Suitable for small businesses or freelancers needing simple contact management.

## Setup
1. Install Go: https://go.dev/doc/install
2. Clone this repo: `git clone https://github.com/OmkarSarkar204/Go-Contact-Manager`
3. Navigate to the project: `cd Go-Contact-Manager`
4. Run: `go run main.go`

## Usage
- Run the program and choose options 1–5 from the menu.
- Example: Select "1" to add a contact like "Alice, alice@example.com, 123-456-7890".
- Note: Currently supports single-word names (e.g., "Alice"). Multi-word names will be added in future versions.


## Future Improvements
- Support multi-word names using `bufio` for input.
- Add file storage to save contacts persistently.
- Extend to a REST API for web-based access

## Author
Omkar Sarkar_
