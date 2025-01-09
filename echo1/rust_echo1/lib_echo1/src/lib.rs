pub fn original_echo1(args: &Vec<String>) {
    let mut s: String = String::new();

    for item in args {
        s += item;
        s += " ";
    }

    println!("{}", s);
}

pub fn echo1(args: &Vec<String>) {
    println!("{}", args[1..].join(" "));
}
