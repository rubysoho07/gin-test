# Gin-Test

Gin을 테스트 해 봅니다. 

## 초기 설정

```shell
go mod init gin-test
go get -u github.com/gin-gonic/gin
```

`go.mod`, `go.sum` 파일이 생성된 것을 확인합니다. 

## 테스트 파일 만들기 

`main.go` 파일을 만들고, [Getting Started](https://gin-gonic.com/docs/quickstart/#getting-started) 문서의 `example.go` 파일 내용을 복사합니다. 

바로 테스트 하려면 `go run main.go` 명령을 입력하고, 웹브라우저에서 `localhost:8080/ping`으로 접속합니다. 

## 빌드하기

```shell
go build
```

`gin-test`라는 실행 파일이 생성된 것을 볼 수 있습니다.

## 참고자료

* [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
* [Gin Web Framework: Quickstart](https://gin-gonic.com/docs/quickstart/)