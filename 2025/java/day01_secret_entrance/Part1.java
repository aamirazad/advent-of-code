package day01_secret_entrance;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Part1 {
	public static void main(String[] args) {
		try {
			Scanner scanner = new Scanner(new File("input.txt"));

			Integer current = 50;
			Integer timesHitsZero = 0;

			while (scanner.hasNextLine()) {
				String line = scanner.nextLine();
				Character direction = line.charAt(0);
				Integer clicks = Integer.parseInt(line.substring(1));
				if (direction == 'L') {
					current -= clicks;
					current = wraparound(current);
				} else if (direction == 'R') {
					current += clicks;
					current = wraparound(current);
				}
				if (current == 0) {
					timesHitsZero += 1;
				}
			}
			scanner.close();
			System.out.println(timesHitsZero);
		} catch (FileNotFoundException e) {
			System.out.println("File not found");
		}
	}

	static Integer wraparound(Integer current) {
		if (current >= 0 && current <= 99) {
			return current;
		}

		if (current < 0) {
			current += 100;
		} else if (current > 99) {
			current -= 100;
		}
		return wraparound(current);
	}
}
