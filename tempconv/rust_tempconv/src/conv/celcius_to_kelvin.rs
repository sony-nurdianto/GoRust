use super::KELVIN_SCALE;

pub fn convert(c: &str) -> f64 {
    let c: f64 = c.parse().unwrap_or(0.0);

    c + KELVIN_SCALE
}
