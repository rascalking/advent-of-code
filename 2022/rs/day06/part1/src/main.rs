use std::env;
use std::fs::File;
use std::io::read_to_string;
use std::collections::VecDeque;
use std::collections::HashSet;

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let mut found = -1;
    let mut chars = contents.chars();
    let mut buf: VecDeque<char> = VecDeque::with_capacity(4);
    for _ in 0..4 {
        buf.push_back(chars.next().unwrap());
    }
    let mut i = 4;
    for c in chars {
        let set: HashSet<&char> = HashSet::from_iter(buf.iter());
        if set.len() == 4 {
            found = i;
            break;
        }
        buf.pop_front();
        buf.push_back(c);
        i += 1;
    }
    dbg!(found);
    Ok(())
}
