use std::env;
use std::fs::File;
use std::io::read_to_string;

#[derive(Debug)]
struct Forest {
    width: usize,
    height: usize,
    trees: Vec<Vec<u32>>, // [x][y], origin top left
}

impl Forest {
    fn new(text: String) -> Self {
        let mut lines = text.lines();
        let mut trees: Vec<Vec<u32>> = lines.next().unwrap().chars().map(|c| vec![c.to_digit(10).unwrap()]).collect();
        for line in lines {
            for (x, c) in line.chars().enumerate() {
                trees[x].push(c.to_digit(10).unwrap());
            }
        }
        Forest {
            width: trees.len(),
            height: trees[0].len(),
            trees,
        }
    }

    fn scenic_score(&self, x: usize, y: usize) -> u32 {
        // perimeter is always 0, since one side has 0 viewing distance
        if (x == 0) || (y == 0) || (x == self.width-1) || (y == self.height-1) {
            return 0;
        }

        let tree_height = self.trees[x][y];

        // left
        let mut left = x;
        for i in (0..x).rev() {
            if self.trees[i][y] >= tree_height {
                left = x - i;
                break;
            }
        }

        // right
        let mut right = self.width - 1 - x;
        for i in x+1..self.width {
            if self.trees[i][y] >= tree_height {
                right = i - x;
                break;
            }
        }

        // top
        let mut top = y;
        for j in (0..y).rev() {
            if self.trees[x][j] >= tree_height {
                top = y - j;
                break;
            }
        }

        // bottom
        let mut bottom = self.height - 1 - y;
        for j in y+1..self.height {
            if self.trees[x][j] >= tree_height {
                bottom = j - y;
                break;
            }
        }

        let score: u32 = (left * right * top * bottom).try_into().unwrap();
        println!("({}, {}) left: {}, right: {}, top: {}, bottom: {}, score: {}", x, y, left, right, top, bottom, score);
        score
    }
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let forest = Forest::new(contents);
    //dbg!(&forest);
    let mut highest = 0;
    for x in 0..forest.width {
        for y in 0..forest.height {
            let score = forest.scenic_score(x, y);
            if score > highest {
                highest = score;
            }
        }
    }
    dbg!(highest);
    Ok(())
}
