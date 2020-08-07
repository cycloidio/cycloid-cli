package externalBackends

import (
	"github.com/spf13/cobra"
)

var regionFlag string
var urls []string
var prefilters map[string]string

func WithFlagAwsRegion(cmd *cobra.Command) string {
	flagName := "region"
	cmd.Flags().StringVar(&regionFlag, flagName, "eu-west-1", "region")
	return flagName
	// cmd.Flags().SetAnnotation("ppurpose", "create", []string{"aws", "log"})
	// Uncomment in v3
	// cmd.Flags().MarkDeprecated("pproject", "Deprecated flag pproject")
	//
	// cmd.Flags().StringVar(&ebPurpose, "purpose", "default-p", "purpose")
}

func WithFlagUrl(cmd *cobra.Command) string {
	flagName := "url"
	cmd.Flags().StringSliceVar(&urls, flagName, nil, "urls")
	return flagName
}
func WithFlagPrefilter(cmd *cobra.Command) string {
	flagName := "prefilter"
	cmd.Flags().StringToStringVar(&prefilters, flagName, nil, "key=value")
	return flagName
}

//   title: New External backend
//   description: >-
//     An external backend contains the configuration needed in order to be
//     plugged into the Cycloid system.
//     A backend is a general purpose concept, but Cycloid specifies which ones
//     are supported and the list of those which are supported for every
//     concrete feature.
//   required:
//     - purpose
//     - configuration
//   properties:
//     purpose:
//       type: string
//       enum: ['events', 'logs', 'remote_tfstate']
//     credential_id:
//       type: integer
//       format: uint32
//       minimum: 1
//     project_canonical:
//       type: string
//       minLength: 3
//       maxLength: 30
//       pattern: '^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$'
//     environment_canonical:
//       type: string
//       minLength: 1
//       maxLength: 30
//       pattern: '^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$'
//     configuration:
//       $ref: '#/definitions/ExternalBackendConfiguration'

// ExternalBackendConfiguration:
//   type: object
//   discriminator: engine
//   properties:
//     engine:
//       type: string
//   required:
//     - engine

// ElasticsearchLogs:
//   description: |
//     Representation of Elasticsearch logs for external backend.
//   allOf:
//     - $ref: '#/definitions/ExternalBackendConfiguration'
//     - type: object
//       required:
//         - urls
//         - version
//         - sources
//       properties:
//         urls:
//           type: array
//           description: |
//             List of the URLs
//           items:
//             type: string
//         version:
//           type: string
//           description: |
//             Only 7 is supported
//         sources:
//           type: object
//           description: |
//             It's an object where the key is the 'environment' and the value
//             another object where the key is the 'source-name' and value a Source.
//           additionalProperties:
//             type: object
//             additionalProperties:
//               type: object
//               properties:
//                 index:
//                   type: string
//                   description: |
//                     Index to use
//                 urls:
//                   type: array
//                   description: |
//                     List of URLs to override the main URL defined
//                   items:
//                     type: string
//                 prefilters:
//                   type: object
//                   description: |
//                     JSON representing the prefilters to apply to the index to get
//                     the specific values.
//                 mapping:
//                   type: object
//                   description: |
//                     Object with the mapping to know which attributes are the ones
//                     we have to map to the ones we want
//                   required:
//                     - host
//                     - timestamp
//                     - message
//                   properties:
//                     host:
//                       type: string
//                       description: |
//                         The Host of the log
//                     timestamp:
//                       type: string
//                       description: |
//                         The Timestamp of the log
//                     message:
//                       type: string
//                       description: |
//                         The Message the user wants to show

// AWSRemoteTFState:
//   description: |
//     Representation of AWS remote tf state for external backend.
//   allOf:
//     - $ref: '#/definitions/ExternalBackendConfiguration'
//     - type: object
//       properties:
//         region:
//           type: string
//           description: |
//             The AWS region were the resource exists
//         bucket:
//           type: string
//           description: |
//             The AWS bucket containing objects
//         key:
//           type: string
//           description: |
//             The S3 Key uniquely identifies an object in a bucket
//         endpoint:
//           type: string
//           description: |
//             A custom endpoint for the S3 API (default: s3.amazonaws.com)
//         s3_force_path_style:
//           type: boolean
//           description: |
//             Always use path-style S3 URLs (https://<HOST>/<BUCKET> instead of https://<BUCKET>.<HOST>)
//         skip_verify_ssl:
//           type: boolean
//           description: |
//             Set this to `true` to not verify SSL certificates
//       required:
//         - region
//         - bucket
//         - key
// AzureRemoteTFState:
//   description: |
//     Representation of azure remote tf state for external backend.
//   allOf:
//     - $ref: '#/definitions/ExternalBackendConfiguration'
//     - type: object
//       properties:
//         container:
//           type: string
//           description: |
//             The Azure container were the resource exists
//         blob:
//           type: string
//           description: |
//             The Azure blob contained in the container
//       required:
//         - container
//         - blob
// GCPRemoteTFState:
//   description: |
//     Representation of GCP remote tf state for external backend.
//   allOf:
//     - $ref: '#/definitions/ExternalBackendConfiguration'
//     - type: object
//       properties:
//         bucket:
//           type: string
//           description: |
//             The GCP bucket containing objects
//         object:
//           type: string
//           description: |
//             The GCP object uniquely identifying an object in a bucket
//       required:
//         - bucket
//         - object

// SwiftRemoteTFState:
//   description: |
//     Representation of Swift remote tf state for external backend.
//   allOf:
//     - $ref: '#/definitions/ExternalBackendConfiguration'
//     - type: object
//       properties:
//         container:
//           type: string
//           description: |
//             The Swift container containing objects
//         object:
//           type: string
//           description: |
//             The swift object uniquely identifying an object in a container
//         skip_verify_ssl:
//           type: boolean
//           description: |
//             Set this to `true` to not verify SSL certificates
//         region:
//           type: string
//           description: |
//             The Swift region were the resource exists
//       required:
//         - container
//         - object
//         - region
