use std::collections::HashSet;
use std::env;
use std::fs::File;
use std::io::read_to_string;

#[derive(Debug)]
enum Direction { Left, Right, Up, Down }

type Position = (i32, i32);
trait Sub {
    fn sub(self, other: Self) -> Self;
}

impl Sub for Position {
    fn sub(self, other: Self) -> Self {
        (self.0-other.0, self.1-other.1)
    }
}

#[derive(Debug)]
struct Motion {
    direction: Direction,
    count: u32,
}

impl Motion {
    fn new(line: &str) -> Self {
        let parts: Vec<&str> = line.splitn(2, ' ').collect();
        Self {
            direction: match parts[0] {
                "L" => Direction::Left,
                "R" => Direction::Right,
                "U" => Direction::Up,
                "D" => Direction::Down,
                _ => panic!(),
            },
            count: parts[1].parse().unwrap(),
        }
    }
}

#[derive(Debug)]
struct Rope {
    knots: Vec<Position>,
    tail_seen: HashSet<Position>,
}

impl Rope {
    fn new() -> Self {
        Self{
            knots: (0..10).map(|_| (0,0)).collect(),
            tail_seen: HashSet::from([(0,0)]),
        }
    }

    fn apply_move(&mut self, motion: Motion) {
        println!("{:?}", self.knots);
        for _ in 0..motion.count {
            // move the actual head
            match motion.direction {
                Direction::Left => {
                    self.knots[0].0 -= 1;
                },
                Direction::Right => {
                    self.knots[0].0 += 1;
                },
                Direction::Up => {
                    self.knots[0].1 += 1;
                },
                Direction::Down => {
                    self.knots[0].1 -= 1;
                },
            }
            // now pull the tail knots behind it
            for head in 0..9 {
                let tail = head + 1;
                let diff = self.knots[head].sub(self.knots[tail]);
                match diff {
                    (0, 0) | (0, 1) | (0, -1) | (1, 0) | (1, 1) | (1, -1) | (-1, 0) | (-1, 1) | (-1, -1) => {}
                    _ => { 
                        self.knots[tail].0 += diff.0.signum();
                        self.knots[tail].1 += diff.1.signum();
                    }
                }
            }
            println!("{:?}", self.knots);
            self.tail_seen.insert(self.knots[9]);
        }
    }
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let mut rope = Rope::new();
    for line in contents.lines() {
        rope.apply_move(Motion::new(line));
    }
    dbg!(rope.tail_seen.len());
    Ok(())
}
