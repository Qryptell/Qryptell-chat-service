package scylla

import (
	"os"
	"time"

	"github.com/Qryptell/Qryptell-chat-service/util"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

// Setting up scylladb
func SetUp() {
	// Getting Redis ports
	var keyspace = os.Getenv("SCYLLA_KEYSPACE")
	var hosts = util.StringToArray(os.Getenv("SCYLLA_HOSTS"))

	var cluster = createCluster(keyspace, hosts...)
	var session, err = gocql.NewSession(*cluster)
	if err != nil {
		panic(err)
	}
	Session = session

	createTables()
}

// Creating scylladb cluster
func createCluster(kepsace string, hosts ...string) *gocql.ClusterConfig {
	var retryPolicy = &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        10 * time.Second,
		NumRetries: 5,
	}

	var cluster = gocql.NewCluster(hosts...)
	cluster.Keyspace = kepsace
	cluster.RetryPolicy = retryPolicy
	cluster.Consistency = gocql.Quorum
	var tokenPolicy = gocql.RoundRobinHostPolicy()

	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(tokenPolicy)
	return cluster
}
