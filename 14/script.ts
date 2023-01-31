const input = Deno.readTextFileSync(Deno.args[0]).trim();
const START_POS = { x: 500, y: 0 } as const;
const tiles = {
  AIR: 0,
  ROCK: 1,
  SAND: 2,
  START: 10,
} as const;

const grid = buildGrid(parse(input));
console.log(sandUnits(grid));

function sandUnits(grid: number[][]) {
  let unitsOfSand = 0;
  let hitBottom = false;
  while (!hitBottom) {
    let { x, y } = START_POS;
    unitsOfSand++;
    while (true) {
      if (
        grid[x][y + 1] > 0 && grid[x - 1][y + 1] > 0 && grid[x + 1][y + 1] > 0
      ) {
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
        hitBottom = true;
        break;
      }
    }
  }
  return unitsOfSand - 1;
}

function buildGrid(paths: { x: number; y: number }[][]) {
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
    }
    paths.push(path);
  }
  return paths;
}
