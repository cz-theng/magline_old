SRC_FILES := $(shell find .  -name "*.go" )

.DEFAULT:all

all : $(SRC_FILES)
	@echo "[SUCCESS] Done go build"
	go build  -v 

fmt : $(SRC_FILES)
	go fmt 
	@echo "[SUCCESS] Done go fmt"

.PHONY:clean

clean:
	go clean
