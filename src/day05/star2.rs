use std::cmp::max;

pub fn star2(input: &str) -> String {
    let input = input.split_once("\n\n").unwrap().0;
    let mut v = input
        .lines()
        .map(|line| line.split_once('-').unwrap())
        .map(|(a, b)| (a.parse::<i64>().unwrap(), b.parse::<i64>().unwrap()))
        .collect::<Vec<_>>();
    v.sort_unstable();
    v.into_iter()
        .fold((0, 0), |(count, end), (start, stop)| {
            let mut diff = max(end, stop) - max(end, start);
            if end > start {
                diff += 1;
            }
            (count + diff, max(end, stop))
        })
        .0
        .to_string()
}
