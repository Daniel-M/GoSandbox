all: variable_types.go variable_types_declaring.go
	$(MAKE) variable_types
	$(MAKE) variable_types_declaring

variable_types: variable_types.go
	go build -o $@ $^

variable_types_declaring: variable_types_declaring.go
	go build -o $@ $^
	
#run: variable_types.go
#go run $^

