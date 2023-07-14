BINARY_NAME=main
 
build:
	go build -o ./review/cmd/coupon_service/${BINARY_NAME} ./review/cmd

deploy-dev:
	docker build -t coupon_service ./review/cmd/coupon_service/${BINARY_NAME}
	docker run -d -p 8080:8080 coupon_service
 
run:
	go build -o ./review/cmd/coupon_service/${BINARY_NAME} ./review/cmd/
	./review/cmd/coupon_service/${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}



