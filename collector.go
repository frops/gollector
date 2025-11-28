package main

// Collector defines a generic interface for collecting data of type T from a specific source.
// It provides type safety and consistent error handling across different data sources.
//
// Type parameter T represents the data type returned by this collector.
type Collector[T any] interface {
	// Collect retrieves data from the source.
	// Returns the collected data or an error if collection fails.
	Collect(ctx context.Context) (T, error)
}
