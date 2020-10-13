package kubeless

import (
	"github.com/kubeless/kubeless/pkg/functions"
)

// Hello sample function with dependencies
func Hello(event functions.Event, context functions.Context) (string, error) {
	return "Hello World!", nil
}
