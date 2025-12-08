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

	public void setDial(int clicks) {
		int location = this.current + clicks;
		System.out.println("MOVING " + clicks + " New location: " + location);
		while (location < 0 || location > 99) {
			// Passed through zero
			if (this.current !== 0) {
				System.out.println("Passed zero because of " + clicks);
				this.count += 1;
			}
			if (location < 0) {
				location += 100;
			} else if (location > 99) {
				location -= 100;
			}
		}
		this.current = location;
		if (location == 0) {
			System.out.println("Landed on zero");
			this.count += 1;
		}
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
					dial.setDial(-clicks);
				} else if (direction == 'R') {
					dial.setDial(clicks);
				}
			}
			scanner.close();
			System.out.println(dial.count);
		} catch (FileNotFoundException e) {
			System.out.println("File not found");
		}
	}
}
