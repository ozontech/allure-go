module github.com/ozontech/allure-go/pkg/framework

go 1.17

replace github.com/ozontech/allure-go/pkg/allure => ../allure

require (
	github.com/ozontech/allure-go/pkg/allure v0.6.5
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)
