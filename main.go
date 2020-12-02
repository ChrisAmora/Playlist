package main

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func createEmp(emp Emp) {
	fmt.Println(" **** Creating new emp ****\n", emp)
	if err := Session.Query("INSERT INTO emps(empid, first_name, last_name, age) VALUES(?, ?, ?, ?)",
		emp.id, emp.firstName, emp.lastName, emp.age).Exec(); err != nil {
		fmt.Println("Error while inserting Emp")
		fmt.Println(err)
	}
}

func getEmps() []Emp {
	fmt.Println("Getting all Employees")
	var emps []Emp
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM emps").Iter()
	for iter.MapScan(m) {
		emps = append(emps, Emp{
			id:        m["empid"].(string),
			firstName: m["first_name"].(string),
			lastName:  m["last_name"].(string),
			age:       m["age"].(int),
		})
		m = map[string]interface{}{}
	}

	return emps
}

func updateEmp(emp Emp) {
	fmt.Printf("Updating Emp with id = %s\n", emp.id)
	if err := Session.Query("UPDATE emps SET first_name = ?, last_name = ?, age = ? WHERE empid = ?",
		emp.firstName, emp.lastName, emp.age, emp.id).Exec(); err != nil {
		fmt.Println("Error while updating Emp")
		fmt.Println(err)
	}
}

func deleteEmp(id string) {
	fmt.Printf("Deleting Emp with id = %s\n", id)
	if err := Session.Query("DELETE FROM emps WHERE empid = ?", id).Exec(); err != nil {
		fmt.Println("Error while deleting Emp")
		fmt.Println(err)
	}
}

var Session *gocql.Session

type Emp struct {
	id        string
	firstName string
	lastName  string
	age       int
}

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
	cluster.Keyspace = "Betinho"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("betinho")
}

func main() {
	emp1 := Emp{"E-1", "Betinho", "Nasus", 20}
	emp2 := Emp{"E-2", "Bati", "Nabewata", 30}
	createEmp(emp1)
	fmt.Println(getEmps())
	createEmp(emp2)
	fmt.Println(getEmps())
	emp3 := Emp{"E-1", "Rahul", "Anand", 30}
	updateEmp(emp3)
	fmt.Println(getEmps())
	deleteEmp("E-2")
	fmt.Println(getEmps())
	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	http.ListenAndServe(":8000", router)
}
