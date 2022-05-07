test:
	go test -v -coverprofile cover.out ./repositories/
	go tool cover -html=cover.out -o cover.html
	open cover.html