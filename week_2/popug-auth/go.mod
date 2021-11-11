module github.com/vctrl/async-architecture/week_2/popug-auth

go 1.17

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/macinnir/jose v1.1.1
	github.com/satori/go.uuid v1.2.0
	github.com/vctrl/async-architecture/week_2/schema v0.0.0-20211109231557-f3e1a85f5c51
	google.golang.org/grpc v1.42.0
	gorm.io/driver/sqlite v1.2.3
	gorm.io/gorm v1.22.2
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/vctrl/async-architecture/week_2/schema => ../schema
