package bitcoinmonitor

import (
	"bitcoinmonitor/internal/application"
)

func main() {
	app := application.NewApplication()
	app.Run()
}
