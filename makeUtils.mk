name := word-highlights-scraper
build_out := out/build
bin := $(build_out)/$(name)
linux_bin := $(build_out)/$(name)-linux
main_dir := ./
main_dir_files := ./...

define build.compile_local
	go build -o $(bin) $(main_dir)
endef

define build.compile_linux
	GOOS=linux go build -o $(linux_bin) $(main_dir)
endef

define build.clean
	go clean -cache -testcache $(main_dir)
endef

define test.run
	go test -coverpkg=$(main_dir_files) -coverprofile c.out $(main_dir_files)
	go tool cover -func c.out
endef

define fmt.check
	gofmt -d -l -e $(main_dir)
endef

define fmt.fix
	gofmt -d -l -w $(main_dir)
endef

define docker.build_image
	docker build . -t $(name)
endef
