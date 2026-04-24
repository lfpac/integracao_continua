lint:
	MSYS_NO_PATHCONV=1 MSYS2_ARG_CONV_EXCL="*" \
	docker run --rm -it \
  		-v "$$(cygpath -m "$$PWD"):/app" \
  		-w /app \
  		golangci/golangci-lint \
  		golangci-lint run controllers/ database/ models/ routes/

test: 
	docker compose exec app go test main_test.go

start:
	docker compose up -d

ci: 
	start lint test