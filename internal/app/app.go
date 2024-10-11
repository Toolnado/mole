package app

import (
	"context"
	"flag"
	"mole/internal/mole"
	"mole/pkg/logger"
	"os"
	"path/filepath"
)

type App struct {
	ctx context.Context
}

func New(ctx context.Context) *App {
	return &App{ctx: ctx}
}

func (app *App) Run() {
	address := flag.String("address", "127.0.0.1:3000", "the address that the mole will listen to")
	sender := flag.Bool("send", false, "indicates whether the file should be sent")
	file := flag.String("file", "", "the file to be sent")
	flag.Parse()

	l := logger.New()
	m := mole.New(context.TODO(), *address, l)
	if !*sender {
		go func() {
			if err := m.Listen(); err != nil {
				l.Fatal("%s", err)
			}
		}()
		m.Wait()
		m.Stop()
	} else {
		if *file == "" {
			l.Fatal("%s", "empty file")
		}

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			l.Fatal("%s", err)
		}

		*file = filepath.Join(dir, *file)
		if err := m.SendFile(*address, *file); err != nil {
			l.Error("%s", err)
		}
	}
}
