[ ] CLI default url should be the prod SaaS one
[ ] Adding output after a login to confirm to the user it's successful login or not
[ ] Make the binary able to check api version and is own builded version to let a warning message you should update it cause the BE version does not match the binairy version
[X] Move and make sure all commands are migrated to middleware
[ ] Put plugins with xdg into home
[ ] Make cycloid url env var or --api-url
[ ] Implement diff pipeline display (same as fly) using see https://github.com/aryann/difflib
[ ] Work on login part -> have a look to login per orgs
  [ ] (Get org if not logged in fail. Should we remove --org and get it from the token ? Or use a user token. And let the cli do login each time on org)
[ ] Create a proper error function to handle or display error correctly (thinking about error message/details in payload)
[ ] Work on output formats
      --field name --field bar --no-header --raw --separator :
      |    name   |  sdfsdf |
       fofo            foo
      fofo:foo
[ ] Define a bit what should be in env vars. Mostly thinking about project/env/org (to be pushed by the pipeline)
[ ] (Keep canonical or name when you need an identifier for delete or create ?))
      cy project delete --project
      cy project delete --name
      cy project delete --canonical
[ ] (Make usage of name possible "canonical like" when there is only ID. Using helper search. Return error if 2 found)
[ ] Download of plugin how to ? embeed 3 last one ? https://github.com/markbates/pkger
[ ] Implement e2e tests
[ ] Create a build pipeline
[ ] Work on help strings
[ ] Add support for others creds types ?
[ ] (Implement stack download: download a stack pipeline template, vars and config samples)
[ ] (Implement create project from downloadded stack. It automatically do the call using the samples on the filesystem)
[ ] Migrate the code into youdeploy api ?
[ ] work on logger (--debug / verbose)
[ ] Migrate /define to git tag
[ ] Reduce swagger generated files
[ ] Add a logger into the cli to be able to display --debug mode
[ ] Feature idea: jobs time history. Basic ascii histogram to display job time/status
[ ] Implement extra flags ?
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// rootCmd.PersistentFlags().Bool("version", false, "Display the version of this tool.")

	// rootCmd.PersistentFlags().BoolP("quiet", "q", false, "....") // also -q ?
	// viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))

	// rootCmd.PersistentFlags().Bool("debug", false, "Turn on debug logging.")
	// viper.BindPFlag("useDebug", rootCmd.PersistentFlags().Lookup("debug"))

	// rootCmd.PersistentFlags().Bool("no-verify-ssl", false, ".....")
	// viper.BindPFlag("noVerifySSL", rootCmd.PersistentFlags().Lookup("noVerifySSL"))

	// --log-http
	//    Log all HTTP server requests and responses to stderr. Overrides the
	//    default core/log_http property value for this command invocation.
[ ] Create alias on all commands to make them singular or plurial like : user/users


Example of all working calls (tested on staging)

export CY_API_URL=https://http-api-staging.cycloid.io
export TOKEN=

# External BE
V=1 ./cy  external-backends create logs ElasticsearchLogs eb2  --project gaeltest --org seraf --env dev --cred 743 --url http://test --prefilter foo=bar
V=1 ./cy  external-backends create events AWSCloudWatchLogs --org seraf --region eu-west-1 --cred 63
V=1 ./cy --env prod --org seraf --project gaeltest external-backends create infraview SwiftRemoteTFState  --cred 767 --region bar
V=1 ./cy --org seraf  external-backends list

# Creds
V=1 ./cy cred  --org seraf create ssh --name gaeltestauth  --ssh-key /home/gael/.ssh/id_rsa_qapa_enovance
V=1 ./cy cred  --org seraf create basic_auth --name gaeltestauth --username foo --password bar
V=1 ./cy cred  --org seraf create custom --name gaeltest  --field bar=bli --field truch=much
V=1 ./cy cred  --org seraf list
V=1 ./cy cred  --org seraf delete --id 882

# Events
V=1 ./cy events --org seraf send --title foo --tag foo=bar --message=/tmposfsdf

# Project
V=1 ./cy project --org seraf list
V=1 ./cy project --org seraf get --project "gaeltest"
V=1 ./cy project --org seraf list-pipelines --project "gaeltest"

V=1 ./cy project --org seraf  create --name Gael --cloud-provider google --stack-ref seraf:stack-dummy \
--config-repo 307 --env test \
--vars /home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/variables.sample.yml \
--pipeline /home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/pipeline.yml \
--config '/home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/variables.sample.yml=($ project $)/pipeline/variables-($ environment $).yml'

V=1 ./cy project --org seraf create-env --project gael --env new  \
--vars /home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/variables.sample.yml \
--pipeline /home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/pipeline.yml  \
--config '/home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/variables.sample.yml=($ project $)/pipeline/variables-($ environment $).yml'

V=1 ./cy project --org seraf delete-env --project gael --env new
V=1 ./cy project --org seraf  delete --project gael

# Org
V=1 ./cy org list
V=1 ./cy org get --org seraf
V=1 ./cy org list-workers --org seraf

# Pipeline
V=1 ./cy pipeline unpause --org seraf --env new --project gael
V=1 ./cy pipeline pause --org seraf --env new --project gael

V=1 ./cy pipeline pause-job --org seraf --env new --project gael --job job-hello-world
V=1 ./cy pipeline unpause-job --org seraf --env new --project gael --job job-hello-world
V=1 ./cy pipeline trigger-build --org seraf --env new --project gael --job job-hello-world

V=1 ./cy pipeline clear-task-cache --org seraf --env new --project gael --job job-hello-world --task hello-world

V=1 ./cy pipeline get-job --org cycloid-demo --env demo --project orange --job terraform-plan
V=1 ./cy pipeline list-jobs --org cycloid-demo --env demo --project orange
V=1 ./cy pipeline list-builds --org cycloid-demo --env demo --project orange --job terraform-plan

V=1 ./cy pipeline diff --org cycloid-demo --env demo --project orange \
--vars /home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/variables.sample.yml \
--pipeline /home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/pipeline.yml

V=1 ./cy pipeline update --org seraf --project gael --env new \
--vars /home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/variables.sample.yml \
--pipeline /home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/pipeline.yml \
--config '/home/gael/Desktop/git/github/cycloidio/stack-dummy/pipeline/variables.sample.yml=($ project $)/pipeline/variables-($ environment $).yml'

V=1 ./cy stack  --org seraf  list
V=1 ./cy stack  --org seraf  get --ref seraf:relationships-summit-661

V=1 ./cy version
