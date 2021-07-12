use std::env;
use std::fs::File;
use std::io::prelude::*;

fn read_numbers<T: std::str::FromStr>(filename: &str) -> Result<Vec<T>, std::io::Error> {
    let mut file = File::open(filename)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    let lines: Vec<&str> = contents.split('\n').collect();
    let mut numbers: Vec<T> = Vec::<T>::new();
    for line in &lines {
        if let Ok(num) = line.trim().parse::<T>() {
            numbers.push(num);
        }
    }
    Ok(numbers)
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let numbers = read_numbers::<i32>(&args[1])?;
    //println!("{:?}", numbers);
    
    for i in 0..numbers.len() {
        for j in (i+1)..numbers.len() {
            for k in (j+1)..numbers.len() {
                if numbers[i] + numbers[j] + numbers[k] == 2020 {
                    println!("{}", numbers[i] * numbers[j] * numbers[k]);
                    return Ok(())
                }
            }
        }
    }

    Ok(())
}

