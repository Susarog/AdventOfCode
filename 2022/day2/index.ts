import * as fs from 'fs';

type OptionsKeyVal = {
  [key: string]: number;
};

const allOptions: OptionsKeyVal = {
  'A X': 4, //rock rock draw
  'A Y': 8, //rock paper W
  'A Z': 3, // rock scissors L
  'B X': 1, // paper rock L
  'B Y': 5, // paper paper draw
  'B Z': 9, // paper scissors W
  'C X': 7, // scissors rock W
  'C Y': 2, // scissors paper L
  'C Z': 6, // scissors scissors draw
};

const data = fs.readFileSync('./input.txt', 'utf-8');

const score = data
  .trim()
  .split('\n')
  .reduce((acc, val) => {
    return (acc += allOptions[val]);
  }, 0);

console.log(score);
