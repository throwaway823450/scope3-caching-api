# Run the caching service locally.
.PHONY: run
run:
	. ./.env && go run main.go

test-curl:
	curl --header "Content-Type: application/json"   --request POST   --data '{"rows": [{"inventoryId": "nytimes.com"}]}'   http://localhost:8080/emissions