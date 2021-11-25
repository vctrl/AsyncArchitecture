module github.com/vctrl/async-architecture/popug-tasks

go 1.17

require (
	github.com/confluentinc/confluent-kafka-go v1.7.0
	github.com/golang/protobuf v1.5.2
	github.com/satori/go.uuid v1.2.0
	github.com/vctrl/async-architecture/schema v0.0.0-20211124023330-5bab94283b9e
	google.golang.org/grpc v1.42.0
	gorm.io/driver/postgres v1.2.2
	gorm.io/gorm v1.22.3
)

require (
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

replace github.com/vctrl/async-architecture/schema => ../schema