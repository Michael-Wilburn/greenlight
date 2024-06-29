
# Greenlight

Greenlight — a JSON API for retrieving and managing information about movies. You can
think of the core functionality as being a bit like the Open Movie Database API.


#### Ultimately, our Greenlight API will support the following endpoints and actions:


| Method  | URL Pattern     | Action               |
| :-------- | :------- | :------------------------- |
| `GET`  | `/v1/healthcheck` | Show application health and version information |
| `GET`  | `/v1/movies` | Show the details of all movies |
| `POST` | `/v1/movies` | Create a new movie |
| `GET`  | `/v1/movies/:id` | Show the details of a specific movie |
| `PATCH`| `/v1/movies/:id` | Update the details of a specific movie |
| `DELETE`| `/v1/movies/:id` | Delete a specific movie |
| `POST` | `/v1/users` | Register a new user |
| `PUT`  | `/v1/users/activated` | Activate a specific user |
| `PUT`  | `/v1/users/password` | Update the password for a specific user |
| `POST` | `/v1/tokens/authentication` | Generate a new authentication token |
| `POST` | `/v1/tokens/password-reset` | Generate a new password-reset token |
| `GET`  | `/debug/vars` | Display application metrics |


## Skeleton directory structure
.
├── bin
├── cmd
│ └── api
│ └── main.go
├── internal
├── migrations
├── remote
├── go.mod
└── Makefile

Let’s take a moment to talk through these files and folders and explain the purpose that
they’ll serve in our finished project.


* The bin directory will contain our compiled application binaries, ready for deployment to a production server.
* The cmd/api directory will contain the application-specific code for our Greenlight API application. This will include the code for running the server, reading and writing HTTP requests, and managing authentication.
* The internal directory will contain various ancillary packages used by our API. It will contain the code for interacting with our database, doing data validation, sending emails and so on. Basically, any code which isn’t application-specific and can potentially be reused will live in here. Our Go code under cmd/api will import the packages in the internal directory (but never the other way around).
* The migrations directory will contain the SQL migration files for our database.
* The remote directory will contain the configuration files and setup scripts for our production server.
* The go.mod file will declare our project dependencies, versions and module path.
* The Makefile will contain recipesfor automating common administrative tasks — like auditing our Go code, building binaries, and executing database migrations.