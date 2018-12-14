NAME   := testq


all: 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o testq .
	docker build -t gherlein/testq .

dependencies:
	go get ./...

clean:
	-rm -f ${NAME}
	-rm -f *~
