use std::env;
use std::fs::File;
use std::io::read_to_string;

fn read_lines(filename: &str) -> Result<Vec<String>, std::io::Error> {
    let contents = read_to_string(File::open(filename)?)?;
    let lines: Vec<&str> = contents.split('\n').collect();
    // there has to be a less messy way to do this
    Ok(lines.iter().map(|&line| String::from(line)).collect())
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let lines = read_lines(&args[1])?;
    //println!("{:?}", lines);

    let mut top = vec![0, 0, 0];
    let mut current = 0;
    for n in &lines {
        if n.trim() == "" {
            top.push(current);
            top.sort_by(|a, b| b.cmp(a));
            top.pop();

            current = 0;
        } else {
            let n: u32 = n.trim().parse().expect("failed to parse line");
            current += n;
        }
    }
    let sum: u32 = top.iter().sum();
    println!("{}", sum);
    Ok(())
}
