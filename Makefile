#COLORS
GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

# Add the following 'help' target to your Makefile
# And add help text after each target name starting with '\#\#'
# A category can be added with @category
HELP_FUN = \
    %Targets; \
    while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
    print "\n${WHITE}usage:\n"; \
	print "  ${YELLOW}make <target>                  ${GREEN} replace the <target> with one of below operations.\n\n"; \
    for (sort keys %help) { \
    print "${WHITE}$$_:${RESET}\n"; \
    for (@{$$help{$$_}}) { \
    $$sep = " " x (32 - length $$_->[0]); \
    print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
    }; \
    print "\n"; }

help: ##@others show target help options.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

start: FORCE ##@app start the http app server at http://localhost:5000
	go run main.go

bdd-test: FORCE ##@test executes godog BDD tests
	godog

bdd-test-report: FORCE ##@test executes godog BDD tests and generates a json and html report (but requires node and cucumber-html-reporter package)
	godog --format=cucumber > cucumber_report.json;\
    node index.js

test: FORCE ##@test executes unit tests
	go test ./...

mutation-test: FORCE ##@test executes go-mutesting in ./api/... dir
	rm -rf go-mutesting-error-report.html go-mutesting-ok-report.html output.txt ;\
    echo "<h1 style="text-align:center" > Mutation Test Error Report </h1><hr>" >> go-mutesting-error-report.html ;\
    echo "<h1 style="text-align:center" > Mutation Test Success OK Report </h1><hr>" >> go-mutesting-ok-report.html ;\
    go-mutesting --exec "./test-mutated-package-custom.sh" --test-recursive api/service/... --debug  > output.txt
    
FORCE:
