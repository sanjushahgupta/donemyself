// package Handle

// import (
// 	"encoding/json"
// 	"firstattemp/Struct"
// 	"log"
// 	"net/http"
// )

// // This is the handler that creates a new post,
// //  we start by creating a new instance of the struct contact.

// // 	 Next we create a random ID with the “math/rand” package
// // 	as well as convert to a string. At last we simply append the post to our posts array, this will save it to memory,
// // 	and right after we return the new post.

// func (a *App) Create(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	request := Struct.Contact{}
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	defer r.Body.Close()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	var data Struct.Contact
// 	stmt := "INSERT INTO newtable(Name,StreetNo,Locality) VALUES($1,$2,$3)"
// 	_, err = a.DBW.Exec(stmt, data.Name, data.StreetNo, data.Locality)
// 	if err != nil {
// 		log.Println(err)

// 	}

// 	//  we’re just setting the header “Content-Type” to “application/json” .w.Header().Set("Content-Type", "application/json")

// }

// func Read(w http.ResponseWriter, r *http.Request) {

// }

// // func Delete(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")
// // 	params := mux.Vars(r)
// // 	for idx, details := range Persons {
// // 		if details.Name == params["name"] {
// // 			Persons = append(Persons[:idx], Persons[idx+1:]...)
// // 			break
// // 		}
// // 	}
// // 	json.NewEncoder(w).Encode(Persons)
// // }

// // func Update(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")
// // 	params := mux.Vars(r)
// // 	for idx, details := range Persons {
// // 		if details.Name == params["name"] {
// // 			Persons = append(Persons[:idx], Persons[idx+1:]...)
// // 			var uperson Struct.Contact
// // 			_ = json.NewDecoder(r.Body).Decode(&uperson)
// // 			uperson.Name = params["name"]
// // 			Persons = append(Persons, uperson)
// // 			json.NewEncoder(w).Encode(Persons)
// // 			return

// // 		}
// // 	}
// // }
