# project-students-go
API to manage Go students

Routes:
- GET /students - List all students
- GET /students?active=<true/false> - List all active/non-active students
- POST /students - Create student
- GET /students/:id - Get data from a specific student
- PUT /students/:id - Update data from a specific student
- DELETE /students/:id - Delete a specific student

Struct Student:
- Name
- CPF
- Email
- Age
- Active