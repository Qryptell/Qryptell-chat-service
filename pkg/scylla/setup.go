package scylla

import "log"

// Creating nessesorry tables for the database to work properly
func createTables() {
	var err = Session.Query("CREATE TABLE IF NOT EXISTS textmsg(id uuid primary key,msg_from text,msg_to text,type text,msg_time text,content text,message text);").Exec()
	if err != nil {
		log.Fatal(err)
	}
}
