package orm

import (
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"gorm.io/gorm"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type CustomePlugin struct{}

func NewCustomePlugin() *CustomePlugin {
	return &CustomePlugin{}
}

func (p *CustomePlugin) Name() string {
	return "CustomePlugin"
}

func (p *CustomePlugin) Initialize(db *gorm.DB) error {
	// Before callbacks
	if err := db.Callback().Create().Before("gorm:createBefore").Register("gorm:createBefore:metric:trace", func(db *gorm.DB) {
		startTime := time.Now().Unix()
		db.InstanceSet("gorm:create_start_time", startTime)

		ctx := db.Statement.Context
		tracer := trace.TracerFromContext(ctx)
		_, span := tracer.Start(ctx,
			"gorm:create",
			oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		)
		db.InstanceSet("gorm:create_span", span)
	}); err != nil {
		return err
	}
	if err := db.Callback().Query().Before("gorm:queryBefore").Register("gorm:queryBefore:metric:trace", func(db *gorm.DB) {
		startTime := time.Now().Unix()
		db.InstanceSet("gorm:query_start_time", startTime)

		ctx := db.Statement.Context
		tracer := trace.TracerFromContext(ctx)
		_, span := tracer.Start(ctx,
			"gorm:query",
			oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		)
		db.InstanceSet("gorm:query_span", span)
	}); err != nil {
		return err
	}
	if err := db.Callback().Update().Before("gorm:updateBefore").Register("gorm:updateBefore:metric:trace", func(db *gorm.DB) {
		startTime := time.Now().Unix()
		db.InstanceSet("gorm:update_start_time", startTime)

		ctx := db.Statement.Context
		tracer := trace.TracerFromContext(ctx)
		_, span := tracer.Start(ctx,
			"gorm:update",
			oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		)
		db.InstanceSet("gorm:update_span", span)
	}); err != nil {
		return err
	}
	if err := db.Callback().Delete().Before("gorm:deleteBefore").Register("gorm:deleteBefore:metric:trace", func(db *gorm.DB) {
		startTime := time.Now().Unix()
		db.InstanceSet("gorm:delete_start_time", startTime)

		ctx := db.Statement.Context
		tracer := trace.TracerFromContext(ctx)
		_, span := tracer.Start(ctx,
			"gorm:delete",
			oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		)
		db.InstanceSet("gorm:delete_span", span)
	}); err != nil {
		return err
	}
	if err := db.Callback().Row().Before("gorm:rowBefore").Register("gorm:rowBefore:metric:trace", func(db *gorm.DB) {
		startTime := time.Now().Unix()
		db.InstanceSet("gorm:row_start_time", startTime)

		ctx := db.Statement.Context
		tracer := trace.TracerFromContext(ctx)
		_, span := tracer.Start(ctx,
			"gorm:row",
			oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		)
		db.InstanceSet("gorm:row_span", span)
	}); err != nil {
		return err
	}
	if err := db.Callback().Raw().Before("gorm:rawBefore").Register("gorm:rawBefore:metric:trace", func(db *gorm.DB) {
		startTime := time.Now().Unix()
		db.InstanceSet("gorm:raw_start_time", startTime)

		ctx := db.Statement.Context
		tracer := trace.TracerFromContext(ctx)
		_, span := tracer.Start(ctx,
			"gorm:raw",
			oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		)
		db.InstanceSet("gorm:raw_span", span)
	}); err != nil {
		return err
	}

	// After callbacks
	if err := db.Callback().Create().After("gorm:createAfter").Register("gorm:createAfter:metric:trace", func(db *gorm.DB) {
		startTime, ok := db.InstanceGet("gorm:create_start_time")
		if !ok {
			return
		}
		stSecond := startTime.(int64)
		st := time.Unix(stSecond, 0)
		metricClientReqDur.Observe(time.Since(st).Milliseconds(), db.Statement.Table, "create")
		metricClientReqErrTotal.Inc(db.Statement.Table, "create", strconv.FormatBool(db.Statement.Error != nil))

		v, ok := db.InstanceGet("gorm:create_span")
		if !ok {
			return
		}
		span := v.(oteltrace.Span)
		if db.Statement.Error != nil {
			span.RecordError(db.Statement.Error)
		}
		span.SetAttributes(
			semconv.DBSQLTableKey.String(db.Statement.Table),
			semconv.DBStatementKey.String(db.Statement.SQL.String()),
		)
		span.End()
	}); err != nil {
		return err
	}
	if err := db.Callback().Query().After("gorm:queryAfter").Register("gorm:queryAfter:metric:trace", func(db *gorm.DB) {
		startTime, ok := db.InstanceGet("gorm:query_start_time")
		if !ok {
			return
		}
		stSecond := startTime.(int64)
		st := time.Unix(stSecond, 0)
		metricClientReqDur.Observe(time.Since(st).Milliseconds(), db.Statement.Table, "query")
		metricClientReqErrTotal.Inc(db.Statement.Table, "query", strconv.FormatBool(db.Statement.Error != nil))

		v, ok := db.InstanceGet("gorm:query_span")
		if !ok {
			return
		}
		span := v.(oteltrace.Span)
		if db.Statement.Error != nil {
			span.RecordError(db.Statement.Error)
		}
		span.SetAttributes(
			semconv.DBSQLTableKey.String(db.Statement.Table),
			semconv.DBStatementKey.String(db.Statement.SQL.String()),
		)
		span.End()
	}); err != nil {
		return err
	}
	if err := db.Callback().Update().After("gorm:updateAfter").Register("gorm:updateAfter:metric:trace", func(db *gorm.DB) {
		startTime, ok := db.InstanceGet("gorm:update_start_time")
		if !ok {
			return
		}
		stSecond := startTime.(int64)
		st := time.Unix(stSecond, 0)
		metricClientReqDur.Observe(time.Since(st).Milliseconds(), db.Statement.Table, "update")
		metricClientReqErrTotal.Inc(db.Statement.Table, "update", strconv.FormatBool(db.Statement.Error != nil))

		v, ok := db.InstanceGet("gorm:update_span")
		if !ok {
			return
		}
		span := v.(oteltrace.Span)
		if db.Statement.Error != nil {
			span.RecordError(db.Statement.Error)
		}
		span.SetAttributes(
			semconv.DBSQLTableKey.String(db.Statement.Table),
			semconv.DBStatementKey.String(db.Statement.SQL.String()),
		)
		span.End()
	}); err != nil {
		return err
	}
	if err := db.Callback().Delete().After("gorm:deleteAfter").Register("gorm:deleteAfter:metric:trace", func(db *gorm.DB) {
		startTime, ok := db.InstanceGet("gorm:delete_start_time")
		if !ok {
			return
		}
		stSecond := startTime.(int64)
		st := time.Unix(stSecond, 0)
		metricClientReqDur.Observe(time.Since(st).Milliseconds(), db.Statement.Table, "delete")
		metricClientReqErrTotal.Inc(db.Statement.Table, "delete", strconv.FormatBool(db.Statement.Error != nil))

		v, ok := db.InstanceGet("gorm:delete_span")
		if !ok {
			return
		}
		span := v.(oteltrace.Span)
		if db.Statement.Error != nil {
			span.RecordError(db.Statement.Error)
		}
		span.SetAttributes(
			semconv.DBSQLTableKey.String(db.Statement.Table),
			semconv.DBStatementKey.String(db.Statement.SQL.String()),
		)
		span.End()
	}); err != nil {
		return err
	}
	if err := db.Callback().Row().After("gorm:rowAfter").Register("gorm:rowAfter:metric:trace", func(db *gorm.DB) {
		startTime, ok := db.InstanceGet("gorm:row_start_time")
		if !ok {
			return
		}
		stSecond := startTime.(int64)
		st := time.Unix(stSecond, 0)
		metricClientReqDur.Observe(time.Since(st).Milliseconds(), db.Statement.Table, "row")
		metricClientReqErrTotal.Inc(db.Statement.Table, "row", strconv.FormatBool(db.Statement.Error != nil))

		v, ok := db.InstanceGet("gorm:row_span")
		if !ok {
			return
		}
		span := v.(oteltrace.Span)
		if db.Statement.Error != nil {
			span.RecordError(db.Statement.Error)
		}
		span.SetAttributes(
			semconv.DBSQLTableKey.String(db.Statement.Table),
			semconv.DBStatementKey.String(db.Statement.SQL.String()),
		)
		span.End()
	}); err != nil {
		return err
	}
	if err := db.Callback().Raw().After("gorm:rawAfter").Register("gorm:rawAfter:metric:trace", func(db *gorm.DB) {
		startTime, ok := db.InstanceGet("gorm:raw_start_time")
		if !ok {
			return
		}
		stSecond := startTime.(int64)
		st := time.Unix(stSecond, 0)
		metricClientReqDur.Observe(time.Since(st).Milliseconds(), db.Statement.Table, "raw")
		metricClientReqErrTotal.Inc(db.Statement.Table, "raw", strconv.FormatBool(db.Statement.Error != nil))

		v, ok := db.InstanceGet("gorm:raw_span")
		if !ok {
			return
		}
		span := v.(oteltrace.Span)
		if db.Statement.Error != nil {
			span.RecordError(db.Statement.Error)
		}
		span.SetAttributes(
			semconv.DBSQLTableKey.String(db.Statement.Table),
			semconv.DBStatementKey.String(db.Statement.SQL.String()),
		)
		span.End()
	}); err != nil {
		return err
	}

	return nil
}
