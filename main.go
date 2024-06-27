package main

import (
	"toolbar"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("toolbar", func(w *unison.Window) {
		w.Content().AddChild(toolbar.New().Layout())
	})
}
