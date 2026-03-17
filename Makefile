.PHONY: run
run:
	go run $(MAIN_PACKAGE_PATH)main.go


.PHONY: test
test: # run unit test
	go test ./fizzbuzz