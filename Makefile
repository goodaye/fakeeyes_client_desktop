GOCC=go
#GC_FLAGS= 

TARGET = bin/fakeeyesclient
M_PREFIX = 
M_SRC = *.go
# PROJECT = skyflow

MAIN_PKG = cmd 

#build target
all :  prepare  $(TARGET) 

prepare:
	@echo "prepare"
	@echo $(GOPATH)
	@mkdir -p bin/

.PHONY: $(TARGET)
$(TARGET): $(M_SRC)
	@echo "build target"
	@echo $(M_SRC)
	$(GOCC) build -o $@ $^


clean:
	@echo $(GOPATH)
	-@rm -rf $(TARGET)