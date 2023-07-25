# home-service  

## test  

request  
`{{host}}:8081/test/asdf  GET`  
response  

```go
{
    "errors": null,
    "data": "teststring"
}
```

## add endpoint  

request  
`{{host}}:8081/endpoint/name1  POST`  

```go
{
 "method": "GET",
 "route": "http://localhost:8080/test/pi"    
}
```  

response  

```go
{
    "errors": [
        null
    ],
    "data": "endpoint added"
}
```

## add condition  

request  
`{{host}}:8081/condition/name1  POST`  

```go
{
 "condition": "always"
}
```  

response  

```go
{
    "errors": [
        null
    ],
    "data": "condition added"
}
```

## restart  

request  
`{{host}}:8081/restart  POST`  
response  

```go
{
    "errors": [
        null
    ],
    "data": "service restarted"
}
```

## start  

request  
`{{host}}:8081/start  POST`  
response  

```go
{
    "errors": [
        null
    ],
    "data": "service started"
}
```
