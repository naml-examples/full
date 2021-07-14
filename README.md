# Full Application

A broader example of how to build a [naml](https://github.com/kris-nova/naml) project.

This is yet another example of the flexibility with naml.

In this example we have a large set of configurable values (similar to a `Values.yaml` file) that we expose in the codebase.

```go
// MySampleAppPublic can be used for any public (or exported) facing mechanisms.
// - Kubernetes Custom Resources
// - Alternative to Values.yaml
// - Exposed over an HTTP API
// - Exposed over a gRPC API
//
// Here is where you could define a large amount of values that another mechanism could "tweak" or "configure"
// just like a Values.yaml.
//
// By making this (and the sub fields) exported we could expose this to other Go packages or even to a Kubernetes
// custom resource.
type MySampleAppPublic struct {

	// In case anyone is wondering this is "the new Values.yaml" as long as you plumb the fields
	// through in the "implementation" below.

	ExampleValue string // See line 84, and line 110
	ExampleNumber int
	ExampleText string
	ExampleToggle bool
	ExampleVerbose int
	ExampleName string
	ExampleAnnotations map[string]string
	ExampleValues map[int]string
	ExampleValue1 string
	ExampleValue2 string
	ExampleValue3 string
}
```
