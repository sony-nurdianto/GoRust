use std::env;

fn main() {
    let input: Vec<String> = env::args().collect();
    lib_echo1::echo1(&input);
    lib_echo1::original_echo1(&input);
}
