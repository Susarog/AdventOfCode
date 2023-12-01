import * as fs from 'fs';

const data = fs.readFileSync('./input.txt', 'utf-8');

const answer = data
  .trim()
  .split('\n')
  .reduce((acc, pair) => {
    const currPair = pair.split(',');
    const firstPair = currPair[0].split('-');
    const secondPair = currPair[1].split('-');
    const firstPairVal = Number(firstPair[0]);
    const firstPairVal2 = Number(firstPair[1]);
    const secondPairVal = Number(secondPair[0]);
    const secondPairVal2 = Number(secondPair[1]);
    const firstRange = firstPairVal2 - firstPairVal + 1;
    const secondRange = secondPairVal2 - secondPairVal + 1;
    if (secondRange > firstRange) {
      if (
        (firstPairVal >= secondPairVal && firstPairVal <= secondPairVal2) ||
        (firstPairVal2 >= secondPairVal && firstPairVal2 <= secondPairVal2)
      ) {
        return (acc += 1);
      }
    } else {
      if (
        (secondPairVal >= firstPairVal && secondPairVal <= firstPairVal2) ||
        (secondPairVal2 >= firstPairVal && secondPairVal2 <= firstPairVal2)
      ) {
        return (acc += 1);
      }
    } 
    return acc;
  }, 0);

console.log(answer);
