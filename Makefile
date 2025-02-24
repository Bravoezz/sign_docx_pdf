BINARY_NAME=app.exe
OUTPUT_DOC=output.docx

build:
	go build -o ${BINARY_NAME} main.go

run: cls build
	./${BINARY_NAME}
	echo open file
	cmd /c start ${OUTPUT_DOC}

cls:
	rm ${OUTPUT_DOC}
	rm ${BINARY_NAME}