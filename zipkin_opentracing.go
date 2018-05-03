package main

import (
	"fmt"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

func main() {
	// 1) Create a opentracing.Tracer that sends data to Zipkin
	server := fmt.Sprintf("http://%s:9411/api/v1/spans", os.Args[1])
	fmt.Println(server)
	collector, err := zipkin.NewHTTPCollector(server)
	if err != nil {
		fmt.Errorf("err:%v", err)
	}
	tracer, err := zipkin.NewTracer(
		zipkin.NewRecorder(collector, false, "127.0.0.1:0", "abcd"))

	if err != nil {
		fmt.Errorf("err:%v", err)
	}

	// 2) Demonstrate simple OpenTracing instrumentation
	parent := tracer.StartSpan("Parent")
	for i := 0; i < 20; i++ {
		parent.LogEvent(fmt.Sprintf("Starting child #%d", i))
		child := tracer.StartSpan("Child", opentracing.ChildOf(parent.Context()))
		time.Sleep(50 * time.Millisecond)
		child.Finish()
	}
	parent.LogEvent("A Log")
	parent.Finish()

	// ... give Zipkin ample time to flush
	time.Sleep(2 * time.Second)
}