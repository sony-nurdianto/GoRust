use core::panic;
use std::{f64::consts::PI, fs::File, io::Write, path::Path};

const WIDTH: usize = 600;
const HEIGHT: usize = 320;
const CELLS: usize = 100;
const XYRANGE: f64 = 30.0;
const XYSCALE: f64 = WIDTH as f64 / 2.0 / XYRANGE;
const ZSCALE: f64 = HEIGHT as f64 * 0.5;
const ANGLE: f64 = PI / 6.0;

fn main() -> std::io::Result<()> {
    let file = match File::create(Path::new("output.svg")) {
        Ok(val) => val,
        Err(err) => {
            panic!("Failed To Open File: {}", err);
        }
    };

    let mut file = std::io::BufWriter::new(file);

    writeln!(
        &mut file,
        "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='{}' height='{}'>",
        WIDTH, HEIGHT
    )?;

    (0..CELLS).for_each(|i| {
        (0..CELLS).for_each(|j| {
            let (ax, ay) = corner(i + 1, j);
            let (bx, by) = corner(i, j);
            let (cx, cy) = corner(i, j + 1);
            let (dx, dy) = corner(i + 1, j + 1);

     match writeln!(&mut file, "<polygon points='{:.2},{:.2} {:.2},{:.2} {:.2},{:.2} {:.2},{:.2}' style='fill: none; stroke: gray' />",
                     ax, ay, bx, by, cx, cy, dx, dy) {
                     Ok(v) => v,
                     Err(err) => panic!("{}",err)
                 };
        });
    });

    writeln!(&mut file, "</svg>")?;
    Ok(())
}

fn corner(i: usize, j: usize) -> (f64, f64) {
    let x = XYRANGE * (i as f64 / CELLS as f64 - 0.5);
    let y = XYRANGE * (j as f64 / CELLS as f64 - 0.5);
    let z = f(x, y);

    let sin30 = ANGLE.sin();
    let cos30 = ANGLE.cos();

    let sx = WIDTH as f64 / 2.0 + (x - y) * cos30 * XYSCALE;
    let sy = HEIGHT as f64 / 2.0 + (x + y) * sin30 * XYSCALE - z * ZSCALE;
    (sx, sy)
}

fn f(x: f64, y: f64) -> f64 {
    let r = (x * x + y * y).sqrt();
    if r == 0.0 { 0.0 } else { (r).sin() / r }
}
