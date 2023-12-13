package scylla

import (
  "chat_v2/logger"
  "github.com/gocql/gocql"
	"go.uber.org/zap"

)

var(
	my_logger = logger.CreateLogger("info")
	session   *gocql.Session
)

func Init(cluster gocql.ClusterConfig) error{
  var err error
	session, err = gocql.NewSession(cluster)
	// if err != nil {
	// 	my_logger.Fatal("unable to connect to scylla", zap.Error(err))
	// }
 return err
}

func Session() *gocql.Session{
  return session
}

func Logger() *zap.Logger{
  return my_logger
}
