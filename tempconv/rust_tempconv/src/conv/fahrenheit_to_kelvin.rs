use super::{FAHRENHEIT_OFFSET, KELVIN_SCALE};

pub fn convert(f: &str) -> f64 {
    let f = f.parse().unwrap_or(0.0);

    ((f - FAHRENHEIT_OFFSET) * 5.0 / 9.0) + KELVIN_SCALE
}
