package main

import (
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

// Session for cassandra sessions
var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "test"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	logrus.Info("cassandra init done")
}

func main() {

}
