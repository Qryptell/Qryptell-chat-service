module github.com/Qryptell/Qryptell-chat-service

go 1.21.5

require (
	github.com/gocql/gocql v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/redis/go-redis/v9 v9.4.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.3.1
