BINARY_NAME=main
 
build:
	go build -o ./review/cmd/coupon_service/${BINARY_NAME} ./review/cmd

deploy-dev:
	# here I would put a search for the previous container and delete it and then...
	docker build -t coupon_service ./review/
	docker run -d -p 8080:8080 coupon_service
 
run:
	go build -o ./review/cmd/coupon_service/${BINARY_NAME} ./review/cmd/
	./review/cmd/coupon_service/${BINARY_NAME}
 
clean:
	go clean
	rm ./review/cmd/coupon_service/${BINARY_NAME}



