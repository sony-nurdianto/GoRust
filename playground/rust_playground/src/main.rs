fn raaa(n: i32) -> i32 {
    if n >= 10 {
        return n;
    }

    raaa(n + 1)
}

fn main() {
    println!("{}", raaa(0));
}
