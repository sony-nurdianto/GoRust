use std::{
    collections::HashMap,
    env,
    fmt::Formatter,
    fs,
    path::Path,
    sync::{Arc, Mutex},
};
use std::{fmt, io::Error as IoError, string::FromUtf8Error};

use lazy_static::lazy_static;
use rayon::{
    iter::{IntoParallelRefIterator, ParallelIterator},
    slice::ParallelSliceMut,
    str::ParallelString,
};
use regex::{Error as RgxError, Regex};

lazy_static! {
    static ref RE: Regex = Regex::new(r"[^a-zA-Z0-9]").unwrap();
}

#[derive(Debug)]
enum DupError {
    Io(IoError),
    Utf8(FromUtf8Error),
    Regex(RgxError),
}

impl fmt::Display for DupError {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        match self {
            DupError::Io(err) => write!(f, "I/O Error: {}", err),
            DupError::Utf8(err) => write!(f, "UTF-8 Error: {}", err),
            DupError::Regex(err) => write!(f, "Regex Error: {}", err),
        }
    }
}

impl From<IoError> for DupError {
    fn from(value: IoError) -> Self {
        DupError::Io(value)
    }
}

impl From<FromUtf8Error> for DupError {
    fn from(value: FromUtf8Error) -> Self {
        DupError::Utf8(value)
    }
}

impl From<RgxError> for DupError {
    fn from(value: RgxError) -> Self {
        DupError::Regex(value)
    }
}

fn find_duplicate(path: &Path, counts: &Arc<Mutex<HashMap<String, u32>>>) -> Result<(), DupError> {
    let fs = fs::read(path)?;
    let text = String::from_utf8(fs)?;
    let data = Arc::clone(counts);

    text.par_split_whitespace().for_each(move |word| {
        let clean_word = RE.replace_all(word, "").to_string();
        if !clean_word.is_empty() {
            match data.lock() {
                Ok(mut value) => *value.entry(clean_word).or_insert(0) += 1,
                Err(err) => eprintln!("Failed to lock data {}", err),
            }
        }
    });

    Ok(())
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let file: &[String] = &args[1..];

    file.iter().for_each(|arg| {
        println!("Finding duplicate Word In: {}", arg);
        let path: &Path = Path::new(arg);
        let counts = Arc::new(Mutex::new(HashMap::<String, u32>::new()));

        match find_duplicate(path, &counts) {
            Ok(_) => {
                let count = counts.lock().unwrap();
                let mut sorted_counts: Vec<(&String, &u32)> = count.par_iter().collect();
                sorted_counts.par_sort_by(|a, b| b.1.cmp(a.1));
                sorted_counts
                    .iter()
                    .for_each(|(k, v)| println!("{}: {}", k, v));
            }
            Err(err) => {
                eprintln!("Error Processing File {}: {}", arg, err)
            }
        };
    });
}
