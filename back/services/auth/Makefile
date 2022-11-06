unit-test:
	go test -v $$(go list ./... | grep -v /integrationtests)
	
integration-test:
	docker-compose -f ./integrationtests/docker-compose.yml up -d 
	go test -v ./integrationtests
	docker-compose -f ./integrationtests/docker-compose.yml down

