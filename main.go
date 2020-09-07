package main

import (
	"context"
	"fmt"

	"github.com/apache/airflow-client-go/airflow"
)

func main() {
	cli := airflow.NewAPIClient(&airflow.Configuration{
		Scheme:   "http",
		Host:     "localhost:28080",
		BasePath: "/api/v1",
	})

	cred := airflow.BasicAuth{
		UserName: "username",
		Password: "password",
	}

	ctx := context.WithValue(context.Background(), airflow.ContextBasicAuth, cred)

	variable, _, err := cli.VariableApi.GetVariable(ctx, "foo")
	if err != nil {
		fmt.Println(variable)
	} else {
		fmt.Println(err)
	}
}
