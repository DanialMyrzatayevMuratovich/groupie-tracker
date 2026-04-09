package updater

import (
	"groupie-tracker/internal/controller"
	"log"
	"time"
)

func Start(interval time.Duration) {
	go func() {
		for {
			if err := controller.WarmCache(); err != nil {
				log.Println("cache refresh error:", err)
			}
			time.Sleep(interval)
		}
	}()
}
