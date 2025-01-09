use criterion::{black_box, criterion_group, criterion_main, Criterion};
use lib_echo1::echo1;

fn echo1_benchmark(c: &mut Criterion) {
    let args: Vec<String> = vec![
        "ignored".to_string(),
        "Hello".to_string(),
        "from".to_string(),
        "Rust".to_string(),
        "Benchmarking".to_string(),
    ];

    c.bench_function("echo1_benchmark", |b| b.iter(|| echo1(black_box(&args))));
}

fn echo1_original_benchmark(c: &mut Criterion) {
    let args: Vec<String> = vec![
        "ignored".to_string(),
        "Hello".to_string(),
        "from".to_string(),
        "Rust".to_string(),
        "Benchmarking".to_string(),
    ];
    c.bench_function("echo1_original_benchmark", |b| {
        b.iter(|| echo1(black_box(&args)))
    });
}

criterion_group!(benches, echo1_benchmark, echo1_original_benchmark);
criterion_main!(benches);
