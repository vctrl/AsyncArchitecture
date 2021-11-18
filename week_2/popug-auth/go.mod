module github.com/vctrl/async-architecture/week_2/popug-auth

go 1.17

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/macinnir/jose v1.1.1
	github.com/satori/go.uuid v1.2.0
	github.com/vctrl/async-architecture/week_2/schema v0.0.0-20211109231557-f3e1a85f5c51
	google.golang.org/grpc v1.42.0
	gorm.io/driver/postgres v1.2.1
	gorm.io/gorm v1.22.2
)

require (
	github.com/confluentinc/confluent-kafka-go v1.7.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/vctrl/async-architecture/week_2/schema => ../schema
