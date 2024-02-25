run:
	# in case you need to re-build it, uncomment the following line
	# docker compose stop && docker compose down && docker compose build && docker compose up origin srt app
	docker compose stop && docker compose build app && docker compose up origin srt app

test:
	# in case you need to re-build it, uncomment the following line
	# docker compose stop test && docker compose down test && docker compose build test && docker compose run --rm test
	docker compose stop test && docker compose down test && docker compose run --rm test

run-srt:
	docker compose stop && docker compose down && docker compose build srt && docker compose up srt

mac-run-local:
	./scripts/mac_local_run.sh

mac-test-local:
	./scripts/mac_local_run_test.sh

html-local-coverage:
	go tool cover -html=coverage.out

lint:
	docker compose stop lint && docker compose down lint && docker compose run --rm lint	

# INCOMPLETE from https://github.com/asticode/go-astiav/blob/master/Makefile
install-ffmpeg:
	./scripts/install_local_ffmpeg.sh

.PHONY: run lint test run-srt mac-run-local mac-test-local html-local-coverage install-ffmpeg
