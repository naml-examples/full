//
// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//   ███╗   ██╗ █████╗ ███╗   ███╗██╗
//   ████╗  ██║██╔══██╗████╗ ████║██║
//   ██╔██╗ ██║███████║██╔████╔██║██║
//   ██║╚██╗██║██╔══██║██║╚██╔╝██║██║
//   ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗
//   ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
//

package main

import (
	"os"

	"github.com/kris-nova/logger"
	"github.com/kris-nova/naml"
	app "github.com/naml-examples/simple"
)

// main is the main entry point for your CLI application
func main() {
	// Define your application.
	//
	// All of these are "Exported" and therefor can be a Kubernetes custom resource
	// or be thought of as a Values.yaml
	publicApp := &app.MySampleAppPublic{
		ExampleValue:       "",
		ExampleNumber:      0,
		ExampleText:        "",
		ExampleToggle:      false,
		ExampleVerbose:     0,
		ExampleName:        "",
		ExampleAnnotations: nil,
		ExampleValues:      nil,
		ExampleValue1:      "",
		ExampleValue2:      "",
		ExampleValue3:      "",
	}

	// Pass the public app fields to the New() function
	a := app.New("my-name", "my-namespace", "my-description", publicApp)

	// Register your application with naml
	naml.Register(a)

	// Run the default CLI tooling
	err := naml.RunCommandLine()
	if err != nil {
		logger.Critical("%v", err)
		os.Exit(1)
	}
}
