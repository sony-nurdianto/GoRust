use clap::{Arg, ArgAction, Command};
use conv::{
    celcius_to_fahrenheit, celcius_to_kelvin, fahrenheit_to_celcius, fahrenheit_to_kelvin,
    kelvin_to_celcius, kelvin_to_fahrenheit,
};

mod conv;

fn main() {
    const CTF_ID: &str = "celcius_to_fahrenheit";
    const CTK_ID: &str = "celcius_to_kelvin";
    const FTC_ID: &str = "fahrenheit_to_celcius";
    const FTK_ID: &str = "fahrenheit_to_kelvin";
    const KTC_ID: &str = "kelvin_to_celcius";
    const KTF_ID: &str = "kelvin_to_fahrenheit";

    let args = Command::new("Temperature Converter")
        .arg(
            Arg::new("Convert Celcius To Fahrenheit")
                .id(CTF_ID)
                .long("ctf")
                .required(false)
                .action(ArgAction::Append),
        )
        .arg(
            Arg::new("Convert Celcius To Kelvin")
                .id(CTK_ID)
                .long("ctk")
                .required(false)
                .action(ArgAction::Append),
        )
        .arg(
            Arg::new("Convert Fahrenheit To Celcius")
                .id(FTC_ID)
                .long("ftc")
                .required(false)
                .action(ArgAction::Append),
        )
        .arg(
            Arg::new("Convert Fahrenheit To Kelvin")
                .id(FTK_ID)
                .long("ftk")
                .required(false)
                .action(ArgAction::Append),
        )
        .arg(
            Arg::new("Kelvin To Celcius")
                .id(KTC_ID)
                .long("ktc")
                .required(false)
                .action(ArgAction::Append),
        )
        .arg(
            Arg::new("Kelvin to Fahrenheit")
                .id("kelvin_to_fahrenheit")
                .long(KTF_ID)
                .required(false)
                .action(ArgAction::Append),
        )
        .get_matches();

    if let Some(ctf) = args.get_one::<String>(CTF_ID) {
        println!("{}°F", celcius_to_fahrenheit::convert(ctf))
    }

    if let Some(ctk) = args.get_one::<String>(CTK_ID) {
        println!("{}°K", celcius_to_kelvin::convert(ctk))
    }

    if let Some(ftc) = args.get_one::<String>(FTC_ID) {
        println!("{}°C", fahrenheit_to_celcius::convert(ftc))
    }

    if let Some(ftk) = args.get_one::<String>(FTK_ID) {
        println!("{}°K", fahrenheit_to_kelvin::convert(ftk))
    }

    if let Some(ktc) = args.get_one::<String>(KTC_ID) {
        println!("{}°C", kelvin_to_celcius::convert(ktc))
    }

    if let Some(ktf) = args.get_one::<String>(KTF_ID) {
        println!("{}°F", kelvin_to_fahrenheit::convert(ktf))
    }
}
