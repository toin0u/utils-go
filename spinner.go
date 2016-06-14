package utils

import (
	"fmt"
	"time"
)

// Spinner produce a little animation
func Spinner(done <-chan struct{}) {
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
