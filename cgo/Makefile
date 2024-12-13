
run: mylib.a
	go run cgo.go

build: mylib.a
	go build -o cgo cgo.go

mylib.a: mylib.o
	ar rv $@ $<

%.o:%.c
	cc -g -fPIC -std=c17 -c $<

clean:
	rm -f *.a *.o cgo
