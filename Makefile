default: compile

deps:
	go get ./...
	go install honnef.co/go/tools/cmd/staticcheck@latest

compile: deps
	go build -o terraform-provider-meltwater

qa: vet lint test

test:
	go test -v ./...

vet:
	go vet ./...

lint:
	staticcheck -tests -fail all -f stylish ./...

install: compile
	mkdir -p .terraform/plugins/meltwater.com/meltwater/meltwater/0.0.1/linux_amd64
	cp ./terraform-provider-meltwater .terraform/plugins/meltwater.com/meltwater/meltwater/0.0.1/linux_amd64/

#test: compile
#	terraform init
#	terraform fmt
#	terraform plan -out terraform.tfplan
#	terraform apply terraform.tfplan
