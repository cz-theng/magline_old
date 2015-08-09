TARGET := maglined

SRC_FILES := $(shell find .. ../agent ../proto ../server -name "*.go" )
SRC_FILES += main.go flag.go config.go

.DEFAULT:all

all : $(TARGET)
	@echo "[SUCCESS] Done go build"


$(TARGET): $(SRC_FILES)
	@echo "src :"$(SRC_FILES)
	go build  -v -o $(TARGET)

fmt : $(SRC_FILES)
	go fmt 
	@echo "[SUCCESS] Done go fmt"

.PHONY:clean

clean:
	go clean
	-rm -rf $(TARGET)
	-rm -rf log
