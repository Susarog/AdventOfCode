import * as fs from 'fs';
import * as readline from 'readline';

let calorie = 0;

const allCalories: Array<number> = [];
async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });
  for await (const line of rl) {
    const currLine = Number(line);
    if (line.length === 0) {
      allCalories.push(calorie);
      calorie = 0;
    } else {
      calorie += currLine;
    }
  }
  allCalories.sort((a, b) => b - a);
  console.log(allCalories[0] + allCalories[1] + allCalories[2]);
}

processLineByLine();
