all:tidy	run

run:
	go run ./app/sales-api
tidy:
	go mod tidy
vendor:
	go mod vendor

# hey -m GET -c 100 -n 10000000 http://localhost:3000/readiness/10
# expvarmon -ports="localhost:4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"
