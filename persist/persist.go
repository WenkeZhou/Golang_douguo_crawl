package persist

import "log"

func ItemSave() chan interface{} {
	out := make(chan interface{})
	itemcount := 0
	go func() {
		for {
			item := <-out
			log.Printf("Item saver:Got%d, %v", itemcount, item)
			itemcount++
		}
	}()

	return out
}
