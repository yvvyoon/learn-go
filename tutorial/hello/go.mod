module hello

go 1.12

// example.com라는 모듈 경로가 실제로 없기 때문에 로컬 경로로 replace
replace example.com/greetings => ../greetings

// 현재 디렉토리에서 go build를 수행하면 아래 코드가 생성됨
// go build로 컴파일된 hello가 example.com/greetings에 종속성을 가진다는 의미
require example.com/greetings v0.0.0-00010101000000-000000000000
