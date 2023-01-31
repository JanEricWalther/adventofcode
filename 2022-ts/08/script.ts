import * as Colors from "https://deno.land/std/fmt/colors.ts";
const input = Deno.readTextFileSync(Deno.args[0]).trim();

const treeGrid = parseMatrix(input);
console.log(`Part 1: ${visibleFromTheOutside(treeGrid)}`);
console.log(`Part 2: ${scenicScore(treeGrid)}`);

function visibleFromTheOutside(grid: number[][]): number {
  const visibilityMatrix: boolean[][] = Array(grid.length).fill(null).map((_) =>
    Array(grid[0].length).fill(false)
  );

  // LEFT
  for (let y = 0; y < grid.length; y++) {
    let max = -1;
    for (let x = 0; x < grid[0].length; x++) {
      const cell = grid[y][x];
      if (cell > max) {
        max = cell;
        visibilityMatrix[y][x] = true;
        if (cell >= 9) {
          break;
        }
      }
    }
  }

  // RIGHT
  for (let y = 0; y < grid.length; y++) {
    let max = -1;
    for (let x = grid[0].length - 1; x >= 0; x--) {
      const cell = grid[y][x];
      if (cell > max) {
        max = cell;
        visibilityMatrix[y][x] = true;
        if (cell >= 9) {
          break;
        }
      }
    }
  }

  // TOP
  for (let x = 0; x < grid[0].length; x++) {
    let max = -1;
    for (let y = 0; y < grid.length; y++) {
      const cell = grid[y][x];
      if (cell > max) {
        max = cell;
        visibilityMatrix[y][x] = true;
        if (cell >= 9) {
          break;
        }
      }
    }
  }

  // BOTTOM
  for (let x = 0; x < grid[0].length; x++) {
    let max = -1;
    for (let y = grid.length - 1; y >= 0; y--) {
      const cell = grid[y][x];
      if (cell > max) {
        max = cell;
        visibilityMatrix[y][x] = true;
        if (cell >= 9) {
          break;
        }
      }
    }
  }
  // print(treeGrid, visibilityMatrix);
  return visibilityMatrix.flat().filter((cell) => cell).length;
}

function scenicScore(grid: number[][]): number {
  let maxScore = 0;

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[0].length; x++) {
      const cell = grid[y][x];

      let leftScore = 0;
      let i = 1;
      while (grid[y][x - i] != null) {
        leftScore++;
        if (cell <= grid[y][x - i]) {
          break;
        }
        i++;
      }

      let rightScore = 0;
      i = 1;
      while (grid[y][x + i] != null) {
        rightScore++;
        if (cell <= grid[y][x + i]) {
          break;
        }
        i++;
      }

      let topScore = 0;
      i = 1;
      while (grid[y - i]?.[x] != null) {
        topScore++;
        if (cell <= grid[y - i][x]) {
          break;
        }
        i++;
      }

      let bottomScore = 0;
      i = 1;
      while (grid[y + i]?.[x] != null) {
        bottomScore++;
        if (cell <= grid[y + i][x]) {
          break;
        }
        i++;
      }

      const score = leftScore * rightScore * topScore * bottomScore;
      maxScore = Math.max(score, maxScore);
    }
  }
  return maxScore;
}

function parseMatrix(input: string): number[][] {
  const matrix: number[][] = [];
  for (const line of input.split("\n")) {
    const row = [];
    for (const char of line.trim().split("")) {
      row.push(+char);
    }
    matrix.push(row);
  }
  return matrix;
}

/** for debugging only
 *  colors in the visible Trees
 *
 * @param treeMatrix - given grid of Trees
 * @param visibilityMatrix - whether or not a given Tree is visible from the outside
 */
function print(treeMatrix: number[][], visibilityMatrix: boolean[][]) {
  const encoder = new TextEncoder();

  for (let y = 0; y < treeMatrix.length; y++) {
    for (let x = 0; x < treeMatrix[0].length; x++) {
      const c = String(treeMatrix[y][x]);
      if (visibilityMatrix[y][x]) {
        Deno.stdout.writeSync(encoder.encode(Colors.green(c)));
      } else {
        Deno.stdout.writeSync(encoder.encode(Colors.gray(c)));
      }
    }
    Deno.stdout.writeSync(encoder.encode("\n"));
  }
}
