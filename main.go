package main

import (
	servicea "github.com/lucas-code42/OTEL-impl-example/services/serviceA"
	serviceb "github.com/lucas-code42/OTEL-impl-example/services/serviceB"
)

func main() {
	infinity := make(chan bool)
	go servicea.RunServiceA()
	go serviceb.RunServiceB()
	<-infinity
}
