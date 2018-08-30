package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"tulongan-backend/src/controllers"
	"tulongan-backend/src/utils"
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

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error closing request body", err)
	}

	fmt.Println("Successfully received request!")

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

	query := apolloQuery["query"]
	variables := apolloQuery["variables"]

	result := graphql.Do(
		graphql.Params{
			Schema:         utils.TulonganSchema,
			RequestString:  query.(string),
			VariableValues: variables.(map[string]interface{}),
		},
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&result)

	return
}

func main() {
	user := controllers.GetUser()

	fmt.Println(user)

	http.HandleFunc("/graphql", graphQLHandler)

	log.Fatal(http.ListenAndServe(":3030", nil))
}
