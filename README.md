## Tasks API
### Description
This is a simple API that allows you to create, read, update and delete tasks. It is built using Go and the Gorilla Mux framework. It uses PostgreSQL database to store the tasks. The API is RESTful and uses JSON for communication.

### Endpoints
#### Get all tasks
```http
GET /api/v1/tasks?page=1&size=10
```
This endpoint returns all the tasks in the database.

#### Get a single task
```http
GET /api/v1/tasks/{id}
```
This endpoint returns a single task with the specified id.

#### Create a task
```http
POST /api/v1/tasks
```
This endpoint creates a new task.

#### Update a task
```http
PUT /api/v1/tasks/{id}
```
This endpoint updates a task with the specified id.

#### Delete a task
```http
DELETE /api/v1/tasks/{id}
```
This endpoint deletes a task with the specified id.

### Installation
1. Clone the repository
```bash
git clone "repository-url"
```
2. Change directory to the project folder
```bash
cd "project-folder"
```
3. Create a `.env` file in the root of the project and add the following environment variables
```env
DB_USER=your-db-username
DB_PASSWORD=your-db-password
DB_NAME=your-db-name
DB_HOST=your-db-host
DB_PORT=your-db-port
```
4. Run the application
```bash
make run
```
5. The application will be running on `http://localhost:8080`
6. You can now use the API to create, read, update and delete tasks.
7. You can also run the tests using the command
```bash
make test
```
### Note: Coming soon
8. You can also run the tests with coverage using the command
```bash
make cover
```
9. You can also run the tests with coverage and generate an HTML report using the command
```bash
make cover-html
```
10. You can also run the linter using the command
```bash
make lint
```
11. You can also run the linter and fix the issues using the command
```bash
make lint-fix
```

