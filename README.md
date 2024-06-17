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

## AWS SDK 이용하기

핵심 SDK 모듈 다운받기

```shell
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
```

서비스별 SDK 다운받기 ([전체 리스트](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2#section-readme))

```shell
# S3 기준
go get github.com/aws/aws-sdk-go-v2/service/s3
```

[문서](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#invoke-an-operation)를 참고하여 코드를 작성하면 됨

## MySQL 연결하기

Dependency 받기

```shell
go get -u github.com/go-sql-driver/mysql
```

임시 DB 실행하기 (Docker Compose 필요)

```shell
docker-compose up -d

# 테스트 테이블 생성 후 임시 데이터 집어넣기
docker exec -i gin-test-mysql mysql  --password=example < test_data.sql
```

CRUD 관련 내용은 `database.go` 파일 참조

## Redis 실행하기

Dependency 설정

```shell
go get github.com/redis/go-redis/v9
```

* Docker Compose로 Redis를 실행할 수 있도록 해 둠

관련 내용은 `redis_example.go` 파일 참조

## Swagger 사용하기

```shell
# swag 명령어 설치 ($HOME/go/bin 경로가 $PATH에 있는지 확인할 것)
go install github.com/swaggo/swag/cmd/swag@latest

# 프로젝트 루트 디렉토리로 이동 후 초기화
swag init

# gin-swagger 패키지 다운로드
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

이후 내용은 [gin-swagger](https://github.com/swaggo/gin-swagger) 프로젝트의 문서 참조. 작성 규칙은 [링크](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format) 참조

실행하려면 `localhost:8080/swagger/index.html`으로 이동하면 됨

## 참고자료

* [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
* [Gin Web Framework: Quickstart](https://gin-gonic.com/docs/quickstart/)
* [Getting Started with the AWS SDK for Go V2](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/)
* [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql?tab=readme-ov-file)
* [Go Redis Documentation](https://redis.uptrace.dev/guide/go-redis.html)
* [GitHub gin-contrib/sessions](https://github.com/gin-contrib/sessions): Gin에서 세션 관리할 때 사용
* [gin-swagger](https://github.com/swaggo/gin-swagger)