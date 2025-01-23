fn comma(s: &str) -> String {
    let n = s.len();
    if n <= 3 {
        return s.to_string();
    }

    let prefix = &s[..n - 3];
    let suffix = &s[n - 3..];

    format!("{},{}", comma(prefix), suffix)
}

fn main() {
    println!("{}", comma("somethingnew"));
}
