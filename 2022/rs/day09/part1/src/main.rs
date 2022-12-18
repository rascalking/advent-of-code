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
struct Grid {
    head: Position,
    tail: Position,
    tail_seen: HashSet<Position>,
}

impl Grid {
    fn new() -> Self {
        Self{
            head: (0,0),
            tail: (0,0),
            tail_seen: HashSet::from([(0,0)]),
        }
    }

    fn apply_move(&mut self, motion: Motion) {
        println!("{:?} {}", motion.direction, motion.count);
        for _ in 0..motion.count {
            let tail_start = self.tail;
            match motion.direction {
                Direction::Left => {
                    self.head.0 -= 1;
                },
                Direction::Right => {
                    self.head.0 += 1;
                },
                Direction::Up => {
                    self.head.1 += 1;
                },
                Direction::Down => {
                    self.head.1 -= 1;
                },
            }
            let diff = self.head.sub(self.tail);
            match diff {
                (0, 0) | (0, 1) | (0, -1) | (1, 0) | (1, 1) | (1, -1) | (-1, 0) | (-1, 1) | (-1, -1) => {}
                _ => { 
                    self.tail.0 += diff.0.signum();
                    self.tail.1 += diff.1.signum();
                }
            }
            println!("H({}, {}), T({}, {}) -> ({}, {})",
                self.head.0, self.head.1, tail_start.0, tail_start.1, self.tail.0, self.tail.1);
            self.tail_seen.insert(self.tail);
        }
    }
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let mut grid = Grid::new();
    for line in contents.lines() {
        grid.apply_move(Motion::new(line));
    }
    dbg!(grid.tail_seen.len());
    Ok(())
}
