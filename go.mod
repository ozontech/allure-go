module github.com/ozontech/allure-go

go 1.17

replace (
	github.com/ozontech/allure-go/pkg/allure => ./pkg/allure
	github.com/ozontech/allure-go/pkg/framework => ./pkg/framework
)

require (
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733
	github.com/ozontech/allure-go/pkg/allure v0.6.9
	github.com/ozontech/allure-go/pkg/framework v0.6.23
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)
