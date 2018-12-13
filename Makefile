NAME   := testq


all: 
	go build

dependencies:
	go get ./...

clean:
	-rm -f ${NAME}
	-rm -f *~
