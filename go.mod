module github.com/cycloidio/cycloid-cli

go 1.24.1

require (
	dario.cat/mergo v1.0.2
	github.com/adrg/xdg v0.5.3
	github.com/go-openapi/errors v0.22.1
	github.com/go-openapi/runtime v0.28.0
	github.com/go-openapi/strfmt v0.23.0
	github.com/go-openapi/swag v0.23.1
	github.com/go-openapi/validate v0.24.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pkg/errors v0.9.1
	github.com/sanity-io/litter v1.5.8
	github.com/spf13/cobra v1.9.2-0.20250531123604-6dec1ae26659
	github.com/spf13/viper v1.20.1
	github.com/stretchr/testify v1.10.0
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.23.0 // indirect
	github.com/go-openapi/inflect v0.21.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.1 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/loads v0.22.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-swagger/go-swagger v0.31.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jessevdk/go-flags v1.5.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/sagikazarmark/locafero v0.9.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.14.0 // indirect
	github.com/spf13/cast v1.8.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/toqueteos/webbrowser v1.2.0 // indirect
	go.mongodb.org/mongo-driver v1.17.3 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// replace github.com/cycloidio/youdeploy-http-api => /home/stammfrei/projects/cycloid/youdeploy-http-api
// // To remove the panic issue of using TF
// //replace github.com/hashicorp/terraform => github.com/cycloidio/terraform v0.13.5-cy
// replace github.com/hashicorp/terraform => github.com/cycloidio/terraform v1.4.6-cy.x
//
// replace github.com/simpleforce/simpleforce => github.com/cycloidio/simpleforce v0.0.0-20210521130438-36c1102a5fba
//
// replace github.com/spf13/afero v1.6.0 => github.com/spf13/afero v1.2.2
//
// // Terracognita replaces copied
// // Force an specific version if not the AWS provider does not compile
// replace github.com/hashicorp/aws-sdk-go-base v0.6.0 => github.com/hashicorp/aws-sdk-go-base v0.5.0
//
// // If we  go to the 1.5.0 then github.com/hashicorp/terraform-plugin-test/ will break
// // as go-getter introduced a break from 1.4 -> 1.5
// replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
//
// // Fork of Azurerm that has the V2 of the SDK
// //replace github.com/terraform-providers/terraform-provider-azurerm => github.com/cycloidio/terraform-provider-azurerm v1.44.1-0.20210517111036-df0beb5af9c3
//
// replace github.com/hashicorp/terraform-provider-azurerm => github.com/cycloidio/terraform-provider-azurerm v1.44.1-0.20230517144901-90a36c6b8ed4
//
// replace github.com/hashicorp/terraform-provider-aws => github.com/cycloidio/terraform-provider-aws v1.60.1-0.20220513132327-e2dbdf90e533
//
// replace github.com/hashicorp/terraform-provider-google => github.com/hashicorp/terraform-provider-google v1.20.1-0.20220201002249-bc5fcb3c89a5
//
// replace github.com/concourse/concourse => github.com/cycloidio/concourse v1.6.1-0.20240611102233-2c09bb557bc9
//
// // Required by TerraCost
// replace github.com/gruntwork-io/terragrunt => github.com/cycloidio/terragrunt v0.0.0-20230905115542-1fe1ff682fd9
//
// replace cloud.google.com/go/bigquery => cloud.google.com/go/bigquery v1.8.0
//
// replace cloud.google.com/go/storage => cloud.google.com/go/storage v1.16.0
//
// replace google.golang.org/api => google.golang.org/api v0.74.0
//
// // End of fixed versions after TerraCost upgrade
//
// replace github.com/aws/aws-sdk-go v1.53.21 => github.com/aws/aws-sdk-go v1.43.34

tool github.com/go-swagger/go-swagger/cmd/swagger
