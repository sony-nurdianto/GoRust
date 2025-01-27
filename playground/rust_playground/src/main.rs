#[derive(Debug)]
struct Wheel {
    radius: i32,
}

fn main() {
    let s = String::from("Something new");
    let s = s.replace("S", "X");

    let mut w = Wheel { radius: 30 };

    w.radius = 20;

    println!("{}", s);
    println!("{:?}", w);
}
