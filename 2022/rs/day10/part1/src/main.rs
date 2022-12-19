use std::env;
use std::fs::File;
use std::io::read_to_string;

enum InstructionKind { Addx, Noop }

struct Instruction {
    kind: InstructionKind,
    num_cycles: u32,
    arg0: i32,
}

impl Instruction {
    fn new(line: &str) -> Self {
        let parts: Vec<&str> = line.splitn(2, ' ').collect();
        match parts[0] {
            "addx" => {
                Self { 
                    kind: InstructionKind::Addx,
                    num_cycles: 2,
                    arg0: parts[1].parse().unwrap(),
                }
            },
            "noop" => {
                Self {
                    kind: InstructionKind::Noop,
                    num_cycles: 1,
                    arg0: 0,
                }
            }
            _ => panic!(),
        }
    }
}

#[derive(Debug)]
struct CPU {
    cycle: i32,
    x: i32,
    signal: i32,
}

impl CPU {
    fn new() -> Self {
        Self {
            cycle: 0,
            x: 1,
            signal: 0,
        }
    }

    fn execute(&mut self, inst: Instruction) {
        for _ in 0..inst.num_cycles {
            self.cycle += 1;
            match self.cycle {
                20 | 60 | 100 | 140 | 180 | 220 => {
                    self.signal += self.cycle * self.x;
                    println!("{:?}", self);
                }
                _ => {}
            }
        }
        match inst.kind {
            InstructionKind::Noop => {},
            InstructionKind::Addx => { self.x += inst.arg0 }
        }
    }
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let mut cpu = CPU::new();
    for line in contents.lines() {
        let inst = Instruction::new(line);
        cpu.execute(inst);
    }
    dbg!(cpu.signal);
    Ok(())
}
