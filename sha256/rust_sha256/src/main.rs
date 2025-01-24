use sha2::{Digest, Sha256};

fn main() {
    let c1 = Sha256::digest(b"x");
    let c2 = Sha256::digest(b"X");
    println!("{:x}\n{:x}\n{}\n{:?}", c1, c2, c1 == c2, c1);
}
