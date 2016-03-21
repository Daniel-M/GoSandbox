all: variable_types.go input_to_variable.go
	$(MAKE) variable_types
	$(MAKE) input_to_variable

variable_types: variable_types.go
	go build -o $@ $^

input_to_variable: input_to_variable.go
	go build -o $@ $^
