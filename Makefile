.PHONY: all test bench bench-all run clean fmt help day-run

DAYS := 01 02 03 04 05 06 07 08 09 10 11 12

# Default day (override with: make test DAY=02)
DAY ?= 01

help:
	@echo "Advent of Code 2025 - Go"
	@echo ""
	@echo "Usage:"
	@echo "  make star1 DAY=01       Run star1 (example + input) for a specific day"
	@echo "  make star2 DAY=01       Run star2 (example + input) for a specific day"
	@echo "  make fmt                Format all Go files"
	@echo "  make clean              Clean build cache"
	@echo "  make 01                 Run both stars (example + input) for Day 01"
	@echo ""

# Run only star1 for a specific day
star1:
	@go run -exec "" ./cmd/run $(DAY) 1 example
	@go run -exec "" ./cmd/run $(DAY) 1

# Run only star2 for a specific day
star2:
	@go run -exec "" ./cmd/run $(DAY) 2 example
	@go run -exec "" ./cmd/run $(DAY) 2

.PHONY: $(DAYS)
$(DAYS):
	@$(MAKE) --no-print-directory DAY=$@ day-run

day-run:
	@echo "== Day $(DAY) =="
	@$(MAKE) --no-print-directory DAY=$(DAY) star1
	@$(MAKE) --no-print-directory DAY=$(DAY) star2

# Format all Go files
fmt:
	go fmt ./...

# Clean build cache
clean:
	go clean -cache -testcache
