module github.com/titrxw/smart-home-server

go 1.19

require (
	github.com/eclipse/paho.mqtt.golang v1.4.1
	github.com/gin-contrib/sessions v0.0.5
	github.com/gin-gonic/gin v1.9.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.14.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/gookit/color v1.5.3
	github.com/gookit/goutil v0.5.14
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/mojocn/base64Captcha v1.3.5
	github.com/redis/go-redis/v9 v9.0.2
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.15.0
	github.com/titrxw/emqx-sdk v0.0.0-20211130142505-b7b21dcc04c4
	github.com/we7coreteam/w7-rangine-go v1.0.16
	github.com/we7coreteam/w7-rangine-go-support v1.0.1
	gorm.io/gorm v1.25.0
)

replace github.com/we7coreteam/w7-rangine-go => ../w7-rangine-go

require (
	github.com/asaskevich/EventBus v0.0.0-20200907212545-49d423059eef // indirect
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golobby/container/v3 v3.0.2 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/imroc/req v0.3.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jedib0t/go-pretty/v6 v6.4.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.9 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/quasoft/memstore v0.0.0-20191010062613-2bce066d2b0b // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/sevlyar/go-daemon v0.1.6 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/wader/gormstore/v2 v2.0.3 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/image v0.1.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/datatypes v1.1.1-0.20230130040222-c43177d3cf8c // indirect
	gorm.io/driver/mysql v1.4.6 // indirect
	gorm.io/driver/sqlite v1.5.0 // indirect
	gorm.io/gen v0.3.22 // indirect
	gorm.io/hints v1.1.0 // indirect
	gorm.io/plugin/dbresolver v1.3.0 // indirect
)
