
build:
	go build -o terraform-provider-fake

test:
	cd fake && TF_ACC=1 go test -race -v .