#!make

# general variables
ENV?=development
PLAN_FILE?=tfplan
# AWS variables
AWS_ACCESS_KEY_ID?=
AWS_SECRET_ACCESS_KEY?=
AWS_DEFAULT_REGION?=us-east-1
# Terraform backend variables
AWS_BACKEND_BUCKET?=
AWS_BACKEND_ACCESS_KEY_ID?=
AWS_BACKEND_SECRET_ACCESS_KEY?=
AWS_BACKEND_REGION?=sa-east-1
# application variables
GOLANG_VERSION?=1.18-alpine

go-vet:
	docker container run --rm -v $$(pwd)/../../lambda-producer:/go/src/app -v producer-cache:/go/pkg/mod -w /go/src/app -e GOOS=linux -e GOARCH=amd64 -e GO111MODULE=on -e CGO_ENABLED=0 golang:$(GOLANG_VERSION) go vet .
	docker container run --rm -v $$(pwd)/../../lambda-consumer:/go/src/app -v consumer-cache:/go/pkg/mod -w /go/src/app -e GOOS=linux -e GOARCH=amd64 -e GO111MODULE=on -e CGO_ENABLED=0 golang:$(GOLANG_VERSION) go vet .

go-test:
	docker container run --rm -v $$(pwd)/../../lambda-producer:/go/src/app -v producer-cache:/go/pkg/mod -w /go/src/app -e GOOS=linux -e GOARCH=amd64 -e GO111MODULE=on -e CGO_ENABLED=0 golang:$(GOLANG_VERSION) sh -c "go test tests/**/*_test.go -v"
	docker container run --rm -v $$(pwd)/../../lambda-consumer:/go/src/app -v consumer-cache:/go/pkg/mod -w /go/src/app -e GOOS=linux -e GOARCH=amd64 -e GO111MODULE=on -e CGO_ENABLED=0 golang:$(GOLANG_VERSION) sh -c "go test tests/**/*_test.go -v"

compile: go-vet go-test
	docker container run --rm -v $$(pwd)/../../lambda-producer:/go/src/app -v producer-cache:/go/pkg/mod -w /go/src/app -e GOOS=linux -e GOARCH=amd64 -e GO111MODULE=on -e CGO_ENABLED=0 golang:$(GOLANG_VERSION) go build -ldflags="-s -w" -o producer
	mv ../../lambda-producer/producer producer
	docker container run --rm -v $$(pwd)/../../lambda-consumer:/go/src/app -v consumer-cache:/go/pkg/mod -w /go/src/app -e GOOS=linux -e GOARCH=amd64 -e GO111MODULE=on -e CGO_ENABLED=0 golang:$(GOLANG_VERSION) go build -ldflags="-s -w" -o consumer
	mv ../../lambda-consumer/consumer consumer

clean:
	docker volume rm -f producer-cache
	docker volume rm -f consumer-cache

create-backend-bucket:
	aws s3 ls s3://$(AWS_BACKEND_BUCKET) --region $(AWS_BACKEND_REGION) || \
		aws s3api create-bucket --bucket $(AWS_BACKEND_BUCKET) --region $(AWS_BACKEND_REGION) --acl private

delete-backend-bucket:
	aws s3 rm --recursive s3://$(AWS_BACKEND_BUCKET) --region $(AWS_BACKEND_REGION)
	aws s3api delete-bucket --bucket $(AWS_BACKEND_BUCKET) --region $(AWS_BACKEND_REGION)

create-workspace:
	-@terraform workspace new development 2> /dev/null
	-@terraform workspace new staging 2> /dev/null
	-@terraform workspace new production 2> /dev/null
	terraform workspace select $(ENV)

init: create-backend-bucket
	terraform init -upgrade=true \
		-backend-config="bucket=$(AWS_BACKEND_BUCKET)" \
		-backend-config="key=state.tfstate" \
		-backend-config="region=$(AWS_BACKEND_REGION)" \
		-backend-config="workspace_key_prefix=terraform/lambda-image-recognition-kinesis" \
		-backend-config="access_key=$(AWS_BACKEND_ACCESS_KEY_ID)" \
		-backend-config="secret_key=$(AWS_BACKEND_SECRET_ACCESS_KEY)" \
		-backend-config="encrypt=true"
	make create-workspace
	make plan

tf-fmt:
	terraform fmt -write=true -recursive

tf-validate:
	terraform validate

tf-lint:
	docker container run --rm -t \
		--name tflint \
		--env TFLINT_LOG=debug \
		-v $$(pwd):/data \
		ghcr.io/terraform-linters/tflint

refresh:
	terraform refresh -var-file=$(ENV).tfvars

output:
	terraform output

plan: tf-validate tf-lint
	terraform plan \
		-out=$(PLAN_FILE) \
		-var-file=$(ENV).tfvars \
		-input=false

apply:
	terraform apply $(PLAN_FILE)

destroy:
	terraform destroy \
		-var-file=$(ENV).tfvars \
		-auto-approve
	rm -rf tfplan consumer consumer.zip producer producer.zip .terraform .terraform.lock.hcl
