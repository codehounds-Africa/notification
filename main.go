package main

import (
	"os"

	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

// Session for cassandra sessions
var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST"))
	cluster.Keyspace = os.Getenv("KEY_SPACE")
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	logrus.Info("cassandra init done")
}

func main() {

}
