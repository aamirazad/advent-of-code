package day01_secret_entrance;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

class Dial {
	public int current;
	public int count;

	public Dial(int current, int count) {
		this.current = current;
		this.count = count;
	}

	public void setDial() {
		int current = this.current;
		while (current < 0 || current > 99) {
			if (current < 0) {
				current += 100;
			} else if (current > 99) {
				current -= 100;
			}
			this.count += 1;
		}
		this.current = current;
	}
}

public class Part2 {
	public static void main(String[] args) {
		try {
			Scanner scanner = new Scanner(new File("sample.txt"));

			Dial dial = new Dial(50, 0);

			while (scanner.hasNextLine()) {
				String line = scanner.nextLine();
				char direction = line.charAt(0);
				int clicks = Integer.parseInt(line.substring(1));
				if (direction == 'L') {
					dial.current -= clicks;
					dial.setDial();
				} else if (direction == 'R') {
					dial.current += clicks;
					dial.setDial();
				}
				if (dial.current == 0) {
					dial.count += 1;
				}
			}
			scanner.close();
			System.out.println(dial.count);
		} catch (FileNotFoundException e) {
			System.out.println("File not found");
		}
	}
}
