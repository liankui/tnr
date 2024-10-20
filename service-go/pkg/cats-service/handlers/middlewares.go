package handlers

import (
	"github.com/chaos-io/gokit/middleware"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/liankui/tnr/go/pkg/cat"
	"github.com/liankui/tnr/service-go/pkg/cats-service/svc"
	"github.com/mojo-lang/core/go/pkg/mojo/core"

	// this service api
	pb "github.com/liankui/tnr/go/pkg/cat/v1"
)

var (
	_ = cat.Cat{}
	_ = core.Null{}
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints, options map[string]interface{}) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/ncraft-io//_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	var tracer stdopentracing.Tracer
	if value, ok := options["tracer"]; ok && value != nil {
		tracer = value.(stdopentracing.Tracer)
	}
	var count *kitprometheus.Counter
	if value, ok := options["count"]; ok && value != nil {
		count = value.(*kitprometheus.Counter)
	}
	var latency *kitprometheus.Histogram
	if value, ok := options["latency"]; ok && value != nil {
		latency = value.(*kitprometheus.Histogram)
	}
	//var validator *middleware.Validator
	//if value, ok := options["validator"]; ok && value != nil {
	//	validator = value.(*middleware.Validator)
	//}

	{ // create_cat
		if tracer != nil {
			in.CreateCatEndpoint = opentracing.TraceServer(tracer, "create_cat")(in.CreateCatEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateCatEndpoint = middleware.Instrumenting(latency.With("method", "create_cat"), count.With("method", "create_cat"))(in.CreateCatEndpoint)
		}
		//if validator != nil {
		//	in.CreateCatEndpoint = validator.Validate()(in.CreateCatEndpoint)
		//}
	}
	{ // update_cat
		if tracer != nil {
			in.UpdateCatEndpoint = opentracing.TraceServer(tracer, "update_cat")(in.UpdateCatEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateCatEndpoint = middleware.Instrumenting(latency.With("method", "update_cat"), count.With("method", "update_cat"))(in.UpdateCatEndpoint)
		}
		//if validator != nil {
		//	in.UpdateCatEndpoint = validator.Validate()(in.UpdateCatEndpoint)
		//}
	}
	{ // get_cat
		if tracer != nil {
			in.GetCatEndpoint = opentracing.TraceServer(tracer, "get_cat")(in.GetCatEndpoint)
		}
		if count != nil && latency != nil {
			in.GetCatEndpoint = middleware.Instrumenting(latency.With("method", "get_cat"), count.With("method", "get_cat"))(in.GetCatEndpoint)
		}
		//if validator != nil {
		//	in.GetCatEndpoint = validator.Validate()(in.GetCatEndpoint)
		//}
	}
	{ // list_cats
		if tracer != nil {
			in.ListCatsEndpoint = opentracing.TraceServer(tracer, "list_cats")(in.ListCatsEndpoint)
		}
		if count != nil && latency != nil {
			in.ListCatsEndpoint = middleware.Instrumenting(latency.With("method", "list_cats"), count.With("method", "list_cats"))(in.ListCatsEndpoint)
		}
		//if validator != nil {
		//	in.ListCatsEndpoint = validator.Validate()(in.ListCatsEndpoint)
		//}
	}
	{ // batch_get_cat
		if tracer != nil {
			in.BatchGetCatEndpoint = opentracing.TraceServer(tracer, "batch_get_cat")(in.BatchGetCatEndpoint)
		}
		if count != nil && latency != nil {
			in.BatchGetCatEndpoint = middleware.Instrumenting(latency.With("method", "batch_get_cat"), count.With("method", "batch_get_cat"))(in.BatchGetCatEndpoint)
		}
		//if validator != nil {
		//	in.BatchGetCatEndpoint = validator.Validate()(in.BatchGetCatEndpoint)
		//}
	}
	{ // delete_cat
		if tracer != nil {
			in.DeleteCatEndpoint = opentracing.TraceServer(tracer, "delete_cat")(in.DeleteCatEndpoint)
		}
		if count != nil && latency != nil {
			in.DeleteCatEndpoint = middleware.Instrumenting(latency.With("method", "delete_cat"), count.With("method", "delete_cat"))(in.DeleteCatEndpoint)
		}
		//if validator != nil {
		//	in.DeleteCatEndpoint = validator.Validate()(in.DeleteCatEndpoint)
		//}
	}

	return in
}

func WrapService(in pb.CatsServer, options map[string]interface{}) pb.CatsServer {
	return in
}
