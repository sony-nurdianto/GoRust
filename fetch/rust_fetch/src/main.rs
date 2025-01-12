use core::fmt;
use std::{
    env,
    io::{self, Error, Write},
};

#[derive(Debug)]
enum FetchErr {
    StdError(Error),
    ReqwestError(reqwest::Error),
}

impl fmt::Display for FetchErr {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::StdError(err) => write!(f, "Std Error: {}", err),
            Self::ReqwestError(err) => write!(f, "Reqwest Error: {}", err),
        }
    }
}

impl From<io::Error> for FetchErr {
    fn from(value: io::Error) -> Self {
        FetchErr::StdError(value)
    }
}

impl From<reqwest::Error> for FetchErr {
    fn from(value: reqwest::Error) -> Self {
        FetchErr::ReqwestError(value)
    }
}

#[tokio::main]
async fn main() -> Result<(), FetchErr> {
    let args: Vec<String> = env::args().collect();

    for arg in &args[1..] {
        let mut url = arg.clone();

        if !url.starts_with("http") {
            url = format!("http://{}", url)
        }

        let mut response = reqwest::get(url).await?;
        let mut out = io::stdout();

        while let Some(chunk) = response.chunk().await? {
            out.write(format!("Response Status: {}\n", response.status()).as_bytes())?;
            out.write_all(&chunk)?;
        }
    }

    Ok(())
}
