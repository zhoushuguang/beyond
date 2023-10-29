package es

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type (
	Config struct {
		Addresses  []string
		Username   string
		Password   string
		MaxRetries int
	}

	Es struct {
		*es8.Client
	}

	// esTransport is a transport for elasticsearch client
	esTransport struct{}
)

func (t *esTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	var (
		ctx        = req.Context()
		span       oteltrace.Span
		startTime  = time.Now()
		propagator = otel.GetTextMapPropagator()
		indexName  = strings.Split(req.URL.RequestURI(), "/")[1]
		tracer     = trace.TracerFromContext(ctx)
	)

	ctx, span = tracer.Start(ctx,
		req.URL.Path,
		oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		oteltrace.WithAttributes(semconv.HTTPClientAttributesFromHTTPRequest(req)...),
	)
	defer func() {
		metricClientReqDur.Observe(time.Since(startTime).Milliseconds(), indexName)
		metricClientReqErrTotal.Inc(indexName, strconv.FormatBool(err != nil))

		span.End()
	}()

	req = req.WithContext(ctx)
	propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err = http.DefaultTransport.RoundTrip(req)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}

	span.SetAttributes(semconv.DBSQLTableKey.String(indexName))
	span.SetAttributes(semconv.HTTPAttributesFromHTTPStatusCode(resp.StatusCode)...)
	span.SetStatus(semconv.SpanStatusFromHTTPStatusCodeAndSpanKind(resp.StatusCode, oteltrace.SpanKindClient))

	return
}

func NewEs(conf *Config) (*Es, error) {
	c := es8.Config{
		Addresses:  conf.Addresses,
		Username:   conf.Username,
		Password:   conf.Password,
		MaxRetries: conf.MaxRetries,
		Transport:  &esTransport{},
	}

	client, err := es8.NewClient(c)
	if err != nil {
		return nil, err
	}

	return &Es{
		Client: client,
	}, nil
}

func MustNewEs(conf *Config) *Es {
	es, err := NewEs(conf)
	if err != nil {
		panic(err)
	}

	return es
}
