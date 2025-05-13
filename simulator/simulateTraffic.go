package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// generate simulated network traffic data and store in Neo4j
func simulateTraffic(driver neo4j.Driver) {
	// generating random network events
	devices := []string{"server1", "server2", "workstation1", "router1"} 
	ports := []int32{80, 443, 22, 3389} // Common network ports (HTTP, HTTPS, SSH, RDP)
	actions := []string{"ALLOW", "BLOCK"} // Possible firewall actions

	// Infinite loop to continuously generate traffic
	for {
		// select source, target, port and action at random
		source := devices[rand.Intn(len(devices))] // source 
		target := devices[rand.Intn(len(devices))] // target
		port := ports[rand.Intn(len(ports))]       //port
		action := actions[rand.Intn(len(actions))] // Allow or block

		// new Neo4j session
		session := driver.NewSession(neo4j.SessionConfig{})
		
		// Execute a write transaction to store the connection data
		_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			// Cypher query to:
			// MERGE - ensure that devices exist
			// CREATE new connection 
			_, err := tx.Run(`
				MERGE (s:Device {name: $source}) 
				MERGE (t:Device {name: $target}) 
				CREATE (s)-[r:CONNECTED {        
					port: $port,                
					action: $action,            
					timestamp: datetime()       
				}]->(t)
				RETURN r                       
			`, map[string]interface{}{
				"source": source,
				"target": target,
				"port":   port,
				"action": action,
			})
			return nil, err
		})
		
		session.Close()

		if err != nil {
			log.Printf("Error writing to Neo4j: %v", err)
		} else {
			log.Printf("Simulated %s → %s:%d (%s)", source, target, port, action)
		}

		// implement a delay
		time.Sleep(time.Second * time.Duration(rand.Intn(3)+1))
	}
}