import * as fs from 'fs';
const input = fs.readFileSync('./input.txt', 'utf-8').trim().split('\n');

const twoDTree: string[][] = [];
for (let i = 0; i < input.length; i++) {
  twoDTree.push(input[i].trim().split(''));
}
let max = 0;
for (let i = 1; i < twoDTree.length - 1; i++) {
  for (let j = 1; j < twoDTree[0].length - 1; j++) {
    const val = twoDTree[i][j];
    const z = 1;
    const w = 1;
    const temp =
      checkColumnIndexes(i, j, z, val, w) *
      checkRowIndexes(i, j, z, val, w) *
      checkColumnIndexes(i, j, -z, val, -w) *
      checkRowIndexes(i, j, -z, val, -w);
    if (max < temp) {
      max = temp;
    }
  }
}
console.log(max);

function checkColumnIndexes(
  i: number,
  j: number,
  z: number,
  val: string,
  w: number
) {
  let temp = 0;
  while (true) {
    if (i + z < 0 || i + z > twoDTree.length - 1) {
      return temp;
    }
    const currentVal = twoDTree[i + z][j];
    temp++;
    z += w;
    if (Number(currentVal) >= Number(val)) {
      return temp;
    }
  }
}
function checkRowIndexes(
  i: number,
  j: number,
  z: number,
  val: string,
  w: number
) {
  let temp = 0;
  while (true) {
    if (j + z < 0 || j + z > twoDTree[0].length - 1) {
      return temp;
    }
    const currentVal = twoDTree[i][j + z];
    temp++;
    z += w;
    if (Number(currentVal) >= Number(val)) {
      return temp;
    }
  }
}
