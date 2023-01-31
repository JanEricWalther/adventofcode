const input = Deno.readTextFileSync(Deno.args[0]).trim();
const START_POS = { x: 500, y: 0 } as const;
const tiles = {
  AIR: 0,
  ROCK: 1,
  SAND: 2,
  START: 10,
} as const;
let heighestY = 0;

const instructions = parse(input);
// Part 1 is off by 1 for some reason... 
console.log(`Part One: ${sandUnits(buildGrid(instructions)) - 1}`);
console.log(`Part Two: ${sandUnits(buildGrid(instructions, true))}`);

function sandUnits(grid: number[][]) {
  let unitsOfSand = 0;
  let keepGoing = true;

  while (keepGoing) {
    let { x, y } = START_POS;
    unitsOfSand++;
    while (true) {
      if (
        grid[x][y + 1] > 0 && grid[x - 1][y + 1] > 0 && grid[x + 1][y + 1] > 0
      ) {
        if (x === START_POS.x && y === START_POS.y) {
          keepGoing = false;
          break;
        }
        grid[x][y] = tiles.SAND;
        break;
      } else if (grid[x][y + 1] > 0) {
        if (grid[x - 1][y + 1] > 0) {
          x++;
        } else {
          x--;
        }
      }
      y++;
      if (y >= 600) {
        keepGoing = false;
        break;
      }
    }
  }
  return unitsOfSand;
}

function buildGrid(paths: { x: number; y: number }[][], partTwo = false) {
  const grid: number[][] = Array(1000).fill(null).map((_) =>
    Array(1000).fill(0)
  );
  grid[START_POS.x][START_POS.y] = tiles.START;
  for (const path of paths) {
    let last;
    for (const node of path) {
      grid[node.x][node.y] = tiles.ROCK;
      if (!last) {
        last = node;
        continue;
      }
      if (node.x !== last.x) {
        for (let i = last.x; i !== node.x; i += Math.sign(node.x - last.x)) {
          grid[i][node.y] = tiles.ROCK;
        }
      } else if (node.y !== last.y) {
        for (let i = last.y; i !== node.y; i += Math.sign(node.y - last.y)) {
          grid[node.x][i] = tiles.ROCK;
        }
      }
      last = node;
    }

    // build floor
    if (partTwo) {
      for (const i of Array(1000).keys()) {
        grid[i][heighestY + 2] = tiles.ROCK;
      }
    }
  }
  return grid;
}

function parse(input: string) {
  const paths: { x: number; y: number }[][] = [];
  for (const line of input.split("\n")) {
    const path: { x: number; y: number }[] = [];
    for (const point of line.trim().split("->")) {
      const [x, y] = point.trim().split(",");
      path.push({ x: +x, y: +y });
      heighestY = Math.max(heighestY, +y);
    }
    paths.push(path);
  }
  return paths;
}
