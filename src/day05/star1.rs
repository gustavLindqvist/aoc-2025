pub fn star1(input: &str) -> String {
    let (p1, p2) = input.split_once("\n\n").unwrap();

    let v = p1
        .lines()
        .map(|line| line.split_once('-').unwrap())
        .map(|(a, b)| (a.parse::<i64>().unwrap(), b.parse::<i64>().unwrap()))
        .collect::<Vec<_>>();
    p2.lines()
        .map(|l| l.parse().unwrap())
        .fold(0, |acc, n: i64| {
            for (start, stop) in &v {
                if n >= *start && n <= *stop {
                    return acc + 1;
                }
            }
            acc
        })
        .to_string()
}
