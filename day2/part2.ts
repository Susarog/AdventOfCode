import * as fs from 'fs';

type OptionsKeyVal = {
  [key: string]: number;
};

const allOptions: OptionsKeyVal = {
  'A X': 4,
  'A Y': 8,
  'A Z': 3,
  'B X': 1,
  'B Y': 5,
  'B Z': 9,
  'C X': 7,
  'C Y': 2,
  'C Z': 6,
};
const data = fs.readFileSync('./input.txt', 'utf-8');

const score = data
  .trim()
  .split('\n')
  .reduce((acc, val) => {
    const currOption = val.split(' ');
    const elfOption = currOption[0];
    const playerOption = currOption[1];
    const filteredOptions = Object.keys(allOptions)
      .filter((key) => key.includes(elfOption))
      .reduce(
        (res: OptionsKeyVal, key: string) => (
          (res[key] = allOptions[key]), res
        ),
        {}
      );
    let value;
    if (playerOption == 'X') {
      value = Object.values(filteredOptions).find((val) => val <= 3);
    } else if (playerOption == 'Y') {
      value = Object.values(filteredOptions).find(
        (val) => val >= 4 && val <= 6
      );
    } else {
      value = Object.values(filteredOptions).find((val) => val >= 7);
    }
    return (acc += value!);
  }, 0);

console.log(score);
