.PHONY: help
help: ## Display this help.
  @awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: wasm
wasm: ## build the game part (WASM)
	cd game && tinygo build -no-debug -o ../backend/assets/game.wasm -target wasm ./...

.PHONY: backend
backend: ## build the backend part 
	cd backend && go build -o ../bin/conway ./main.go

.PHONY: app
app: wasm backend ## build the entire application

.PHONY: run
run: app ## build and run the entire application
	./bin/conway

