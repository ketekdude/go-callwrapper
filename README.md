# go-callwrapper

`go-callwrapper` is a Go library designed to wrap external calls to databases, Redis, services, or any downstream system with built-in reliability and performance features.

## Features

This wrapper helps you build resilient and efficient external call logic by providing:

1. **Circuit Breaker**  
   Automatically stops calls to unstable downstream services to prevent cascading failures.

2. **Call With Timeout**  
   Ensures that calls do not block indefinitely by enforcing a maximum wait time, protecting your system from high latency spikes.

3. **Memcache Integration**  
   Caches data that changes infrequently to reduce unnecessary calls to downstream services, improving response time and reducing load.

## Why use go-callwrapper?

When your application depends on external systems like databases, caches, or remote services, issues like latency spikes, failures, or overload can impact your entire system.  
`go-callwrapper` offers a simple way to add robustness by:

- Detecting failures and temporarily halting requests to failing services (circuit breaker).  
- Timing out slow calls to avoid blocking your app (timeout).  
- Leveraging caching to minimize redundant calls for relatively static data (memcache).

## Typical Usage

Wrap your external calls (e.g., DB queries, Redis commands, service requests) inside `go-callwrapper` to gain resilience and caching benefits transparently.

```go
response, err := callWrapper.Execute(func(ctx context.Context) (interface{}, error) {
    // Your logic here: memcache check → backend call → memcache set
})
