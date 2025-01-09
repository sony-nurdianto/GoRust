use std::{
    collections::HashMap,
    io::{self, BufRead},
};

/// Program utama untuk menghitung frekuensi kemunculan setiap baris input
/// yang dibaca dari stdin.
///
/// Program ini akan membaca setiap baris yang dimasukkan, menghitung
/// jumlah kemunculan setiap baris, dan kemudian mencetak hasilnya dalam
/// bentuk `HashMap` di mana kunci adalah baris input dan nilai adalah
/// jumlah kemunculannya.
fn main() {
    // Membuat HashMap untuk menyimpan jumlah kemunculan setiap baris.
    let mut counts: HashMap<String, u32> = HashMap::new();

    // Membaca input dari stdin dan memproses setiap barisnya.
    io::stdin().lock().lines().for_each(|line| {
        if let Ok(value) = line {
            // Memperbarui jumlah kemunculan untuk setiap baris
            let count = counts.entry(value).or_insert(0);
            *count += 1;
        }
    });

    // for line in input.lock().lines() {
    //     match line {
    //         Ok(value) => {
    //             let count = counts.entry(value).or_insert(0);
    //             *count += 1;
    //         }
    //         Err(_) => continue,
    //     };
    // }

    // Menampilkan hasil penghitungan
    println!("{:?}", counts);
}
