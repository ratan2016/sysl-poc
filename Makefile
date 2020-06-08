all: sysl

input = api/project.sysl
app = simple
dependencies = jsonplaceholder # this can be a list separated by a space or left empty
outdir = gen
# Current go import path
basepath = github.com/anz-bank/sysl-template

####################################################################
#                                                                  #
#                                                                  #
#                                                                  #
# START SYSL MAKEFILE: you shouldn't need to edit anything below   #
#                                                                  #
#                                                                  #
#                                                                  #
####################################################################

docker:
	GOOS=linux GOARCH=amd64 go build main.go
	docker build -t joshcarp/sysltemplate .
	docker run -p 8080:8080 joshcarp/sysltemplate


TMP = .tmp# Cache the server lib directory in tmp
SERVERLIB = /var/tmp
TRANSLOCATION = .tmp/sysl-go/codegen/transforms
TRANSFORMS= svc_error_types.sysl svc_handler.sysl svc_interface.sysl svc_router.sysl svc_types.sysl
DOWNSTREAMTRANSFORMS = svc_client.sysl svc_error_types.sysl svc_types.sysl
GRAMMAR=$(wildcard .tmp/sysl-go/codegen/grammars/go.gen.g)
START=goFile


# Always execute these with just `make`
.PHONY: setup clean gen
sysl: clean setup gen downstream format tmp

# try to clone, then try to fetch and pull
setup:
	# Syncing server-lib to $(SERVERLIB)
	git clone https://github.com/anz-bank/sysl-go/ $(SERVERLIB)/sysl-go || true  # Don't fail
	cd  $(SERVERLIB)/sysl-go || true
	mkdir -p $(TMP)/sysl-go/
	mkdir -p ${outdir}/${app}
	# Copying server-lib to $(TMP)
	cp -rf $(SERVERLIB)/sysl-go $(TMP)/
	$(foreach path, $(dependencies), $(shell mkdir -p ${outdir}/$(path)))
    $(foreach path, $(app), $(shell mkdir -p ${outdir}/$(path)))

# Generate files with internal git service
gen:
	$(foreach file, $(TRANSFORMS), $(shell sysl codegen --basepath=$(basepath)/${outdir}/ --transform $(TRANSLOCATION)/$(file) --grammar ${GRAMMAR} --start ${START} --outdir=${outdir}/${app} --app-name ${app} $(input)))

downstream:
	$(foreach file, $(DOWNSTREAMTRANSFORMS), $(foreach downstream, $(dependencies), $(shell sysl codegen --basepath=$(basepath)/${outdir}/ --transform $(TRANSLOCATION)/$(file) --grammar ${GRAMMAR} --start ${START} --outdir=${outdir}/${downstream} --app-name ${downstream} $(input))))

format:
	gofmt -s -w ${outdir}/${app}/*
	goimports -w ${outdir}/${app}/*
	$(foreach path, $(dependencies), $(shell gofmt -s -w ${outdir}/${path}/*))
	$(foreach path, $(dependencies), $(shell goimports -w ${outdir}/${path}/*))
	./replace.sh


# Remove the tmp directory after
tmp:
	rm -rf $(TMP)

# Remove the generated files
clean:
	rm -rf $(outdir)