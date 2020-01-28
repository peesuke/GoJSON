package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Member struct{
  Id  int
  FirstName string
  LastName string
  Color string
}

func main(){
  HandleRequest()
}

func HandleRequest(){
  http.HandleFunc("/", RouteHomePage)
  http.HandleFunc("/Members", RouteGetMembers)
  http.ListenAndServe(":9090", nil)
}

func RouteHomePage(w http.ResponseWriter, r *http.Request){
  fmt.Println("HomePage has been accessed")
}

func RouteGetMembers(w http.ResponseWriter, r *http.Request){
  member := []Member{}
	n := Member{}
	n = Member{1, "Bae", "Joohyun", "Pink"}
	member = append(member, n)

	n = Member{2, "Kang", "Seulgi", "Yellow"}
	member = append(member, n)

	n = Member{3, "Son", "Seung-wan", "Blue"}
	member = append(member, n)

	n = Member{4, "Park", "Soo-young", "Green"}
	member = append(member, n)

	n = Member{5, "Kim", "Yerim", "Purple"}
	member = append(member, n)

	json.NewEncoder(w).Encode(member)
}
