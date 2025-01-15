fn main() {
    let freezing_f = 32.0;
    let boilling_c = 212.0;

    println!("{}°F = {}°C", &freezing_f, f_to_c(&freezing_f));
    println!("{}°F = {}°C", &boilling_c, f_to_c(&boilling_c));
}

fn f_to_c(f: &f64) -> f64 {
    return (f - 32_f64) * 5_f64 / 9_f64;
}
