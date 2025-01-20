use super::{FAHRENHEIT_OFFSET, FAHRENHEIT_SCALE};

pub fn convert(f: &str) -> f64 {
    let f: f64 = f.parse().unwrap_or(0.0);

    (f * FAHRENHEIT_SCALE) + FAHRENHEIT_OFFSET
}
