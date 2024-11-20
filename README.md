
# ViewLens Backend

ViewLens Backend is a REST API designed to dynamically fetch and display MSSQL views and their contents in JSON format. Built using the Gin framework, it allows users to query views based on a specified prefix.

---

## Features
- Dynamically list views starting with a specified prefix (`VIEW_PREFIX`).
- Retrieve columns and data of a specific view.
- Secure connection to MSSQL database.

---

## Setup

### 1. Clone the Repository
```sh
git clone https://github.com/yasirsey/viewlens-backend.git
cd viewlens-backend
```

### 2. Install Dependencies
```sh
go mod tidy
```

### 3. Configure the `.env` File
```env
DB_USER=your_db_user # Database username 
DB_PASSWORD=your_db_password # Database password 
DB_SERVER=your_db_server # Database server address (e.g., localhost)
DB_PORT=1433 # Database port (default: 1433) 
DB_DATABASE=your_db_name # Database name 
VIEW_PREFIX=VEX_ # Prefix for dynamic views 
APP_PORT=8080 # API server port
```

### 4. Run the Application
`go run cmd/main.go`

## API Usage
### 1. List Views

 - **Endpoint:** `GET /views`
 - **Description:** Returns a list of views starting with the specified `VIEW_PREFIX`.
 - **Response Example:**
 ```json
 { "views": ["VEX_Users", "VEX_Orders"] }
 ```

### 2. View Details

 - **Endpoint:** `GET /view/:viewName`
 - **Description:** Returns the column names and data of the specified view.
 - **Parameters:**
- - `viewName` (path param): Name of the view to retrieve.
- **Response Example:**
```json
{ 
	"columns": ["ID", "Name", "Email"], 
	"data": [ 
		{"ID": 1, "Name": "Alice", "Email": "alice@example.com"}, 
		{"ID": 2, "Name": "Bob", "Email": "bob@example.com"} 
	] 
}
```
## Project Structure
```plaintext
viewlens-backend/ 
├── cmd/ # Application entry point (main.go) 
├── config/ # Configuration and environment handling 
├── db/ # Database connection settings 
├── internal/ # Core application logic 
│ ├── controllers/ # Request handlers 
│ ├── services/ # Business logic layer 
│ ├── repositories/ # Database operations 
│ └── models/ # Data models 
├── .env # Environment variables 
├── go.mod # Go module file 
└── README.md # Project documentation
```
## Contribution
If you'd like to contribute to this project, please fork the repository and submit a pull request. Suggestions, improvements, and bug reports are always welcome.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.

