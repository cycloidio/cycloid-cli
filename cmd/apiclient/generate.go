package apiclient

//go:generate go tool mockgen -destination=mock/middleware.go -mock_names=Middleware=Middleware -package=mock -source=middleware.go
