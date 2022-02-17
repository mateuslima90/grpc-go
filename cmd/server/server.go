package main

import (
	"github.com/mateuslima90/grpc-go/pb"
	"github.com/mateuslima90/grpc-go/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

//func startHttpServer(r *mux.Router) {
//	log.Fatal(http.ListenAndServe(":8081", r))
//}

type Person struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

//func getAllUser(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("Content-Type", "application/json")
//
//	var persons []Person
//
//	persons = append(persons, Person{Username: "mateus_lima", Name: "Mateus Lima Fonseca", Email: "mateuslima90@gmail.com"})
//	json.NewEncoder(writer).Encode(persons)
//}

//func getUserByUsername(writer http.ResponseWriter, request *http.Request) {
//
//	params := mux.Vars(request)
//	fmt.Println(params)
//	writer.Header().Set("Content-Type", "application/json")
//	//json.NewEncoder(writer).Encode(p)
//	//id := strings.TrimPrefix(request.URL.Path, "/customer/")
//	//fmt.Println(id)
//	//fmt.Println(request.URL.Path)
//	//
//	// err := json.NewDecoder(request.Body).Decode(&p)
//	// if err != nil {
//	// 	http.Error(writer, err.Error(), http.StatusBadRequest)
//	// 	return
//	// }
//	// //res, _ := json.Marshal(p)
//	user := repository.GetUser(params["username"])
//
//	res2, _ := json.Marshal(user)
//	writer.WriteHeader(http.StatusOK)
//	writer.Write(res2)
//}
//
//func authenticate(writer http.ResponseWriter, request *http.Request) {
//
//	//Pega os campos do formulário
//	username := request.FormValue("username")
//	password := request.FormValue("password")
//
//	resp := struct {
//		Username string
//		Password string
//	}{
//		username,
//		password,
//	}
//
//	res, _ := json.Marshal(resp)
//	writer.WriteHeader(http.StatusOK)
//	writer.Write(res)
//
//}

//func saveUser(writer http.ResponseWriter, request *http.Request) {
//
//	var p Person
//
//	// Try to decode the request body into the struct. If there is an error,
//	// respond to the client with the error message and a 400 status code.
//	err := json.NewDecoder(request.Body).Decode(&p)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	repository.InsertUser2(p.Username, p.Name, p.Email)
//
//	res, _ := json.Marshal(p)
//	writer.WriteHeader(http.StatusOK)
//	writer.Write(res)
//}

func main() {

	//r := mux.NewRouter()
	//
	//r.HandleFunc("/users.go", getAllUser).Methods("GET")
	//
	//r.HandleFunc("/user/{username}", getUserByUsername).Methods("GET")
	//
	//r.HandleFunc("/user/authenticate", authenticate).Methods("POST")
	//
	//r.HandleFunc("/user", saveUser).Methods("POST")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	//go startHttpServer(r)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
