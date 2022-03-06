module github.com/Silencezzzz/go-gin-example

go 1.16

require (
	github.com/Shopify/sarama v1.31.1
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.0
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/Silencezzzz/go-gin-example/conf => ./pkg/conf
	github.com/Silencezzzz/go-gin-example/kafka => ./kafka
	github.com/Silencezzzz/go-gin-example/middleware => ./middleware
	github.com/Silencezzzz/go-gin-example/models => ./models
	github.com/Silencezzzz/go-gin-example/pkg/e => ./pkg/e
	github.com/Silencezzzz/go-gin-example/pkg/setting => ./pkg/setting
	github.com/Silencezzzz/go-gin-example/pkg/util => ./pkg/util
	github.com/Silencezzzz/go-gin-example/routers => ./routers

)
