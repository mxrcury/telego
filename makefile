BINARY_NAME="main"

run:
	go build -o ${BINARY_NAME}
	./${BINARY_NAME}
	rm ${BINARY_NAME}
	go clean
