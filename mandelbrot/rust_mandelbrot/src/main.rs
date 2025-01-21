use image::{ImageBuffer, Rgba};
use num::Complex;

const XMIN: f64 = -2.0;
const YMIN: f64 = -2.0;
const XMAX: f64 = 2.0;
const YMAX: f64 = 2.0;
const HEIGHT: usize = 1024;
const WIDTH: usize = 1024;

fn main() -> std::io::Result<()> {
    let mut img = ImageBuffer::new(HEIGHT as u32, WIDTH as u32);

    (0..HEIGHT).for_each(|py| {
        let y = (py as f64) / (HEIGHT as f64) * (YMAX - YMIN) + YMIN;
        (0..WIDTH).for_each(|px| {
            let x = (px as f64) / (WIDTH as f64) * (XMAX - XMIN) + XMIN;
            let z = num::Complex::new(x, y);

            img.put_pixel(px as u32, py as u32, mandelbrot(z));
        });
    });

    img.save("out.png").expect("Failed To Create Image");

    Ok(())
}

fn mandelbrot(z: Complex<f64>) -> Rgba<u8> {
    const CONTRAST: u32 = 25;
    const ITERATIONS: u8 = 200;

    let mut v = Complex::new(0.0, 0.0);
    for n in 0..ITERATIONS {
        v = v * v + z;
        if v.norm().abs() > 2.0 {
            let color: u8 = 255 - (CONTRAST * n as u32) as u8;
            return Rgba([color, color, color, 255]);
        }
    }

    Rgba([0, 0, 0, 255])
}
