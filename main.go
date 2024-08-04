package main

import (
	"toolbar"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("toolbar", func(w *unison.Window) {
		w.Content().AddChild(toolbar.New().Layout())
	})
}
