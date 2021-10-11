# Cycloid Developer Tips

**NOTE!** These are internal notes meant for cycloid develovers only since they require access to private git repositories.

This file gives some tips to how to test, change or upgrade cli, for cycloid developers.

The cli pipeline is available in the [cycloid-stacks](https://github.com/cycloidio/cycloid-stacks/tree/stacks/cycloid-cli) and the prod and staging cli can be checked [here](https://console.cycloid.io/organizations/cycloid/projects/cycloid-cli).


## Update cli client 

To update the cli client you can either use either:

- the **latest swagger version stored in the the [cycloid docs](https://docs.cycloid.io/api/swagger.yml)**. For that, you just need to run ```make generate-client```, which will run a docker local swagger,generate the new client version and then commit it to the cli git repo where it will be tested in the cli pipeline. </br>
**Note!** After running the command you should validate that the e2e testing has been succefull and then if no issue has been found, you should edit the CHANGELOG.md to add the new version of swagger.

- a **copied swagger file**, if you wish to use another version of the swagger file, you can create a folder called /gen-swagger and copy it inside. Once that's finished you can run ```make generate-local-client``` that will create the new client version files. After, you should change the client/version file with the corresponding BE version.</br> 


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

    2. ```API_LICENCE_KEY```: you can find the e2e license to use in the cycloid console. The name can be retrieved on the [pipeline variables file](https://github.com/cycloidio/cycloid-stacks/blob/config/cycloid-cli/pipeline/staging/variables.yaml#L28) in the config repository of the cli pipeline of staging.

4. Then all you need to do is launch ```make local-ci-test```, which will launch a local BE server using fake generated data and then launch the e2e tests using the following variables, as required and defined in [e2e/e2e.go](e2e/e2e.go):
  
    - **CY_API_URL** : corresponding to the BE API URL. Set by default has `http://172.42.0.3:3001`

    - **CY_TEST_ROOT_API_KEY** : the cycloid API_KEY generated in the BE repository. Set by default has `cat ${LOCAL_BE_GIT_PATH}/API_KEY)`

    - **CY_TEST_GIT_CR_URL** : The local git server launched in the docker-compose. Set by default has `git@172.42.0.14:/git-server/repos/backend-test-config-repo.git`

5. To re-run the tests, you can just re-run the previous command has many times has you want.

6. Once you're finished with the tests you can run ```make delete-local-be```, that will delete all the created docker instances.