import * as fs from 'fs';
import * as util from 'util';

const input = fs.readFileSync('./input.txt', 'utf-8').trim().split('\n');
const map: { [key: string]: any } = {};
const currDirectory: string[] = [];
let currMap: { [key: string]: any } = map;
let output = 0;
function sumDirectory() {
  let sum = 0;
  for (const val in currMap) {
    sum += currMap[val];
  }
  if (sum < 100000) {
    output += sum;
  }
  const lastDirectory = currDirectory.pop();

  if (!lastDirectory) {
    throw 'input aint correct or smth';
  } else {
    currMap = map;
    for (let i = 0; i < currDirectory.length; i++) {
      currMap = currMap[currDirectory[i]];
    }
    currMap[lastDirectory] = sum;
  }
}
input.forEach((currInput) => {
  if (currInput === '$ cd ..') {
    sumDirectory();
  } else if (currInput.includes('$ cd')) {
    currDirectory.push(currInput.split(' ')[2]);
    currMap = map;
    for (let i = 0; i < currDirectory.length; i++) {
      const temp = currMap[currDirectory[i]];
      if (!temp) {
        currMap[currDirectory[i]] = {};
        currMap = currMap[currDirectory[i]];
        break;
      }
      currMap = currMap[currDirectory[i]];
    }
  } else if (!isNaN(Number(currInput.split(' ')[0]))) {
    const temp = currInput.split(' ');
    currMap[temp[1]] = Number(temp[0]);
  }
});

while (currDirectory.length != 0) {
  sumDirectory();
}

console.log(output);
