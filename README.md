# Greenlight
Greenlight — a JSON API for retrieving and managing information about movies. You can think of the core functionality as being a bit like the Open Movie Database API.

| **Method** | **URL Pattern**           | **Handler**        | **Action**                                      |
|------------|---------------------------|--------------------|-------------------------------------------------|
| GET        | /v1/healthcheck           | healthcheckHandler | Show application health and version information |
| GET        | /v1/movies                | listMoviesHandler  | Show the details of all movies                  |
| POST       | /v1/movies                | createMovieHandler | Create a new movie                              |
| GET        | /v1/movies/:id            | showMovieHandler   | Show the details of a specific movie            |
| PATCH      | /v1/movies/:id            | editMovieHandler   | Update the details of a specific movie          |
| DELETE     | /v1/movies/:id            | deleteMovieHandler | Delete a specific movie                         |
| POST       | /v1/users                 |                    | Register a new user                             |
| PUT        | /v1/users/activated       |                    | Activated a specific user                       |
| PUT        | /v1/users/password        |                    | Update the password for a specific user         |
| POST       | /v1/tokens/authentication |                    | Generate a new authentication token             |
| POST       | /v1/tokens/password-reset |                    | Generate a new password-reset token             |
| GET        | /debug/vars               |                    | Display application metrics                     |


# HTTP Verbs in REST API

| **Method** 	| **Usage**                                                                                                                                                                                                	|
|------------	|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| **GET**    	| Use for actions that retrieve information only and don’t change the state of your<br>application or any data.                                                                                            	|
| **POST**   	| Use for non-idempotent actions that modify state. In the context of a REST API, POST<br>is generally used for actions that _create_ a new resource.                                                      	|
| **PUT**    	| Use for idempotent actions that modify the state of a resource at a specific URL. In<br>the context of a REST API, PUT is generally used for actions that _replace_ or _update_<br>an existing resource. 	|
| **PATCH**  	| Use for actions that _partially update_ a resource at a specific URL. It’s OK for the<br>action to be either idempotent or non-idempotent.                                                               	|
| **DELETE** 	| Use for actions that _delete_ a resource at a specific URL.                                                                                                                                              	|

# JSON Encoding
The following table summarizes how different Go types are mapped to JSON data types
during encoding:

| **Go Type**                                        	| **⇒** 	| **JSON type**              	|
|----------------------------------------------------	|-------	|----------------------------	|
| bool                                               	| ⇒     	| JSON boolean               	|
| string                                             	| ⇒     	| JSON string                	|
| int*, uint*, float*, rune                          	| ⇒     	| JSON number                	|
| array, slice                                       	| ⇒     	| JSON array                 	|
| struct, map                                        	| ⇒     	| JSON object                	|
| nil pointers, interface values, slices, maps, etc. 	| ⇒     	| JSON null                  	|
| chan, func, complex*                               	| ⇒     	| Not supported              	|
| time.Time                                          	| ⇒     	| RFC3339-format JSON string 	|
| []byte                                             	| ⇒     	| Base64-encoded JSON string 	|


# Triaging the Decode error

| **Error types** | **Reason** |
|---|---|
| json.SyntaxError io.ErrUnexpectedEOF | There is a syntax problem with the JSON being decoded. |
| json.UnmarshalTypeError | A JSON value is not appropriate for the destination Go type. |
| json.InvalidUnmarshalError | The decode destination is not valid (usually because it is not a  pointer). This is actually a problem with our application code,  not the JSON itself. |
| io.EOF | The JSON being decoded is empty. |

