module github.com/chargehive/configuration

go 1.19

require (
	github.com/chargehive/proto v1.11.0
	github.com/go-playground/assert/v2 v2.2.0
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.15.4
)

// validator 10.15 breaks diving
replace github.com/go-playground/validator/v10 => github.com/go-playground/validator/v10 v10.14.0

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)
