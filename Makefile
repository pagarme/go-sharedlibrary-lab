
all: tester

clean:
	rm -rf tmp out

tmp/libtms.a: libc.go tms.go
	@mkdir -p tmp/libtms
	go build -buildmode=c-archive -o tmp/libtms.a libc.go tms.go
	tar xf tmp/libtms.a -C tmp/libtms/

out/libtms.dylib: tmp/libtms.a
	@mkdir out
	clang -dynamiclib -o out/libtms.dylib tmp/libtms/*.o
	
tester: out/libtms.dylib
	clang -Lout -ltms -o tester test.c

