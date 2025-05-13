// monitor/monitor.go
package main

import (
	"context"
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func monitorNetwork(driver neo4j.Driver) {
	for {
		checkConnectionFloods(driver)
		checkBlockedAttempts(driver)
		checkPortScans(driver)
		time.Sleep(30 * time.Second)
	}
}

func checkConnectionFloods(driver neo4j.Driver) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, _ := session.Run(`
		MATCH (src)-[r:CONNECTED]->()
		WHERE r.timestamp > datetime().subtract(duration('PT1M'))
		WITH src, count(r) as connections
		WHERE connections > 50
		RETURN src.name as source, connections
		ORDER BY connections DESC
	`, nil)

	for result.Next() {
		rec := result.Record()
		log.Printf("FLOOD ALERT: %s made %d connections/minute", 
			rec.Values[0].(string), 
			rec.Values[1].(int64))
	}
}