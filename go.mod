module go-gin-smallproject

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/criminologiste/go-gin-smallproject/conf => ./go-gin-smallproject/conf
	github.com/criminologiste/go-gin-smallproject/middleware => ./go-gin-smallproject/middleware
	github.com/criminologiste/go-gin-smallproject/models => ./go-gin-smallproject/models
	github.com/criminologiste/go-gin-smallproject/pkg => ./go-gin-smallproject/pkg
	github.com/criminologiste/go-gin-smallproject/pkg/e => ./go-gin-smallproject/pkg/e
	github.com/criminologiste/go-gin-smallproject/pkg/setting => ./go-gin-smallproject/pkg/setting
	github.com/criminologiste/go-gin-smallproject/pkg/util => ./go-gin-smallproject/pkg/util
	github.com/criminologiste/go-gin-smallproject/routers => ./go-gin-smallproject/routers
	github.com/criminologiste/go-gin-smallproject/runtime => ./go-gin-smallproject/runtime
)
