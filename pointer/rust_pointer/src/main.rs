use std::time;

fn main() {
    let start = time::Instant::now();
    let mut p: i32 = 0;

    (0..10).for_each(|_| {
        count(&mut p);
    });

    println!("{}", p);

    let elapsed = start.elapsed();
    println!("Elapsed time{:.2?}", elapsed);
}

fn count(p: &mut i32) {
    *p += 1;
}
