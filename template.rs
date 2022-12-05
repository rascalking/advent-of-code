use std::env;
use std::fs::File;
use std::io::read_to_string;

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let lines = contents.lines().collect::<Vec<_>>();
    println!("{:?}", lines);
    Ok(())
}
