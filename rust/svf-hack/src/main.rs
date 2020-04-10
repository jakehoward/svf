use std::io::{self, Read};
use structopt::StructOpt;

/// Read newline delimited data from stdin and structurally parse based on delimiters
#[derive(StructOpt, Debug)]
struct Cli {
    /// The character delimiter used to split a line into fields
    #[structopt(short = "d", long = "delimiter", default_value = ",")]
    delim: String,

    /// The fields you would like to return. E.g. for the first, fifth to seventh inclusive, and tenth onwards: '1,5-7,10-' (spaces will be ignored)
    #[structopt(short = "f", long = "fields")]
    fields: String,

    /// Escape character, fields wrapped in this character will not be split by the delimiter. Double escape character will not escape and be read as the character.
    #[structopt(short = "e", long = "escape-character", default_value = "\"")]
    escape_char: String,

    /// Use zero based indexing for fields
    #[structopt(short = "z", long = "zero-index")]
    zero_index: bool,
}

fn main() -> io::Result<()> {
    let args = Cli::from_args();
    // println!("Args: {:?}", args);

    let mut content = String::new();
    match io::stdin().read_to_string(&mut content) {
        Ok(_) => (),
        Err(e) => return Err(e),
    }
    // println!("{}", content);

    // println!("{:?}", "a,b,c,d".split(&args.delim).collect::<Vec<&str>>());

    let lines: Vec<&str> = content.lines().collect();
    let items: Vec<Vec<&str>> = lines.into_iter().map(|x| x.split(&args.delim).collect()).collect();
    println!("{:?}", items);
    Ok(())
}
