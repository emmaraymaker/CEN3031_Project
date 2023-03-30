package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	backend "test/Structs"
)

var memL backend.MemberList

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/user/reg", registerUser).Methods("Post")
	r.HandleFunc("/user/get", getUser).Methods("GET")
	r.HandleFunc("/user/update", updateUser).Methods("POST")
	r.HandleFunc("/user/search", searchUser).Methods("GET")
	r.HandleFunc("/listing/set", setListing).Methods("POST")
	r.HandleFunc("/listing/get", getListing).Methods("GET")
	r.HandleFunc("/listing/update", updateListing).Methods("POST")
	r.HandleFunc("/listing/search", searchListing).Methods("GET")

	http.ListenAndServe(":8000", r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + os.Getenv("PORT"),
	}

	log.Fatal(srv.ListenAndServe())

	// Test set for pre-front-end integration

	/*
		var memL backen.MemberList

			memL.AddMember("1", "Test, t", "888-888-8888", "test@test.com", "password", false, false)
			memL.AddMember("2", "Test2, t", "888-888-8888", "test@test.com", "password", false, false)
			memL.AddMember("3", "Test3, t", "888-888-8888", "test@test.com", "password", false, false)

			var name, _ = memL.SearchMember("2")
			fmt.Println(name)

			memL.RemoveMember("2")

			for i := 0; i < 2; i++ {
				fmt.Println(memL.Ml[i].Name)
			}
			var memL MemberList
			var lisL ListingSet

			var uid string
			var name string
			var phone string
			var email string
			fmt.Println("Hello, please enter your Union Id:")
			fmt.Scanln(&uid)
			fmt.Println("Please enter your name:")
			fmt.Scanln(&name)
			fmt.Println("Please enter your Phone number:")
			fmt.Scanln(&phone)
			fmt.Println("Please enter your Email address:")
			fmt.Scanln(&email)
			memL.addMember(uid, name, phone, email, isL, isT)

			for {
				fmt.Println("Please select a menu option:\n1: Rent\n2:List")

				var option string
				fmt.Scanln(&option)

				var location string
				var price int
				if option == "1" {
					fmt.Println("displaying listings:")
					fmt.Println(lisL.ls[0].location)
				} else if option == "2" {

					fmt.Println("Please enter the Listing's Location:")
					fmt.Scanln(&location)
					fmt.Println("Please enter the Listing's Price:")
					fmt.Scanln(&price)
					lisL.addListing(location, price)
				} else {
					break
				}
			}
	*/
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var usr backend.Member
	_ = json.NewDecoder(r.Body).Decode(&usr)
	memL.Ml = append(memL.Ml, usr)
	//json.NewEncoder(w).Encode(memL)
}
func getUser(w http.ResponseWriter, r *http.Request) {

}
func updateUser(w http.ResponseWriter, r *http.Request) {

}
func searchUser(w http.ResponseWriter, r *http.Request) {

}

func setListing(w http.ResponseWriter, r *http.Request) {

}
func getListing(w http.ResponseWriter, r *http.Request) {

}
func updateListing(w http.ResponseWriter, r *http.Request) {

}
func searchListing(w http.ResponseWriter, r *http.Request) {

}
