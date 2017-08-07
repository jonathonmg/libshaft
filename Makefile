BIN_DIR=bin


default: all

all: libshaft

libshaft: clean
	mkdir -p $(BIN_DIR)
	godep go build -o $(BIN_DIR)/libshaft apps/*.go
	cp bin/libshaft ${GOPATH}/bin/libshaft
	cp -Rf resources ${GOPATH}/bin/.

clean:
	rm -rf $(BIN_DIR)/*