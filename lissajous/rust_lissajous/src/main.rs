use std::{
    f64::{self, consts::PI},
    fs::File,
    i32,
};

use image::{
    codecs::gif::{GifEncoder, Repeat},
    Delay, Frame, ImageBuffer, RgbaImage,
};
use rand::Rng;

const CYCLES: f64 = 5.0;
const RES: f64 = 0.00_1;
const SIZE: u32 = 100;
const NFRAMES: u32 = 64;
const DELAY: u32 = 8;

const PALETE: [image::Rgba<u8>; 8] = [
    image::Rgba([255, 255_u8, 255_u8, 255_u8]), // White
    image::Rgba([0, 0, 0, 255]),                // Black
    image::Rgba([255, 0, 0, 255]),              // Red
    image::Rgba([0, 255, 0, 255]),              // Green
    image::Rgba([0, 0, 255, 255]),              // Blue
    image::Rgba([255, 255, 0, 255]),            // Yellow
    image::Rgba([255, 0, 255, 255]),            // Purple
    image::Rgba([0, 255, 255, 255]),            // Cyan
];

fn main() {
    let something = File::create("out.gif").unwrap();
    let freq = rand::thread_rng().gen_range(0.0..3.0);
    let mut anim = GifEncoder::new(something);
    anim.set_repeat(Repeat::Infinite).unwrap();

    let mut phase: f64 = 0.0;

    let mut index_palete = 1;

    (0..NFRAMES).for_each(|_| {
        if index_palete >= 7 {
            index_palete = 0;
        }

        let mut img: RgbaImage = ImageBuffer::from_pixel(2 * SIZE + 1, 2 * SIZE * 1, PALETE[0]);

        let mut t = 0.0;
        while t < CYCLES * 2.0 * PI {
            let x = (t.sin() * SIZE as f64) as i32;
            let y = ((t * freq + phase).sin() * SIZE as f64) as i32;

            let pixel_x = SIZE as i32 + x;
            let pixel_y = SIZE as i32 + y;

            if pixel_x >= 0 && pixel_y >= 0 {
                img.put_pixel(pixel_x as u32, pixel_y as u32, PALETE[index_palete]);
            }

            t += RES;
        }
        index_palete += 1;
        phase += 0.1;
        let frame = Frame::from_parts(img, 0, 0, Delay::from_numer_denom_ms(DELAY, 1));
        anim.encode_frame(frame).unwrap();
    });
}
