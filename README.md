<h1>Golang CRUD API</h1>

> Status: finished. ✔

---

## About
<p>CRUD API server made with Golang for an album management system, using gorilla/mux package.</p>

<p>To download and run this project, you must have GIT and Golang installed in your machine. Also, if you want to test the routes and methods, I recommend tools like Insomnia or Postman.</p>

<p>There's no database connected to the project, all the data is inside of an Array. Main focus here was to build a server in Golang with Rest archtecture.</p>

---

## Download
<p>After installed Git and Golang in you machine, follow the commands below to downaload the project.</p>

```bash
# step 1: clone the repository
git clone git@github.com:celiovjunior/go-crud-simple.git

# step 2: enter the folder
cd go-crud-simple

# step 3: install the necessary packages
go mod tidy

# step 4: execute the server
go run main.go
```

<p>You will see a message from the terminal saying</p>

```code
Server is running at port 8080
```
<p>Use <code>http://localhost:8080/</code> as the base URL to make the tests.</p>

---

## Exploring

<p>The models are configured as structs, as you can see below</p>

```go
// Model for ALBUM
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Artist *Artist `json: "artist"`
}

// Model for ARTIST
type Artist struct {
	Name string `json:"name"`
}
```

<p>The application use JSON format to work with data, so there's a example on how to write your tests.</p>

```json
{
    "title": "di melo",
    "year": 1975,
    "artist": "di melo"
}
```

