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
    if (firstRange > secondRange) {
      if (!(secondPairVal2 <= firstPairVal2 && secondPairVal >= firstPairVal)) {
        return acc;
      }
    } else if (firstRange < secondRange) {
      if (!(firstPairVal2 <= secondPairVal2 && firstPairVal >= secondPairVal)) {
        return acc;
      }
    } else {
      if (secondPairVal2 != firstPairVal2 && secondPairVal != firstPairVal) {
        return acc;
      }
    }
    return (acc += 1);
  }, 0);

console.log(answer);
