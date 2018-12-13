NAME   := testq


all: 
	go build
	docker build -t gherlein/testq .

dependencies:
	go get ./...

clean:
	-rm -f ${NAME}
	-rm -f *~
