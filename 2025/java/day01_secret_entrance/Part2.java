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

	public int setDial(int clicks) {
		int newLocation = this.current + clicks;
		if (newLocation < 0) {
			newLocation += 100;
			this.count += 1;
		} else if (newLocation > 99) {
			newLocation -= 100;
			this.count += 1;
		} else if (newLocation == 0) {
			this.count += 1;
		} else if (clicks > 100 || clicks < -100) {
			this.count += Math.floor(Math.abs(clicks) / 100);
		} else {
			return 
		}
		return setDial(0)
	}

	/*
	 * Pseudocode
	 * Set the new location
	 * If the number is (newLocation < 0 || newLocation > 99):
	 * if the number is less than zero, add 100 (else subtract 100), return
	 * setDial(newNumber)
	 * Else if the number is zero, increase counter, return number
	 * Else if the number is greater than 100 past the old number, use mod to find
	 * the number times passed
	 */
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
