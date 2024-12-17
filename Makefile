# Run the caching service locally.
.PHONY: run
run:
	. ./.env && go run main.go