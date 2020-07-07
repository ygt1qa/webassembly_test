build:
	GOOS=js GOARCH=wasm go build -o main.wasm
start:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'	