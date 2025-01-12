use core::fmt;
use std::{env, time};

use tokio::{
    io::{self},
    sync::mpsc::{self, UnboundedSender},
};

#[derive(Debug)]
enum FetchAllErr {
    ReqwestErr(reqwest::Error),
    TokioIoErr(io::Error),
}

impl fmt::Display for FetchAllErr {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::ReqwestErr(err) => write!(f, "Reqwest Error: {}", err),
            Self::TokioIoErr(err) => write!(f, "Tokio I/O Error: {}", err),
        }
    }
}

impl From<reqwest::Error> for FetchAllErr {
    fn from(value: reqwest::Error) -> Self {
        FetchAllErr::ReqwestErr(value)
    }
}

impl From<io::Error> for FetchAllErr {
    fn from(value: io::Error) -> Self {
        FetchAllErr::TokioIoErr(value)
    }
}

#[tokio::main]
async fn main() {
    let start = time::Instant::now();
    let (tx, mut rx) = mpsc::unbounded_channel::<String>();
    let args: Vec<String> = env::args().collect();

    for url in args.into_iter().skip(1) {
        let transmit = tx.clone();
        tokio::spawn(async move { fetch(&url, transmit).await });
    }

    drop(tx);

    while let Some(msg) = rx.recv().await {
        println!("{}", msg);
    }

    println!("{:.2} s elapsed\n", start.elapsed().as_secs_f64())
}

async fn fetch(url: &str, tx: UnboundedSender<String>) -> Result<(), FetchAllErr> {
    let start = time::Instant::now();

    match reqwest::get(url).await {
        Ok(val) => {
            let body = val.bytes().await?;
            let nbytes = body.len() as u64;
            io::copy(&mut body.as_ref(), &mut io::sink()).await?;
            let message = format!(
                "{:.2} s  {:7}  {}",
                start.elapsed().as_secs_f64(),
                nbytes,
                url
            );
            tx.send(message).unwrap();
        }
        Err(err) => tx.send(err.to_string()).unwrap(),
    };

    Ok(())
}
