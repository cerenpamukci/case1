Hello, this project has been created for the FillLabs technical test. The project primarily involves working with Go for the backend, while React and Next.js will be used on the frontend. The system will be supported by REST APIs.

---    

## Tasks

### Q1: Sort Words by Character "a"   
**Task:**  
Write a function to sort a list of words by the number of occurrences of the character "a" (in descending order). If two words have the same number of "a"s, sort them by length.

**Example Input:**  
`["aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"]`

**Example Output:**  
`["aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"]`

---

### Q2: Recursive Function
**Task:**  
Write a recursive function that takes an integer parameter. The output should demonstrate a specific pattern based on the algorithm.

**Example Input:**  
`9`

**Example Output:**  

### Q3: Find Most Repeated Data
**Task:**  
Write a function to find the most repeated element in an array.

**Example Input:**  
`["apple", "pie", "apple", "red", "red", "red"]`

**Example Output:**  
`"red"`

---

### Q4: User Management Project

**Description:**  
Develop a user management system with the following requirements:

1. **Frontend (React & Next.js with TypeScript):**
   - **Master View:** Displays all users in a data grid with CRUD functionality.  
     - Buttons: `New`, `Edit`, `Delete`.
     - `Edit` and `Delete` require row selection.
   - **Detailed View:** A form for creating, editing, or deleting a user.
     - Buttons:  
       - **Action Button:** Text changes dynamically:  
         - `New → Create`  
         - `Edit → Save`  
         - `Delete → Delete`  
       - **Back Button:** Returns to the master view.

2. **Backend (Go & SQLite):**
   - Implement a REST API with the following endpoints:
     - `GET /users` → Returns all users.
     - `GET /users/{id}` → Returns a user by ID.
     - `POST /users` → Creates a new user.
     - `PUT /users/{id}` → Updates an existing user.
     - `DELETE /users/{id}` → Deletes a user by ID.
   - Use SQLite for data persistence.

**Note:** Include the SQLite database file in the project folder. Follow REST API standards for paths, methods, and status codes.

---
