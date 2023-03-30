all: app

cparser.c: cparser.l
	flex -o cparser.c cparser.l
cparser.tab.c: cparser.y
	bison -t -d -o cparser.tab.c cparser.y

cparser.tab.h: cparser.tab.c
	:

parser: cparser.c cparser.tab.c cparser.tab.h
	gcc -o parser cparser.c cparser.tab.c -lfl

libcparser.so: cparser.c cparser.tab.c cparser.tab.h
	gcc -shared -fPIC -o libcparser.so cparser.c cparser.tab.c

app: cparser.go libcparser.so
	go build

clean:
	rm -f cparser.tab.h cparser.tab.c cparser.c libcparser.so parser
