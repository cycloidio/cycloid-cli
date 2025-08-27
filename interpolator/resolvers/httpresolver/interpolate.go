package httpresolver

import (
	"github.com/cycloidio/cycloid-cli/interpolator/formatters"
	"github.com/cycloidio/cycloid-cli/interpolator/resources"
	"github.com/cycloidio/cycloid-cli/interpolator/transformers"
)

func (r HTTPResolver) Interpolate(uri string) (string, error) {
	ref, err := resources.NewResourceReference(uri)
	if err != nil {
		return "", err
	}

	data, err := r.Resolve(ref)
	if err != nil {
		return "", err
	}

	formatter := formatters.New(ref.Params)
	dataStr, err := formatter.Format(data)
	if err != nil {
		return "", err
	}

	return transformers.Transform(dataStr, ref.Params), nil
}
