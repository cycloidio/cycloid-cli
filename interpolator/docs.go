package interpolator

import (
	"strings"

	"github.com/cycloidio/cycloid-cli/interpolator/formatters"
	"github.com/cycloidio/cycloid-cli/interpolator/transformers"
)

const URIDoc = `An URI resource path is defined as follows:

  cy://<api_path>[?<param>=<value>&<param2...>]

Where:
  <api_path>
    Any valid path from our API as defined in our api docs: https://docs.cycloid.io/api/

  <params> 
    Urlencoded parameters, they can be either an API parameter, or a formatting parameter.
    formatting parameters allow you to format the output of the API.

    By default the output will be JSON for objects and array and a string for any simple value.

	<option>
		Each params can have options, some with value assigned with =

Params:
  key
    Query the output value using gojq syntax ( https://github.com/itchyny/gojq )
    that match jq syntax, see the manual here https://jqlang.org/manual

    The basic syntax is starting with a . that represent the current object.
    You can access keys by their names .<key> and use array syntax for lists .[]

    If no key is provided, the output will be the basic API Response.
    You can check API responses on our API documentation https://docs.cycloid.io/api/
`

var Docs = strings.Join([]string{URIDoc, formatters.Docs, transformers.Docs}, "\n")
