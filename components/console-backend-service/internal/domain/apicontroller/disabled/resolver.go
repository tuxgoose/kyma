// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import context "context"
import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

// Resolver is an autogenerated failing mock type for the Resolver type
type Resolver struct {
	err error
}

// NewResolver creates a new Resolver type instance
func NewResolver(err error) *Resolver {
	return &Resolver{err: err}
}

// APIsQuery provides a failing mock function with given fields: ctx, namespace, serviceName, hostname
func (_m *Resolver) APIsQuery(ctx context.Context, namespace string, serviceName *string, hostname *string) ([]gqlschema.API, error) {
	var r0 []gqlschema.API
	var r1 error
	r1 = _m.err

	return r0, r1
}
