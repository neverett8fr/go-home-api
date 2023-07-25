# home-service  
  
## add endpoint  

`{{host}}:8081/endpoint/name1`  

```go
{
 "method": "GET",
 "route": "http://localhost:8080/test/pi"    
}
```  

## add condition  

`{{host}}:8081/condition/name1`  

```go
{
 "condition": "always"
}
```  

## restart  

`{{host}}:8081/restart`  

## start  

`{{host}}:8081/start`  
