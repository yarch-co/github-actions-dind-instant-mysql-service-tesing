package fut

import (
	"github.com/gofrs/flock"
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	log.Println("run at fut")

	l := flock.New("/tmp/foo.lock")

	var err error
	var locked bool
	for {
		locked, err = l.TryLock()
		if err != nil {
			panic(err)
		}
		if locked {
			break
		}
		log.Println("wait unlock at fut")
		time.Sleep(100 * time.Millisecond)
	}

	log.Println("locked at fut")
	code := m.Run()
	log.Println("unlock at fut")
	l.Unlock()
	log.Println("done at fut")

	os.Exit(code)
}
