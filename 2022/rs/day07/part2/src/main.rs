use std::collections::HashMap;
use std::collections::HashSet;
use std::env;
use std::fs::File;
use std::io::read_to_string;
use std::path::Path;
use std::path::PathBuf;


fn size(dir: &PathBuf, files: &HashMap<PathBuf, u32>) -> u32 {
    let mut size = 0;
    for (file, fsize) in files {
        if file.starts_with(dir) {
            size += fsize;
        }
    }
    size
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");

    let mut files = HashMap::new();
    let mut dirs = HashSet::from([Path::new("/").to_path_buf()]);
    let mut cwd = PathBuf::new();
    for line in contents.lines() {
        let toks: Vec<_> = line.split_whitespace().collect();
        match toks[0] {
            "$" => {
                if toks[1] == "cd" {
                    match toks[2] {
                        ".." => {
                            cwd.pop();
                        }
                        "/" => {
                            cwd.clear();
                            cwd.push(toks[2]);
                        }
                        _ => {
                            cwd.push(toks[2]);
                        }
                    }
                }
            }
            "dir" => {
                cwd.push(toks[1]);
                dirs.insert(cwd.clone());
                cwd.pop();
            }
            num => {
                let num = num.parse::<u32>().unwrap();
                cwd.push(toks[1]);
                files.insert(cwd.clone(), num);
                cwd.pop();
            }
        }
    }
    dbg!(&dirs);
    dbg!(&files);
    let sizes: HashMap<&PathBuf, u32> =
        HashMap::from_iter(dirs.iter().map(|d| (d, size(d, &files))));
    dbg!(&sizes);

    let root = Path::new("/").to_path_buf();
    let free = 70000000 - sizes.get(&root).unwrap();
    let wanted = 30000000 - free;
    dbg!(free);
    dbg!(wanted);

    let mut sizes = Vec::from_iter(sizes.iter().filter(|&(_,s)| s >= &wanted).map(|(d,s)| (s,d)));
    sizes.sort_by(|a,b| a.cmp(b));
    dbg!(&sizes);
    let (size, dir) = sizes[0];
    dbg!(size, dir);

    Ok(())
}
