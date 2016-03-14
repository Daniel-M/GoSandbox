all: variable_types.go variable_types_declaring.go
	$(MAKE) variable_types
	$(MAKE) variable_types_declaring
	$(MAKE) input_to_variable

variable_types: variable_types.go
	go build -o $@ $^

variable_types_declaring: variable_types_declaring.go
	go build -o $@ $^
	

input_to_variable: input_to_variable.go
	go build -o $@ $^
