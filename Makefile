BINARY_NAME=signer.app
OUTPUT_DOC=output.docx

build:
	GOOS=linux go build -o ./bin/${BINARY_NAME} ./pkg/program.go

run: build
	./bin/${BINARY_NAME}

cls:
	rm ${BINARY_NAME}