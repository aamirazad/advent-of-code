use std::fs;

const FILE: &str = "sample.txt";

fn main() {
    let input = fs::read_to_string(FILE).expect("Unable to read file");
}
