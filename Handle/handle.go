package Handle

import (
	"encoding/json"
	"firstattemp/Struct"
	"net/http"

	"github.com/gorilla/mux"
)

var Person Struct.Contact
var Persons []Struct.Contact

// This is the handler that creates a new post,
//  we start by creating a new instance of the struct contact.

// 	 Next we create a random ID with the “math/rand” package
// 	as well as convert to a string. At last we simply append the post to our posts array, this will save it to memory,
// 	and right after we return the new post.

func Create(w http.ResponseWriter, r *http.Request) {
	//  we’re just setting the header “Content-Type” to “application/json” .
	w.Header().Set("Content-Type", "application/json")
	// 	Then we decode the data that is sent with the request and insert it into the Person.
	_ = json.NewDecoder(r.Body).Decode(&Person)
	Persons = append(Persons, Person)
	json.NewEncoder(w).Encode(Persons)
}

func Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Persons)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, details := range Persons {
		if details.Name == params["name"] {
			Persons = append(Persons[:idx], Persons[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Persons)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, details := range Persons {
		if details.Name == params["name"] {
			Persons = append(Persons[:idx], Persons[idx+1:]...)
			var uperson Struct.Contact
			_ = json.NewDecoder(r.Body).Decode(&uperson)
			uperson.Name = params["name"]
			Persons = append(Persons, uperson)
			json.NewEncoder(w).Encode(Persons)
			return

		}
	}
}
