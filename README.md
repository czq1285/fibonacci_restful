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

fibnacci_restful use [httprouter](https://github.com/julienschmidt/httprouter) as the http request router and [logrus](https://github.com/sirupsen/logrus) as the logger. To build the project, these two pakcages shall be fetched as well.

	$ go get github.com/julienschmidt/httprouter
	$ go get github.com/sirupsen/logrus

Then this project could be built and run:

	$ cd github.com/czq1285/fibonacci_restful
	$ go build
	$ ./fibonacci_restful

## Test

The project provides two kinds of test cases, unit test cases and functinal test cases.

To run unit tests:

	$ go test
	PASS
	ok  	github.com/czq1285/fibonacci_restful	0.020s

To run functional tests: (fibonacci_restful shall be started firstly):

	$ cd function_test/
	$ go run fibonacci_restful_suite.go 
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


