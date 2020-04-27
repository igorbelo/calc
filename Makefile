all:
	rm -rf ./bin
	mkdir bin
	go run main.go > ./bin/calc.ll
	llc -filetype=obj ./bin/calc.ll
	clang ./bin/calc.o main.c -o ./bin/calc
	make clean

clean:
	rm ./bin/calc.ll
	rm ./bin/calc.o

