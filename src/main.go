package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"tulongan-backend/src/schema"
)

func graphQLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error reading request body", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	fmt.Println("Successfully received request!")
	fmt.Println(body)

	var apolloQuery map[string]interface{}

	if err := json.Unmarshal(body, &apolloQuery); err != nil { // unmarshall body contents as a type query
		fmt.Println(err)

		fmt.Println("Error in unmarshalling JSON string.")

		w.WriteHeader(http.StatusUnprocessableEntity)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error in parsing JSON body", err)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	fmt.Println(apolloQuery)

	query := apolloQuery["query"]
	variables := apolloQuery["variables"]

	result := graphql.Do(
		graphql.Params{
			Schema:         schema.TulonganSchema,
			RequestString:  query.(string),
			VariableValues: variables.(map[string]interface{}),
		},
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&result)

	fmt.Println(result)

	return
}

func main() {
	port := ":3030"

	http.HandleFunc("/graphql", graphQLHandler)
	http.Handle("/playground/", http.StripPrefix("/playground/", http.FileServer(http.Dir("./graphql-playground"))))

	fmt.Println("Server running at port", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
