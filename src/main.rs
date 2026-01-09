use std::env;
use std::time::Instant;

mod day05;

type SolverFn = fn(&str) -> String;

struct Day {
    number: u8,
    star1: SolverFn,
    star2: SolverFn,
    input: &'static str,
}

fn get_days() -> Vec<Day> {
    vec![
    ]
}

/// Returns the total time for this day in milliseconds
fn run_day(day: &Day, bench_iterations: Option<u32>) -> f64 {
    println!("═══════════════════════════════════════");
    println!("  Day {:02}", day.number);
    println!("═══════════════════════════════════════");

    let day_time_ms = if let Some(iterations) = bench_iterations {
        // Benchmark mode
        let start = Instant::now();
        let mut result1 = String::new();
        for _ in 0..iterations {
            result1 = (day.star1)(day.input);
        }
        let elapsed1 = start.elapsed();
        let avg1_ms = elapsed1.as_secs_f64() * 1000.0 / iterations as f64;
        println!(
            "  ⭐ Star 1: {} ({:.3}ms avg over {} runs)",
            result1, avg1_ms, iterations
        );

        let start = Instant::now();
        let mut result2 = String::new();
        for _ in 0..iterations {
            result2 = (day.star2)(day.input);
        }
        let elapsed2 = start.elapsed();
        let avg2_ms = elapsed2.as_secs_f64() * 1000.0 / iterations as f64;
        println!(
            "  ⭐ Star 2: {} ({:.3}ms avg over {} runs)",
            result2, avg2_ms, iterations
        );

        let total_avg = avg1_ms + avg2_ms;
        println!("  ─────────────────────────────────────");
        println!("  Total: {:.3}ms avg", total_avg);
        total_avg
    } else {
        // Normal mode
        let start = Instant::now();
        let result1 = (day.star1)(day.input);
        let elapsed1 = start.elapsed();
        let ms1 = elapsed1.as_secs_f64() * 1000.0;
        println!("  ⭐ Star 1: {} ({:.3}ms)", result1, ms1);

        let start = Instant::now();
        let result2 = (day.star2)(day.input);
        let elapsed2 = start.elapsed();
        let ms2 = elapsed2.as_secs_f64() * 1000.0;
        println!("  ⭐ Star 2: {} ({:.3}ms)", result2, ms2);

        let total = ms1 + ms2;
        println!("  ─────────────────────────────────────");
        println!("  Total: {:.3}ms", total);
        total
    };
    println!();
    day_time_ms
}

fn print_usage() {
    println!("Advent of Code 2025 - Rust Solutions");
    println!();
    println!("Usage: aoc-2025 [OPTIONS] [DAY]");
    println!();
    println!("Arguments:");
    println!("  [DAY]  Run a specific day (1-25). If omitted, runs all days.");
    println!();
    println!("Options:");
    println!("  --bench [N]  Benchmark mode, run each solution N times (default: 100)");
    println!("  --help       Show this help message");
    println!();
    println!("Examples:");
    println!("  aoc-2025           Run all implemented days");
    println!("  aoc-2025 5         Run day 5 only");
    println!("  aoc-2025 --bench   Benchmark all days (100 iterations)");
    println!("  aoc-2025 --bench 1000 5  Benchmark day 5 with 1000 iterations");
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let days = get_days();

    let mut bench_iterations: Option<u32> = None;
    let mut specific_day: Option<u8> = None;
    let mut i = 1;

    while i < args.len() {
        match args[i].as_str() {
            "--help" | "-h" => {
                print_usage();
                return;
            }
            "--bench" | "-b" => {
                // Check if next arg is a number (iterations)
                if i + 1 < args.len() {
                    if let Ok(n) = args[i + 1].parse::<u32>() {
                        bench_iterations = Some(n);
                        i += 1;
                    } else {
                        bench_iterations = Some(100);
                    }
                } else {
                    bench_iterations = Some(100);
                }
            }
            arg => {
                if let Ok(day) = arg.parse::<u8>() {
                    specific_day = Some(day);
                } else {
                    eprintln!("Unknown argument: {}", arg);
                    print_usage();
                    return;
                }
            }
        }
        i += 1;
    }

    println!();
    println!("🎄 Advent of Code 2025 🎄");
    if bench_iterations.is_some() {
        println!("   (Benchmark Mode)");
    }
    println!();

    let mut total_time_ms = 0.0;

    if let Some(day_num) = specific_day {
        if let Some(day) = days.iter().find(|d| d.number == day_num) {
            total_time_ms += run_day(day, bench_iterations);
        } else {
            eprintln!("Day {} is not implemented yet!", day_num);
            eprintln!(
                "Implemented days: {:?}",
                days.iter().map(|d| d.number).collect::<Vec<_>>()
            );
        }
    } else {
        for day in &days {
            total_time_ms += run_day(day, bench_iterations);
        }
    }

    println!("═══════════════════════════════════════");
    println!("  Total runtime: {:.3}ms", total_time_ms);
    println!("═══════════════════════════════════════");
}
