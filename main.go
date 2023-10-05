package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//___________ 1. Models______________

type MyFriends struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Nickname string  `json:"nick-name"`
	Home     string  `json:"home"`
	Phone    string  `json:"phone"`
	Skills   *Skills `json:"skills"`
}

type Skills struct {
	Domain    string   `json:"domain"`
	Languages []string `json:"languages"`
	Projects  []string `json:"projects"`
}


//_____________2. Fake Db________________

 var Friends = []MyFriends{
        {
            Id:       "1",
            Name:     "John Doe",
            Nickname: "Johnny",
            Home:     "New York",
            Phone:    "123-456-7890",
            Skills: &Skills{
                Domain:    "Software Development",
                Languages: []string{"Go", "Python"},
                Projects:  []string{"Project A", "Project B"},
            },
        },
        {
            Id:       "2",
            Name:     "Jane Smith",
            Nickname: "Janey",
            Home:     "Los Angeles",
            Phone:    "987-654-3210",
            Skills: &Skills{
                Domain:    "Web Development",
                Languages: []string{"JavaScript", "HTML/CSS"},
                Projects:  []string{"Website X", "Website Y"},
            },
        },
        // Add more dummy data as needed
    }


//_____________3.Middlewares , Controllers , Helpers________________

func IsEmpty(c *MyFriends) bool {
	return c.Id == "" && c.Name == ""
}

func getAllFriends( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(Friends)
}

func getFriend(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    friendID := params["id"]

    for _, friend := range Friends {
        if friend.Id == friendID {
            json.NewEncoder(w).Encode(friend)
            return
        }
    }

    http.NotFound(w, r)
}


func updateFriend(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    friendID := params["id"]

    for index, friend := range Friends {
        if friend.Id == friendID {
            var updatedFriend MyFriends
            err := json.NewDecoder(r.Body).Decode(&updatedFriend)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

    
            Friends[index] = updatedFriend
            json.NewEncoder(w).Encode(updatedFriend)
            return
        }
    }

    http.NotFound(w, r)
}


func deleteFriend(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    friendID := params["id"]

    for index, friend := range Friends {
        if friend.Id == friendID {
            Friends = append(Friends[:index], Friends[index+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.NotFound(w, r)
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("working......")
	w.Write([]byte("<h1>Server Starting!</h1>"))
}


//_____________MAIN FUNCTION________________

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET");
	r.HandleFunc("/friends", getAllFriends).Methods("GET")
	r.HandleFunc("/friends/{id}", getFriend).Methods("GET")
    r.HandleFunc("/friends/{id}", updateFriend).Methods("PUT")    
    r.HandleFunc("/friends/{id}", deleteFriend).Methods("DELETE")
	fmt.Println("Server is listening on port 8080...")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
