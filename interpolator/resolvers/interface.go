package resolvers

import "github.com/cycloidio/cycloid-cli/interpolator/resources"

type ResourceResolver interface {
	Resolve(ref *resources.Reference) ([]any, error)
	Interpolate(input string) (string, error)
}
