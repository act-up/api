module github.com/act-up/api/config

go 1.13

require (
	github.com/act-up/api/controllers v0.0.0-20200410130308-9b4a3811ce8c
	github.com/act-up/api/models v0.0.0-20200410130308-9b4a3811ce8c // indirect
	github.com/gin-gonic/gin v1.6.2
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.3.0
)

replace github.com/act-up/api/controllers => /Users/antonellawilby/ActUp/api/controllers
replace github.com/act-up/api/models => /Users/antonellawilby/ActUp/api/models
