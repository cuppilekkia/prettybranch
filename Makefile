PROJECTNAME=ptb

.PHONY: build clean link
build: 
	go build -o bin/$(PROJECTNAME) main.go

clean: ## Remove build related file
	rm -fr ./bin

install: ## Copy the binary into $HOME/bin
	mkdir -p $(HOME)/bin
	cp ./bin/$(PROJECTNAME) $(HOME)/bin/