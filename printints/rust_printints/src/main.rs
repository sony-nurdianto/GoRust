fn int_to_string(value: &[i32]) -> String {
    let formated = value
        .iter()
        .map(|v| v.to_string())
        .collect::<Vec<String>>()
        .join(",");

    format!("[{}]", formated)
}

fn main() {
    let r = [1, 2, 3, 4, 5, 6, 7];

    println!("{}", int_to_string(&r));
}
