fn basename(s: &str) -> String {
    let mut s = s;

    if let Some(pos) = s.rfind('/') {
        s = &s[pos + 1..];
    }

    if let Some(pos) = s.rfind('.') {
        s = &s[..pos];
    }

    s.to_string()
}

fn main() {
    println!("{}", basename("a/b/c.go"));
    println!("{}", basename("c.d.go"));
    println!("{}", basename("abc"));
}
