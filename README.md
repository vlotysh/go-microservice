# Go microservice test project

## Run project

1. Docker build
```text  
docker build -t gomicroservice . 
```  

2. Docker run
```text  
docker run -it -p 8081:8081  --rm --name gomicroservice gomicroservice  
```

M1 Develop run steps
```text 
[1] Download and install the ARM64 installer package from https://golang.org - https://go.dev/dl/go1.18.darwin-arm64.pkg
[2] Run go install github.com/go-delve/delve/cmd/dlv@latest from the command-line.
[3] Run go install github.com/aarzilli/gdlv@latest from the command line.
```