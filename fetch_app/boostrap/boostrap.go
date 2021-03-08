package boostrap

import (
	"fmt"

	"github.com/fetch_app/pkg/logger"
)

// RunApp :nodoc:
func RunApp() {
	initService()

	router := initREST()
	if err := router.Run(); err != nil {
		logger.Panic(
			fmt.Sprintf("Receiving error: %v", err),
		)
	}
}
