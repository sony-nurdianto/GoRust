use std::{
    collections::HashMap,
    io::{self, BufRead},
};

fn main() {
    let mut counts: HashMap<String, u32> = HashMap::new();

    io::stdin().lock().lines().for_each(|line| match line {
        Ok(value) => {
            let count = counts.entry(value).or_insert(0);
            *count += 1;
        }
        Err(_) => {}
    });

    // for line in input.lock().lines() {
    //     match line {
    //         Ok(value) => {
    //             let count = counts.entry(value).or_insert(0);
    //             *count += 1;
    //         }
    //         Err(_) => continue,
    //     };
    // }

    println!("{:?}", counts);
}
