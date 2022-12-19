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
struct Crt {
    pixels: [[bool; 40]; 6],
}

impl Crt {
    fn new() -> Self {
        Self { pixels: [[false; 40]; 6] }
    }

    fn print(&self) {
        for row in 0..6 {
            println!("{}", self.pixels[row].map(|p| match p {
                true => "#",
                false => ".",
            }).join(""));
        }
    }
}

#[derive(Debug)]
struct Cpu {
    cycle: u32,
    x: i32,
    crt: Crt,
}

impl Cpu {
    fn new() -> Self {
        Self {
            cycle: 0,
            x: 1,
            crt: Crt::new(),
        }
    }

    fn execute(&mut self, inst: Instruction) {
        for _ in 0..inst.num_cycles {
            let row = self.cycle / 40;
            let pos = self.cycle % 40;
            let pixel = (pos as i32 - self.x).unsigned_abs() <= 1;
            self.crt.pixels[row as usize][pos as usize] = pixel;
            self.cycle += 1;
        }
        match inst.kind {
            InstructionKind::Noop => {},
            InstructionKind::Addx => { self.x += inst.arg0; }
        }
    }
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let mut cpu = Cpu::new();
    for line in contents.lines() {
        let inst = Instruction::new(line);
        cpu.execute(inst);
    }
    cpu.crt.print();
    Ok(())
}
