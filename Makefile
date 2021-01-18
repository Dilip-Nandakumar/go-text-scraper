include ./makeUtils.mk

compile:
	$(call build.compile_local)

clean:
	$(call build.clean)

test:
	$(call test.run)

check-fmt:
	$(call fmt.check)

fix-fmt:
	$(call fmt.fix)

build: clean check-fmt compile test

build-docker:
	$(call build.compile_linux)
	$(call docker.build_image)

run-docker: build-docker
	$(call docker.run)
