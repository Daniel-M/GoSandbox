all: arraysSlices.go maps.go
	$(MAKE) arraysSlices
	$(MAKE) maps

arraysSlices: arraysSlices.go
	go build -o $@ $^

maps: maps.go
	go build -o $@ $^
