# Run the caching service locally.
.PHONY: run
run:
	. ./.env && go run main.go

.PHONY: test-curl
test-curl:
	curl --header "Content-Type: application/json" --request POST -d @example_request.json http://localhost:8080/emissions

.PHONY: test-curl-routine
test-curl-routine:
	$(MAKE) test-curl
	sleep 5
	$(MAKE) test-curl
	sleep 15
	$(MAKE) test-curl