
run123: mylib.a
	go run cgo-receiver-go123.go

run124: mylib.a
	go run cgo-receiver-go124.go

mylib.a: mylib.o
	ar rv $@ $<

%.o:%.c
	cc -g -fPIC -std=c17 -c $<

clean:
	rm -f *.a *.o cgo
