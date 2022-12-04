use std::env;
use std::fs::File;
use std::io::read_to_string;


#[derive(Debug)]
enum RPS { Rock, Paper, Scissors }

#[derive(Debug)]
struct Round {
    opponent: RPS,
    player: RPS,
}

impl Round {
    fn score(&self) -> u32 {
        let shape = match self.player {
            RPS::Rock => 1,
            RPS::Paper => 2,
            RPS::Scissors => 3,
        };
        let outcome = match (&self.player, &self.opponent) {
            (RPS::Rock, RPS::Rock) => 3,
            (RPS::Rock, RPS::Paper) => 0,
            (RPS::Rock, RPS::Scissors) => 6, 
            (RPS::Paper, RPS::Rock) => 6,
            (RPS::Paper, RPS::Paper) => 3,
            (RPS::Paper, RPS::Scissors) => 0, 
            (RPS::Scissors, RPS::Rock) => 0,
            (RPS::Scissors, RPS::Paper) => 6,
            (RPS::Scissors, RPS::Scissors) => 3,
        };
        shape + outcome
    }
}

#[derive(Debug, Clone)]
struct ParseError;


fn parse_rounds(lines: Vec<String>) -> Result<Vec<Round>, ParseError> {
    let mut rounds = Vec::new();
    for line in &lines {
        let (opponent, player): (RPS, RPS);
        let plays: Vec<&str> = line.splitn(2, ' ').collect();
        match plays[0] {
            "A" => opponent = RPS::Rock,
            "B" => opponent = RPS::Paper,
            "C" => opponent = RPS::Scissors,
            _ => {
                println!("{:?}", plays[0]);
                return Err(ParseError);
            },
        }
        match plays[1] {
            "X" => {
                // lose
                player = match opponent {
                    RPS::Rock => RPS::Scissors,
                    RPS::Paper => RPS::Rock,
                    RPS::Scissors => RPS::Paper,
                }
            },
            "Y" => {
                // draw
                player = match opponent {
                    RPS::Rock => RPS::Rock,
                    RPS::Paper => RPS::Paper,
                    RPS::Scissors => RPS::Scissors,
                }
            },
            "Z" => {
                // win
                player = match opponent {
                    RPS::Rock => RPS::Paper,
                    RPS::Paper => RPS::Scissors,
                    RPS::Scissors => RPS::Rock,
                }
            },
            _ => {
                println!("{:?}", plays[0]);
                return Err(ParseError);
            },
        }
        rounds.push(Round{opponent: opponent, player: player});
    }
    Ok(rounds)
}

fn read_lines(filename: &str) -> Result<Vec<String>, std::io::Error> {
    let contents = read_to_string(File::open(filename)?)?;
    let lines: Vec<&str> = contents.trim().split('\n').collect();
    Ok(lines.iter().map(|&line| String::from(line)).collect())
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let lines = read_lines(&args[1]).expect("unable to read input");
    println!("{:?}", lines);
    let rounds = parse_rounds(lines).expect("unable to parse input");
    let mut total = 0;
    for round in &rounds {
        let score = round.score();
        println!("{:?}: {}", round, score);
        total += score;
    }
    println!("{}", total);
}
