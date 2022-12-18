import * as fs from 'fs';

const data = fs.readFileSync('./input.txt', 'utf-8');

const score = data.trim().split('\n');
let priorityValue = 0;
for (let i = 0; i < score.length; i += 3) {
  const groupRuckSack = {
    score0: score[i],
    score1: score[i + 1],
    score2: score[i + 2],
  };
  const temp: { [key: string]: string } = {};
  const temp2: { [key: string]: string } = {};
  for (let i = 0; i < groupRuckSack.score0.length; i++) {
    temp[groupRuckSack.score0[i]] = groupRuckSack.score0[i];
  }
  for (let i = 0; i < groupRuckSack.score1.length; i++) {
    temp2[groupRuckSack.score1[i]] = groupRuckSack.score1[i];
  }
  for (let i = 0; i < groupRuckSack.score2.length; i++) {
    if (!temp[groupRuckSack.score2[i]] || !temp2[groupRuckSack.score2[i]]) {
      continue;
    }
    if (groupRuckSack.score2[i].charCodeAt(0) >= 97) {
      priorityValue += groupRuckSack.score2[i].charCodeAt(0) - 96;
    } else {
      priorityValue += groupRuckSack.score2[i].charCodeAt(0) - 38;
    }
    break;
  }
}

console.log(priorityValue);
