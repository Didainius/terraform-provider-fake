
build:
	go build -o terraform-provider-fake

test:
	@echo "Sometimes it must be run multiple times as it doesn't always hit, but I would say 3 out of 5 runs do"
	cd fake && TF_ACC=1 go test -race -v .