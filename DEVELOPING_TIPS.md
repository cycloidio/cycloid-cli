# Cycloid Developer Tips

**NOTE!** These are internal notes meant for cycloid develovers only since they require access to private git repositories.

This file gives some tips to how to test, change or upgrade cli, for cycloid developers.

The cli pipeline is available in the [cycloid-stacks](https://github.com/cycloidio/cycloid-stacks/tree/stacks/cycloid-cli) and the prod and staging cli can be checked [here](https://console.cycloid.io/organizations/cycloid/projects/cycloid-cli).


## Update cli client 

To update the cli client you can either use either:

- the **latest swagger version stored in the the [cycloid docs](https://docs.cycloid.io/api/swagger.yml)**. For that, you just need to run ```make generate-client-from-docs```, which will run a docker local swagger,generate the new client version and then commit it to the cli git repo where it will be tested in the cli pipeline. </br>
**Note!** After running the command you should validate that the e2e testing has been succefull and then if no issue has been found, you should edit the CHANGELOG.md to add the new version of swagger.

- a **copied swagger file**, if you wish to use another version of the swagger file, you can create a folder called /gen-swagger and copy it inside. Once that's finished you can run ```make generate-client-from-local``` that will create the new client version files. After, you should change the client/version file with the corresponding BE version.</br> 

## Add nem commands to the cli

To add a new command you can start by verifying the api endpoint to implement in the [docs](https://docs.cycloid.io/api/index.html). And then you should:

1. Define the method that implements the endpoint in ```cmd/middleware/<feature>.go```. </br>
    
    1. In this method you start by defining the set of params required for the endpoint as defined in ```client/client/<feature>/<method_parameters>.go``` </br> 
    
    2. Then, you call the api equivalent method and then return the reply.</br> **Note!** The method should always return error and depending on the method you may also be required to return an object. You can check the available object structs in the folder ```client/models```.
    Here's some general logic: 
        - delete methods -> return error only
        - create/get methods -> return *models.object, error
        - list methods -> return []*models.object, error</br> </br>


2. Once you're method is defined you should add it to the middleware interface at ```cmd/middleware/middleware.go``` that implements the api endpoints.</br> 

3. Now you need to define the cobra command at ```cmd/cycloid/<feature>/``` so that the method previously defined can be used in the cli. In this folder, you find the files: ```cmd.go``` where you define the available cobra commands for the feature, the ```common.go``` where you define the flag methods to associate with the cobra commands and then a set of files, one per command of the feature.

    1. You should start by adding the type of command to the set of list of available commands in ```cmd.go```

    2. Then create a new file on this folder, where you will specify the method that will return the cobra command that defines the cli command to implement with the flags possible to use. </br> **Note** You can define the flags as required or not for that you can use the methods specified in ```cmd/cycloid/common```

    3. Then you define a method that will be run by this cobra command that will take as argument the multiple flags and pass them to the middleware interface method that you previsouly defined in step 1.

4. To validate the created command you can build a local version of the cli using ```make build```and then use the generated binary to test the created commands. </br> **Tip!** Don't forget to login in the cli using ```./ci login --org cycloid-sandbox --api-key API_KEY```. You can generate a temporary key on the staging organisation of cycloid using the console.

5. Finally, create the correspond e2e tests in ```e2e/<feature>_test.go``` and add a new entry on the changelog. To see how to launch tests locally you can check the CLI local testing section.


## CLI local testing 
### Requirements

To perform local test you need:
- aws installed locally ([official doc](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)), required to connect to the cycloid docker images stored in a ECR
- [backend repo local clone](https://github.com/cycloidio/youdeploy-http-api), required to launch local be

### Procedure

In order to automate the procedure has much as possible we tried to create some Makefile targets to help making the testing easier and faster, but it still requires some manual actions.

To launch the e2e tests, locally you should:

1. Connect to cycloid's ECR, by typing ```make ecr-connect```, it will configure aws cli with the cycloid credentials retrieved from vault and login to the ECR

2. Clone BE git repository and change the variable ```LOCAL_BE_GIT_PATH``` in the makefile acordingly

3. To run the BE infra we need to change the following variables in the Makefile:
    1. ```YD_API_TAG```: set by default has ```staging``` in the Makefile, so no need to change it. However, you can change it to a specific version, if required. You can check the available versions in the backend repository or pipeline and as well on the AWS ECR directly. Tip: use the following command to list the available images : ```aws ecr list-images --repository-name youdeploy-http-api``` 

    2. **[/!\ Required]** ```API_LICENCE_KEY```: you can find the e2e license to use in the cycloid console. The name can be retrieved on the [pipeline variables file](https://github.com/cycloidio/cycloid-stacks/blob/config/cycloid-cli/pipeline/staging/variables.yaml#L28) in the config repository of the cli pipeline of staging.

4. Then all you need to do is launch ```make local-e2e-test```, which will launch a local BE server using fake generated data and then launch the e2e tests using the following variables, as required and defined in [e2e/e2e.go](e2e/e2e.go):
  
    - **CY_API_URL** : corresponding to the BE API URL. You should change it in the Makefile to `http://172.42.0.3:3001`

    - **CY_TEST_ROOT_API_KEY** : the cycloid API_KEY generated in the BE repository. Set by default has `cat ${LOCAL_BE_GIT_PATH}/API_KEY)`

    - **CY_TEST_GIT_CR_URL** : The local git server launched in the docker-compose. You should change it in the Makefile to `git@172.42.0.14:/git-server/repos/backend-test-config-repo.git`

5. To re-run the tests, you can just re-run the previous command has many times has you want.

6. Once you're finished with the tests you can run ```make delete-local-be```, that will delete all the created docker instances.
