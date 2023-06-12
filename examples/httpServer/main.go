// @title        Example HTTP API
// @version      1.0
// @description  An example HTTP API to demonstrate the usage of the Arsenal framework.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

package main

import (
	"examples/httpServer/cmd"
	_ "examples/httpServer/docs"
)

func main() {
	cmd.Execute()
}
