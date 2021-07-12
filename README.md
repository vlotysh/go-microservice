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