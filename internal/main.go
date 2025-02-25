package internal

import (
	"context"
	"time"
)

var YamlLocation string

func Run(ctx context.Context) {
	t := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			gather()
		}
	}
}

func gather() {

}
