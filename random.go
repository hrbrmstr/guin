package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Object struct {
	Name  string
	Count int
}

func generateRandomObjects() []Object {
	objects := make([]Object, 10)
	for i := range objects {
		objects[i].Name = fmt.Sprintf("Name%d", i)
		objects[i].Count = rand.Intn(100)
	}
	return objects
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request")
		objects := generateRandomObjects()
		jsonData, err := json.Marshal(objects)
		if err != nil {
			http.Error(w, "Error generating JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization, Origin")
		w.Write(jsonData)
	})

	fmt.Println("Server running on port 9292")
	http.ListenAndServe(":9292", nil)
}
