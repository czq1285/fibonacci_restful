include ./build.mk

#add dependencies
DEPS:=github.com/julienschmidt/httprouter github.com/sirupsen/logrus

#start the web service
run: build
	./${BINARY} &

#stop the web service
stop:
	pkill ${BINARY}

#run function test
ft:
	@${MAKE} -C function_test run
	@${MAKE} -C function_test clean

.PHONY:run stop ft
