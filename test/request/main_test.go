package request

import (
  "database/sql"
  "fmt"
	"log"
	"os"
	"testing"

  _ "github.com/go-sql-driver/mysql"
  "github.com/ory/dockertest/v3"
  "github.com/ory/dockertest/v3/docker"
)

func TestMain(m *testing.M) {
	log.Println("run at request")

  // 明示しないとよしなに設定してくれる
  pool, err := dockertest.NewPool("")
  if err != nil {
    panic(err)
  }
  resource, err := pool.RunWithOptions(&dockertest.RunOptions{
    Repository: "mysql",
    Tag: "8.0",
    Env: []string{"MYSQL_ROOT_PASSWORD=secret"},
  }, func(config *docker.HostConfig) {
    config.AutoRemove = true
    config.RestartPolicy = docker.RestartPolicy{
      Name: "no",
    }
  })
  if err != nil {
    panic(err)
  }

  dbHost := os.Getenv("DB_HOST")
  if dbHost == "" {
    dbHost = "localhost"
  }
  var db *sql.DB
  if err := pool.Retry(func() error {
    var err error
    log.Printf("trying to connect mysql at request: %s %s(port: %v)", dbHost, resource.Container.Name, resource.GetPort("3306/tcp"))
    db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(%s:%s)/mysql", dbHost, resource.GetPort("3306/tcp")))
    if err != nil {
      return err
    }
    return db.Ping()
  }); err != nil {
    panic(err)
  }

	log.Printf("initiated mysql at request: %s %s(port: %v)", dbHost, resource.Container.Name, resource.GetPort("3306/tcp"))
	code := m.Run()
	log.Println("done at request")

  if err := pool.Purge(resource); err != nil {
    panic(err)
  }

	os.Exit(code)
}
