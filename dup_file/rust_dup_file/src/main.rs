use std::{
    collections::HashMap,
    env::{self},
    fs::File,
    io::{self, BufRead, BufReader},
};

use regex::Regex;

fn content_word<R: BufRead>(reader: R, counts: &mut HashMap<String, u32>) {
    let re = Regex::new(r"[^a-zA-Z0-9]").unwrap();

    reader.lines().for_each(|line| match line {
        Ok(value) => {
            value.split_whitespace().for_each(|word| {
                let clean_word = re.replace_all(word, "").to_string();
                if !clean_word.is_empty() {
                    *counts.entry(clean_word).or_insert(0) += 1
                }
            });
        }
        Err(err) => {
            eprintln!("Error reading line: {}", err)
        }
    });
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let mut counts: HashMap<String, u32> = HashMap::new();

    if args[1..].is_empty() {
        let input = io::stdin().lock();
        content_word(input, &mut counts);
    } else {
        let file = File::open(&args[1]).unwrap_or_else(|err| {
            eprintln!("Error opening file: {}", err);
            std::process::exit(1);
        });

        let reader = BufReader::new(file);
        content_word(reader, &mut counts);
    }

    counts.iter().for_each(|(k, v)| println!("{}: {}", k, v));
}
