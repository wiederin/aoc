use std::{str::FromStr};

use anyhow::Result;

#[derive(Debug)]
struct RoundScore {
    value: usize,
}
#[derive(Debug)]
struct RoundScore1 {
    value: usize,
}

const WIN_LOSE: [usize; 3] = [3, 6, 0];
const WIN_LOSE_2: [usize; 3] = [0, 3, 6];
const CHOICE_VALUE: [usize; 3] = [3, 1, 2];

fn to_number(c: &str) -> usize {
    return match c {
        "A" => 0,
        "B" => 2,
        "C" => 1,
        "X" => 1,
        "Y" => 2,
        "Z" => 3,
        _ => unreachable!("invalid move")
    }
}

fn to_number_2(c: &str) -> usize {
    return match c {
        "A" => 0,
        "B" => 1,
        "C" => 2,
        "X" => 0,
        "Y" => 1,
        "Z" => 2,
        _ => unreachable!("invalid move")
    }
}

impl FromStr for RoundScore {
    type Err = anyhow::Error;
    fn from_str(s: &str) -> Result<Self> {
        let (o, p) = match s.split_once(" ") {
            Some((o, p)) => (o, p),
            None => return Err(anyhow::anyhow!("invalid input")),
        };
        let opp_move = to_number(o);
        let my_move = to_number(p);
        let score = my_move + WIN_LOSE[
            (2 + opp_move + my_move) % WIN_LOSE.len()
        ];
        return Ok(RoundScore { value: score });
    }
}

impl FromStr for RoundScore1 {
    type Err = anyhow::Error;
    fn from_str(s: &str) -> Result<Self> {
        let (o, p) = match s.split_once(" ") {
            Some((o, p)) => (o, p),
            None => return Err(anyhow::anyhow!("invalid input")),
        };
        let opp_move = to_number_2(o);
        let my_move = to_number_2(p);
        let score = WIN_LOSE_2[my_move] + CHOICE_VALUE[
            (opp_move + my_move) % CHOICE_VALUE.len()
        ];
        return Ok(RoundScore1 { value: score });
    }
}
fn main() -> Result<()> {
    let values: usize = include_str!("input2_1.prod")
        .lines()
        .flat_map(|x| x.parse::<RoundScore>())
        .map(|x| x.value)
        .sum();

    println!("Score part 1: {:?}", values);

    let values1: usize = include_str!("input2_1.prod")
        .lines()
        .flat_map(|x| x.parse::<RoundScore1>())
        .map(|x| x.value)
        .sum();

    println!("Score part 2: {:?}", values1);
    return Ok(());
}
