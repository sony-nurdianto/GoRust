use super::FAHRENHEIT_OFFSET;

pub fn convert(c: &str) -> f64 {
    let c: f64 = c.parse().unwrap_or(0_f64);

    (c * FAHRENHEIT_OFFSET) + 32.0
}
