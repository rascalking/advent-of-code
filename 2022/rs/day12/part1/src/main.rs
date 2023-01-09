use std::cmp::Reverse;
use std::collections::HashMap;
use std::collections::HashSet;
use std::convert::TryFrom;
use std::env;
use std::fs::File;
use std::io::read_to_string;
use std::str::FromStr;
use std::string::ToString;

use priority_queue::PriorityQueue;

#[derive(Debug, Copy, Clone, PartialEq, Eq)]
struct Height(u8);

impl TryFrom<char> for Height {
    type Error = ();

    fn try_from(c: char) -> Result<Self, Self::Error> {
        match c {
            'a'..='z' => Ok(Height((c as u8) - 96)),
            'S' => Ok(Height(0)),
            'E' => Ok(Height(27)),
            _ => Err(()),
        }
    }
}

impl From<Height> for char {
    fn from(val: Height) -> Self {
        match val {
            Height(0) => 'S',
            Height(27) => 'E',
            Height(h) => (h + 96) as char,
        }
    }
}

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
struct Coord {
    x: usize,
    y: usize,
}

#[derive(Debug)]
struct Grid {
    start: Coord,
    end: Coord,
    heights: Vec<Vec<Height>>,
}

impl FromStr for Grid {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut heights: Vec<Vec<Height>> = Vec::new();
        let mut start = Coord {
            x: usize::MAX,
            y: usize::MAX,
        };
        let mut end = Coord {
            x: usize::MAX,
            y: usize::MAX,
        };
        for (y, line) in s.lines().enumerate() {
            let mut row: Vec<Height> = Vec::new();
            for (x, c) in line.chars().enumerate() {
                let height = Height::try_from(c)?;
                match height {
                    Height(0) => {
                        start = Coord { x, y };
                    }
                    Height(27) => {
                        end = Coord { x, y };
                    }
                    _ => {}
                }
                row.push(height);
            }
            heights.push(row);
        }
        Ok(Grid {
            start,
            end,
            heights,
        })
    }
}

impl ToString for Grid {
    fn to_string(&self) -> String {
        let y_bound = self.heights.len();
        let x_bound = self.heights[0].len();
        let mut buf = String::with_capacity((x_bound + 1) * y_bound);
        for y in 0..y_bound {
            for x in 0..x_bound {
                let c: char = self.heights[y][x].into();
                buf.push(c);
            }
            buf.push('\n');
        }
        buf
    }
}

impl Grid {
    fn height_at(&self, c: Coord) -> Height {
        self.heights[c.y][c.x]
    }

    fn heuristic(&self, a: Coord, b: Coord) -> usize {
        ((a.x.abs_diff(b.x).pow(2) + a.y.abs_diff(b.y).pow(2)) as f64).sqrt() as usize
    }

    // implements A* as explained at https://en.wikipedia.org/wiki/A*_search_algorithm#Pseudocode
    fn path(&self, s: Coord, d: Coord) -> Vec<Coord> {
        let mut open_set = HashSet::new();
        open_set.insert(s);

        let mut came_from: HashMap<Coord, Coord> = HashMap::new();

        let mut g_score = HashMap::new(); // defaultdict(Infinity)
        g_score.insert(s, 0usize);

        let mut f_score = HashMap::new(); // defaultdict(Infinity)
        f_score.insert(s, self.heuristic(s, d));

        while !open_set.is_empty() {
            let mut open = Vec::from_iter(open_set.clone());
            open.sort_by_key(|c| Reverse(f_score[c]));
            let mut current = open.pop().unwrap();

            if current == d {
                let mut path: Vec<Coord> = Vec::new();
                while current != s {
                    current = came_from[&current];
                    path.push(current);
                }
                path.reverse();
                return path;
            }

            open_set.remove(&current);
            g_score.entry(current).or_insert_with(|| usize::MAX);
            f_score.entry(current).or_insert_with(|| usize::MAX);
            println!("considering {:?}, g_score {:?}, f_score {:?}", current, g_score[&current], f_score[&current]);
            for neighbor in self.valid_neighbors(current) {
                g_score.entry(neighbor).or_insert_with(|| usize::MAX);
                let tmp_g_score = g_score[&current].saturating_add(1);
                println!("\tneighbor {:?} current g_score: {:?}, new g_score: {:?}", neighbor, g_score[&neighbor], tmp_g_score);
                if tmp_g_score < g_score[&neighbor] {
                    came_from.insert(neighbor, current);
                    g_score.insert(neighbor, tmp_g_score);
                    f_score.insert(neighbor, tmp_g_score + self.heuristic(neighbor, d));
                    println!("\t\tnew g_score={:?}, f_score={:?}", tmp_g_score, f_score[&neighbor]);
                    if !open_set.contains(&neighbor) {
                        open_set.insert(neighbor);

                    }
                }
            }
        }
        dbg!(came_from);
        dbg!(g_score);
        dbg!(f_score);
        vec![]
    }

    fn solve(&self) -> Vec<Coord> {
        self.path(self.start, self.end)
    }

    fn valid_neighbors(&self, c: Coord) -> Vec<Coord> {
        let mut valid = Vec::new();
        for (x_mod, y_mod) in [(-1, 0), (0, -1), (1, 0), (0, 1)] {
            if let (Some(x), Some(y)) =
                (c.x.checked_add_signed(x_mod), c.y.checked_add_signed(y_mod))
            {
                let n = Coord { x, y };
                if (y < self.heights.len())
                    && (x < self.heights[0].len())
                    && (self.height_at(n).0.abs_diff(self.height_at(c).0) <= 1)
                {
                    valid.push(n);
                }
            }
        }
        valid
    }
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let grid = Grid::from_str(&contents).unwrap();
    let path: Vec<Coord> = grid.solve();
    println!("{:?} -> {:?} = {}", grid.start, grid.end, path.len());
    Ok(())
}
