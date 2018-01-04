# fibonacci_restful

Fibonacci_restful is a RESTful API implementation for fetching Fibonacci numbers. It accept two kinds of inputs.

	"127.0.0.1:1234/fibonaccisequence/<n>" to get the first "n" fibonacci numbers
	"127.0.0.1:1234/fibonaccinumber/<n>" to get the "nth" fibonacci number

Here are some examples:

	$ curl http://127.0.0.1:1234/fibonaccisequence/5
	[0 1 1 2 3]
	$ curl http://127.0.0.1:1234/fibonaccinumber/5
	5

## Build

fibonacci_restful is implemented by Go language. You can fetch it via `go get` command

	$ go get github.com/czq1285/fibonacci_restful

fibnacci_restful use [httprouter](https://github.com/julienschmidt/httprouter) as the http request router and [logrus](https://github.com/sirupsen/logrus) as the logger. These two pakcages are fetched by default during building the package.

	$ make
	github.com/julienschmidt/httprouter (download)
	github.com/sirupsen/logrus (download)
	github.com/julienschmidt/httprouter
	github.com/sirupsen/logrus
	go build -o ./fibonacci_restful 

Then this project could be started and stopped via the make command:

	$ make run
	$ make stop

## Test

The project provides two kinds of test cases, unit test cases and functinal test cases. These tests cases could be run via the make command:

To run unit tests:

	$ make ut
	go test -cover -v
	=== RUN   TestInitSequence
	--- PASS: TestInitSequence (0.00s)
	=== RUN   TestParameterCheck
	--- PASS: TestParameterCheck (0.00s)
	=== RUN   TestGetFobinacciSequence
	--- PASS: TestGetFobinacciSequence (0.00s)
	=== RUN   TestGetFobinacciNumber
	--- PASS: TestGetFobinacciNumber (0.00s)
	=== RUN   TestHandlers
	--- PASS: TestHandlers (0.00s)
	PASS
	coverage: 95.2% of statements
	ok  	github.com/czq1285/fibonacci_restful	0.021s

To run functional tests: (fibonacci_restful shall be started firstly):

	$ make ft
	go build -o ./function_test 
	./function_test
	Testing GET "/" page ... passed!
	Testing GET "/fibonaccisequence/10" page ... passed!
	Testing GET "/fibonaccisequence/abc" page ... passed!
	Testing GET "/fibonaccisequence/-10" page ... passed!
	Testing GET "/fibonaccinumber/10" page ... passed!
	Testing GET "/fibonaccinumber/abc" page ... passed!
	Testing GET "/fibonaccinumber/-10" page ... passed!
	Testing POST "/" page ... passed!
	Testing POST "/fibonaccisequence/10" page ... passed!
	Testing POST "/fibonaccisequence/abc" page ... passed!
	Testing POST "/fibonaccisequence/-10" page ... passed!
	Testing POST "/fibonaccinumber/10" page ... passed!
	Testing POST "/fibonaccinumber/abc" page ... passed!
	Testing POST "/fibonaccinumber/-10" page ... passed!
	go clean


