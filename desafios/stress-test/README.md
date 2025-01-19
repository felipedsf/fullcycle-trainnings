# Stress Test Tool
___
## Parametros default:
>> url = https://www.webpagetest.org/
>
>> requests = 3000
> 
>> concurrency = 100


## Execução da aplicação:
### Via Docker:

> docker run IMAGEM —url=http://google.com —requests=1000 —concurrency=10

### Via Go Local:

> go run ./cmd/main.go —url=http://google.com —requests=1000 —concurrency=10
___