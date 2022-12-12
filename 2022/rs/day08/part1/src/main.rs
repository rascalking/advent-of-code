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

    fn tree_visible(&self, x: usize, y: usize) -> bool {
        // short circuits as soon as we find one direction it's visible from
        // perimeter is always visible
        if (x == 0) || (y == 0) || (x == self.width-1) || (y == self.height-1) {
            return true;
        }

        let tree_height = self.trees[x][y];
        let mut visible;

        // left
        visible = true;
        for i in 0..x {
            if self.trees[i][y] >= tree_height {
                visible = false;
                break;
            }
        }
        if visible {
            return true;
        }

        // right
        visible = true;
        for i in x+1..self.width {
            if self.trees[i][y] >= tree_height {
                visible = false;
                break;
            }
        }
        if visible {
            return true;
        }

        // top
        visible = true;
        for j in 0..y {
            if self.trees[x][j] >= tree_height {
                visible = false;
                break;
            }
        }
        if visible {
            return true;
        }

        // bottom
        visible = true;
        for j in y+1..self.height {
            if self.trees[x][j] >= tree_height {
                visible = false;
                break;
            }
        }
        if visible {
            return true;
        }

        // if none of the directions are visible, it's not visible
        false
    }
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let forest = Forest::new(contents);
    dbg!(&forest);
    let mut visible = 0;
    for x in 0..forest.width {
        for y in 0..forest.height {
            if forest.tree_visible(x, y) {
                visible += 1;
            }
        }
    }
    dbg!(visible);
    Ok(())
}
