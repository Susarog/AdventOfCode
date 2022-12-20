import * as fs from 'fs';

const data = fs.readFileSync('./input.txt', 'utf-8');
const arrayOfStacks: Array<Array<string>> = [];
for (let i = 0; i < 10; i++) {
  arrayOfStacks.push([]);
}
arrayOfStacks[1].push('Q', 'M', 'G', 'C', 'L');
arrayOfStacks[2].push('R', 'D', 'L', 'C', 'T', 'F', 'H', 'G');
arrayOfStacks[3].push('V', 'J', 'F', 'N', 'M', 'T', 'W', 'R');
arrayOfStacks[4].push('J', 'F', 'D', 'V', 'Q', 'P');
arrayOfStacks[5].push('N', 'F', 'M', 'S', 'L', 'B', 'T');
arrayOfStacks[6].push('R', 'N', 'V', 'H', 'C', 'D', 'P');
arrayOfStacks[7].push('H', 'C', 'T');
arrayOfStacks[8].push('G', 'S', 'J', 'V', 'Z', 'N', 'H', 'P');
arrayOfStacks[9].push('Z', 'F', 'H', 'G');

data
  .trim()
  .split('\n')
  .forEach((input) => {
    const arrInput = input.split(' ');
    const amountToMove = Number(arrInput[1]);
    const from = Number(arrInput[3]);
    const to = Number(arrInput[5]);
    for (let i = 0; i < amountToMove; i++) {
      const cargo = arrayOfStacks[from].pop();
      if (cargo) {
        arrayOfStacks[to].push(cargo);
      }
    }
  });

for (let i = 0; i < arrayOfStacks.length; i++) {
  console.log(arrayOfStacks[i].pop());
}
