# Observer

`observer` is a lightweight **observability toolkit for Go** that combines:

- **Structured Logging** (pluggable backends, e.g. Zap)
- **Distributed Tracing** (pluggable backends, e.g. OpenTelemetry via Prospan)
- **Context Propagation** (logs and traces follow your `context.Context`)

It is designed to make **logs, metrics, and traces** seamlessly work together with minimal setup.

---

## Why Observer?

Modern applications need **correlation** between logs and traces. With `observer`:

- Logs are automatically enriched with **trace IDs** (great for Loki/Grafana).
- Traces are easy to create and propagate through functions.
- You can swap logging/tracing backends without changing application code.
- Both logging and tracing providers support **graceful shutdown**.

---

## Features

- ✅ **Pluggable Logger** – default: [Zap](https://github.com/uber-go/zap)
- ✅ **Pluggable Tracer** – default: [OpenTelemetry](https://opentelemetry.io/) via Prospan
- ✅ **Trace-aware Logging** – `span.Logger()` includes `trace_id` in logs
- ✅ **Context-first API** – no global state pollution, everything flows through `context.Context`
- ✅ **Graceful Lifecycle** – start/stop loggers and tracers cleanly

---

## Installation

```bash
go get github.com/not-for-prod/observer
