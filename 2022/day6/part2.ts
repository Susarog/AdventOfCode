import * as fs from 'fs';

const input = fs.readFileSync('./input.txt', 'utf-8').trim();

for (let i = 0; i < input.length; i++) {
  const fourUniqueChar = input.substring(i, i + 14);
  const tempObj: { [key: string]: string } = {};
  let hasDuplicate = false;
  for (let j = 0; j < fourUniqueChar.length; j++) {
    if (!tempObj[fourUniqueChar[j]]) {
      tempObj[fourUniqueChar[j]] = fourUniqueChar[j];
    } else {
      hasDuplicate = true;
      break;
    }
  }
  if (hasDuplicate) {
    continue;
  } else {
    console.log(i + 14);
    break;
  }
}
