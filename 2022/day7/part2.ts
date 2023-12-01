import * as fs from 'fs';

const input = fs.readFileSync('./input.txt', 'utf-8').trim().split('\n');
const map: { [key: string]: any } = {};
const currDirectory: string[] = [];
let currMap: { [key: string]: any } = map;
let output = 0;
function sumDirectory() {
  let sum = 0;
  for (const val in currMap) {
    if (typeof currMap[val] === 'object') {
      sum += currMap[val].sum;
      continue;
    }
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
    currMap[lastDirectory] = { sum: sum, ...currMap[lastDirectory] };
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
const unusedSpace = 70000000 - map['/'].sum;
const neededSizeForSpace = 30000000 - unusedSpace;
const smallest = { smallest: map['/'].sum };

const iterate = (obj: any, smallest: { [key: string]: number }) => {
  Object.keys(obj).forEach((key) => {
    if (
      obj[key].sum < smallest.smallest &&
      obj[key].sum >= neededSizeForSpace
    ) {
      smallest.smallest = obj[key].sum;
    }
    if (typeof obj[key] === 'object' && obj[key] !== null) {
      iterate(obj[key], smallest);
    }
  });
};
iterate(map, smallest);
console.log(smallest.smallest);
