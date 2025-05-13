func monitor(driver neo4j.Driver) {
	for {
		session := driver.NewSession(neo4j.SessionConfig{})
		result, err := session.Run(`
			MATCH (n)-[r]->(m)
			WHERE r.timestamp > datetime().subtract(duration('PT5M'))
			WITH n.name AS source, count(r) AS connection_count
			WHERE connection_count > 50
			RETURN source, connection_count
			ORDER BY connection_count DESC
		`, nil)
		
		if err == nil {
			for result.Next() {
				record := result.Record()
				log.Printf("ALERT: %s made %d connections.",
					record.Values[0].(string),
					record.Values[1].(int64))
			}
		}
		session.Close()
		time.Sleep(30 * time.Second)
	}
}