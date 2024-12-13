
tidy-main-module-A:
	cd main-module-A; go mod tidy

run-main-module-A:
	cd main-module-A; go run .

build-main-module-A:
	cd main-module-A; go build -o main .

tidy-main-module-B:
	cd main-module-B; go mod tidy

run-main-module-B:
	cd main-module-B; go run .

build-main-module-B:
	cd main-module-B; go build -o main .

.PHONY: submodule-A
submodule-A:
	cd submodule-A; go build -o sub .

.PHONY: submodule-B
submodule-B:
	cd submodule-B; go build -o sub .

clean:
	rm -f */main submodule-*/sub
