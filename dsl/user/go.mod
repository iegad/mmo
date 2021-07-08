module github.com/iegad/mmo/dsl/user

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/iegad/kraken v0.0.1
	github.com/iegad/mmo v0.0.1
	github.com/olivere/elastic/v7 v7.0.26
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace (
	github.com/iegad/kraken v0.0.1 => ../../../kraken
	github.com/iegad/mmo v0.0.1 => ../../
)
