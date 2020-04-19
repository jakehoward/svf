use csv::ReaderBuilder;
use std::io;
use structopt::StructOpt;

/// Read newline delimited data from stdin and structurally parse based on delimiters
#[derive(StructOpt, Debug)]
struct Cli {
    /// The character delimiter used to split a line into fields
    #[structopt(short = "d", long = "delimiter", default_value = ",")]
    delimiter: String,

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

struct SvfBufReader {
    buffer: Vec<u8>
}

impl SvfBufReader {
    fn new() -> SvfBufReader {
        // TODO: check the performance impact
        SvfBufReader { buffer: Vec::with_capacity(1000) }
    }

    // fn read_record(&self) {
        // vec!["testing", "testing", "123"]
    // }
}

fn main() -> io::Result<()> {
    let args = Cli::from_args();

    let mut rdr = ReaderBuilder::new()
        .delimiter(args.delimiter.as_bytes()[0])
        .from_reader(io::stdin());

    for record in rdr.records() {
        println!("Record is: {:?}", record);
    }

    Ok(())
}
