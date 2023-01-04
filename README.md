<h1>Golang CRUD API</h1>

> Status: finished. ✔

---

## About
CRUD API server made with Golang for an album management system, using [gorilla/mux](https://github.com/gorilla/mux) package.<br>
To download and run this project, you must have [GIT](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) and [Golang](https://go.dev/learn/) installed in your machine.<br>
Also, if you want to test the routes and methods, I recommend tools like [Insomnia](https://insomnia.rest/download) or [Postman](https://www.postman.com/downloads/).<br>
There's no database connected to the project, all the data is inside of an Array. Main focus here was to build a server in Golang with Rest archtecture.

---

## Download
<p>After install Git and Golang in you machine, follow the commands below to download the project.</p>

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

<p>The models were configured as structs, as you can see below</p>

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
> Here some orientation of how to set the right router and method for the CRUD action.

<table>
    <tr>
        <td>Action</td>
        <td>HTTP Method</td>
        <td>Router</td>
    </tr>
    <tr>
        <td>Get the list with all albums</td>
        <td><code>GET</code></td>
        <td><code>"/albums"</code></td>
    </tr>
    <tr>
        <td>Get the data of one Album</td>
        <td><code>GET</code></td>
        <td><code>"/albums/{id}"</code></td>
    </tr>
    <tr>
        <td>Create a new Album</td>
        <td><code>POST</code></td>
        <td><code>"/albums"</code></td>
    </tr>
    <tr>
        <td>Update data of one Album</td>
        <td><code>PUT</code></td>
        <td><code>"/albums/{id}"</code></td>
    </tr>
    <tr>
        <td>Delete a specific album</td>
        <td><code>DELETE</code></td>
        <td><code>"/albums/{id}"</code></td>
    </tr>
</table>

---
## Final

<p>Any question about this project or if you want to know more about me and my skills, contact me sending an email to 📩 cl.juniorr@gmail.com </p>
