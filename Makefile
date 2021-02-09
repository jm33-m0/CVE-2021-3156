all:
	rm -rf libnss_X
	mkdir libnss_X
	go build -o exploit exploit.go
	gcc -fPIC -shared -o 'libnss_X/P0P_SH3LLZ_ .so.2' lib.c
clean:
	rm -rf libnss_X exploit
