module yeat.dev/shortlink

go 1.18

require github.com/uptrace/bunrouter v1.0.14

require (
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.18.1 // indirect
)

require (
	github.com/go-redis/redis v6.15.9+incompatible
	yeat.dev/shortlink/handlers v0.0.0-00010101000000-000000000000
)

replace yeat.dev/shortlink/handlers => ./handlers
