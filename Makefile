

# Include
include scripts/make-rules/tools.mk


## help: show this help message
.PHONY: help
help: Makefile
	@printf "\nUsage: make <TARGET> <OPTIONS> ...\n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"