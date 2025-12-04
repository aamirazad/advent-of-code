import * as fs from 'fs';

const filePath: string = 'sample.txt';
const content: string = fs.readFileSync(filePath, 'utf-8');
const lines: string[] = content.split(/\r?\n/);

for (const line of lines) {
  console.log(line.trim());
}