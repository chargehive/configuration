module github.com/chargehive/configuration

go 1.24.0

toolchain go1.24.4

require (
	github.com/chargehive/proto v1.15.2
	github.com/go-playground/assert/v2 v2.2.0
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.15.4
)

// validator 10.15 breaks diving
replace github.com/go-playground/validator/v10 => github.com/go-playground/validator/v10 v10.14.0

require (
	github.com/chargehive/grpc v0.0.0-20260213121241-9642c028c9d2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gabriel-vasile/mimetype v1.4.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/crypto v0.43.0 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)
