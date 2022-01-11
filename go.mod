module github.com/koodeex/allure-testify

go 1.16

replace (
	github.com/koodeex/allure-testify/pkg/allure => ./pkg/allure
	github.com/koodeex/allure-testify/pkg/framework => ./pkg/framework
	github.com/koodeex/allure-testify/pkg/provider => ./pkg/provider
)

require (
	github.com/koodeex/allure-testify/pkg/allure v0.1.1
	github.com/koodeex/allure-testify/pkg/framework v0.1.1
	github.com/koodeex/allure-testify/pkg/provider v0.1.1
	github.com/stretchr/testify v1.7.0
)
