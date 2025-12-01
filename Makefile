.PHONY: all test bench bench-all run clean fmt help

# Default day (override with: make test DAY=02)
DAY ?= 01

help:
	@echo "Advent of Code 2025 - Go"
	@echo ""
	@echo "Usage:"
	@echo "  make test DAY=01        Run tests for a specific day"
	@echo "  make test-all           Run all tests"
	@echo "  make bench DAY=01       Benchmark a specific day"
	@echo "  make bench-all          Benchmark all days"
	@echo "  make star1 DAY=01       Run star1 test for a specific day"
	@echo "  make star2 DAY=01       Run star2 test for a specific day"
	@echo "  make fmt                Format all Go files"
	@echo "  make clean              Clean build cache"
	@echo ""

# Run tests for a specific day
test:
	go test -v ./internal/days/day$(DAY)/...

# Run all tests
test-all:
	go test -v ./internal/days/...

# Benchmark a specific day
bench:
	go test -bench=. -benchmem ./internal/days/day$(DAY)/...

# Benchmark all days
bench-all:
	go test -bench=. -benchmem ./internal/days/...

# Run only star1 for a specific day
star1:
	@go run -exec "" ./cmd/run $(DAY) 1

# Run only star2 for a specific day
star2:
	@go run -exec "" ./cmd/run $(DAY) 2

# Run star1 with example input
star1-example:
	go test -v -run TestStar1Example ./internal/days/day$(DAY)/...

# Run star2 with example input
star2-example:
	go test -v -run TestStar2Example ./internal/days/day$(DAY)/...

# Benchmark only star1 for a specific day
bench-star1:
	go test -bench=BenchmarkStar1 -benchmem ./internal/days/day$(DAY)/...

# Benchmark only star2 for a specific day
bench-star2:
	go test -bench=BenchmarkStar2 -benchmem ./internal/days/day$(DAY)/...

# Format all Go files
fmt:
	go fmt ./...

# Clean build cache
clean:
	go clean -cache -testcache
