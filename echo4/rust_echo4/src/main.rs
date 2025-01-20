use clap::{Arg, Command};

fn main() {
    let matches = Command::new("example")
        .arg(
            Arg::new("n")
                .short('n')
                .long("newline")
                .help("Omit trailing newline")
                .action(clap::ArgAction::SetTrue),
        )
        .arg(
            Arg::new("s")
                .short('s')
                .long("separator")
                .help("Separator")
                .default_value(" "),
        )
        .arg(
            Arg::new("args")
                .help("Input arguments")
                .action(clap::ArgAction::Append),
        )
        .get_matches();

    let omit_newline = matches.get_flag("n");
    let separator = matches.get_one::<String>("s").unwrap();
    let args: Vec<String> = matches.get_many("args").unwrap().cloned().collect();

    if !omit_newline {
        return println!("{}", args.join(separator));
    }

    print!("{}", args.join(separator));
}
