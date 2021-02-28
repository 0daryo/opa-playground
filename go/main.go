package main

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

var (
	//go:embed example.rego
	module string
)

func main() {
	ctx := context.Background()
	query, err := rego.New(
		rego.Query("x = data.example.authz.allow"),
		rego.Module("example.rego", module),
	).PrepareForEval(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	input := map[string]interface{}{
		"role":   "customer",
		"action": "read",
		"object": "admin_api",
	}

	checkCtx := context.TODO()
	results, err := query.Eval(checkCtx, rego.EvalInput(input))
	if err != nil {
		fmt.Println(err)
	} else if len(results) == 0 {
		fmt.Println("zero result")
	} else if result, ok := results[0].Bindings["x"].(bool); ok {
		fmt.Println("bool result")
		fmt.Println(result)
	} else {
		fmt.Println("other")
		fmt.Println(results)
	}
}
