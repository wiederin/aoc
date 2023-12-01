use anyhow::Result;

fn main() -> Result<()> {
    let mut max: Vec<usize> = include_str!("./input1_1.prod")
        .split("\n\n")
        .map(|x| {
            return x
                .lines()
                .flat_map(str::parse::<usize>)
                .sum();
        })
        .collect();
    
    max.sort_by(|a, b| b.cmp(a));

    println!("max: {:?}", max
        .into_iter()
        .take(3)
        .sum::<usize>()
    );

    return Ok(());
}
