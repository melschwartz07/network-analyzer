package main

import (
	"fmt"
	"log"
	"os"
	"time" 

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	// Neo4j connection
	uri := os.Getenv("NEO4J_URI")
	user := os.Getenv("NEO4J_USER")
	password := os.Getenv("NEO4J_PASSWORD")

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(user, password, ""))
	if err != nil {
		log.Fatal("Failed to create Neo4j driver:", err)
	}
	defer driver.Close()

	// Verify connectivity
	err = driver.VerifyConnectivity()
	if err != nil {
		log.Fatal("Failed to connect to Neo4j:", err)
	}
	fmt.Println("Connected to Neo4j!")

	// test node
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run("CREATE (n:Node {name: $name}) RETURN n", 
			map[string]interface{}{"name": "Test Node"})
		return nil, err
	}, neo4j.WithTxTimeout(5*time.Second)) 
	
	if err != nil {
		log.Fatal("Failed to create node:", err)
	}
	fmt.Println("Created test node in Neo4j!")
}