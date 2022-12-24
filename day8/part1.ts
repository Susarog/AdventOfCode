import * as fs from 'fs';
const input = fs.readFileSync('./input.txt', 'utf-8').trim().split('\n');

const twoDTree: string[][] = [];
for (let i = 0; i < input.length; i++) {
  twoDTree.push(input[i].trim().split(''));
}

let visibleTreeCount = 2 * (twoDTree.length - 1) + 2 * (twoDTree[0].length - 1);
for (let i = 1; i < twoDTree.length - 1; i++) {
  for (let j = 1; j < twoDTree[0].length - 1; j++) {
    const val = twoDTree[i][j];
    const z = 1;
    const w = 1;
    if (
      checkColumnIndexes(i, j, z, val, w) ||
      checkRowIndexes(i, j, z, val, w) ||
      checkColumnIndexes(i, j, -z, val, -w) ||
      checkRowIndexes(i, j, -z, val, -w)
    ) {
      continue;
    }
  }
}

console.log(visibleTreeCount);

function checkColumnIndexes(
  i: number,
  j: number,
  z: number,
  val: string,
  w: number
) {
  while (true) {
    if (i + z < 0 || i + z > twoDTree.length - 1) {
      visibleTreeCount++;
      return true;
    }
    const currentVal = twoDTree[i + z][j];
    z += w;
    if (Number(currentVal) >= Number(val)) {
      return false;
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
  while (true) {
    if (j + z < 0 || j + z > twoDTree[0].length - 1) {
      visibleTreeCount++;
      return true;
    }
    const currentVal = twoDTree[i][j + z];
    z += w;

    if (Number(currentVal) >= Number(val)) {
      return false;
    }
  }
}
