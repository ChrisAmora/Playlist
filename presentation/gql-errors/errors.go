package gql_errors

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GraphqlUnauthorized(ctx context.Context, message string) error {
	return &gqlerror.Error{Message: "Unauthorized", Path: graphql.GetPath(ctx), Extensions: map[string]interface{}{
		"httpStatusCode": "401",
		"customMessage":  message,
	}}
}

func GraphqlUserAlreadyExist(ctx context.Context) error {
	return &gqlerror.Error{Message: "User Already Exist", Path: graphql.GetPath(ctx), Extensions: map[string]interface{}{
		"httpStatusCode": "200",
		"customMessage":  "",
	}}
}

func GraphqlInvalidInput(ctx context.Context, message string) error {
	return &gqlerror.Error{Message: "Unprocessable Entity", Path: graphql.GetPath(ctx), Extensions: map[string]interface{}{
		"httpStatusCode": "422",
		"customMessage":  message,
	}}
}
