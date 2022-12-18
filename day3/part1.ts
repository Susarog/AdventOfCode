import * as fs from 'fs';

const data = fs.readFileSync('./input.txt', 'utf-8');

const score = data
  .trim()
  .split('\n')
  .reduce((acc: number, ruckSack: string): number => {
    const firstRuckSack = ruckSack.substring(0, ruckSack.length / 2);
    const secondRuckSack = ruckSack.substring(
      ruckSack.length / 2,
      ruckSack.length
    );
    const temp: { [key: string]: string } = {};
    let priorityValue: number;
    for (let i = 0; i < firstRuckSack.length; i++) {
      temp[firstRuckSack[i]] = firstRuckSack[i];
    }
    for (let i = 0; i < secondRuckSack.length; i++) {
      if (!temp[secondRuckSack[i]]) {
        continue;
      }
      if (secondRuckSack[i].charCodeAt(0) >= 97) {
        priorityValue = secondRuckSack[i].charCodeAt(0) - 96;
      } else {
        priorityValue = secondRuckSack[i].charCodeAt(0) - 38;
      }
      break;
    }
    return (acc += priorityValue!);
  }, 0);
console.log(score);
