use super::{FAHRENHEIT_OFFSET, FAHRENHEIT_SCALE, KELVIN_SCALE};

pub fn convert(k: &str) -> f64 {
    let k: f64 = k.parse().unwrap_or(0.0);

    ((k - KELVIN_SCALE) * FAHRENHEIT_SCALE) + FAHRENHEIT_OFFSET
}
