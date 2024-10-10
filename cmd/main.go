package main

import (
	"fmt"
	"log"
	"mole/internal/mole"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	moles := []*mole.Mole{}
	for i := 0; i < 6; i++ {
		moles = append(moles, mole.NewMole(fmt.Sprintf("localhost:808%d", i+1)))
	}
	go moles[0].Listen()
	wg := sync.WaitGroup{}
	for i, lm := range moles[1:] {
		wg.Add(1)
		littleMole := lm
		filename := filepath.Join(dir, "from", fmt.Sprintf("test_%d.txt", i+1))
		go func() {
			if err := littleMole.SendFile("localhost:8081", filename); err != nil {
				log.Println("littleMole: ", err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	moles[0].Wait()
	moles[0].Stop()
}
