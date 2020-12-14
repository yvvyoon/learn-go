# learn-go

### go build vs go install
* go build

    * 작성한 .go 파일을 컴파일하여 build 커맨드를 수행한 디렉토리에 생성

* go install

    * 작성한 .go 파일을 컴파일하여 `GOBIN` 위치에 생성
    * `$GOBIN` = `$GOPATH/bin`
    * 라이브러리 파일의 경우 `.a` 형태로 `$GOPATH/pkg` 위치에 생성되고, 컴파일이 캐싱되므로 다음 컴파일 시 변경된 부분만 수행