package goutil

import (
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"
)

// WaitForServices tests and waits on the availability of a TCP host and port
func WaitForServices(services []url.URL, timeOut time.Duration) error {
	var depChan = make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(len(services))
	go func() {
		for _, s := range services {
			go func(s url.URL) {
				defer wg.Done()
				for {
					_, err := net.Dial(s.Scheme, s.Host)
					if err == nil {
						return
					}
					time.Sleep(5 * time.Second)
				}
			}(s)
		}
		wg.Wait()
		close(depChan)
	}()

	select {
	case <-depChan: // services are ready
		return nil
	case <-time.After(timeOut):
		return fmt.Errorf("services aren't ready in %s", timeOut)
	}
}
