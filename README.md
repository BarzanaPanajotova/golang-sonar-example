## Golang Sonar Example

To run the tests execute from the root directory
```shell
ginkgo -cover -coverprofile=coverage.out -outputdir=.coverage ./...   
```

To run sonar analysis on an existing SonarQube server:
1. Edit sonar-scanner.properties with your respective config
2. Run the following:
   ```shell
    sonar-scanner -Dproject.settings=sonar-scanner.properties 
    ```