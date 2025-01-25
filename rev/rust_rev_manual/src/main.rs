fn reverse(s: &mut [i32]) {
    let mut i = 0;
    let mut j = s.len() - 1;

    while i < j {
        s.swap(i, j);
        i += 1;
        j -= 1;
    }
}

fn main() {
    let mut arr = [1, 2, 3, 4, 5];
    reverse(&mut arr);
    println!("{:?}", arr);
}
