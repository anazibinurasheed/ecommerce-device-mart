package helper

import (
	"sync"
	"time"
)

func GoClean(m map[string]string, uuid string, mutex *sync.Mutex) {

	time.Sleep(1 * time.Minute)
	mutex.Lock()
	delete(m, uuid)
	mutex.Unlock()

}
