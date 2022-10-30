SERVICE_NAME := "resexample"
CURRENT_DIR = $(shell pwd)
GOPATH = $(shell echo ${HOME})
ifdef BUILD_VERSION
	VERSION := "-$(BUILD_VERSION)"
else
	VERSION := ""
endif

.SILENT:

deps:
	cd ${CURRENT_DIR}/app && go mod download

clean:
	rm -rf ./.bin/${SERVICE_NAME}


build: clean deps
	cd ${CURRENT_DIR}/app && CGO_ENABLED=0 GOOS=darwin go build -o ./.bin/terraform-provider-${SERVICE_NAME} ./cmd/${SERVICE_NAME}/main.go

move:
	mkdir -p ~/.terraform.d/plugins/k11s.ru/custom-provider/resexample/1.0.0/darwin_amd64/
	mv ${CURRENT_DIR}/app/.bin/terraform-provider-resexample ~/.terraform.d/plugins/k11s.ru/custom-provider/resexample/1.0.0/darwin_amd64/terraform-provider-resexample

chmod:
	chmod +x ~/.terraform.d/plugins/k11s.ru/custom-provider/resexample/1.0.0/darwin_amd64/terraform-provider-resexample