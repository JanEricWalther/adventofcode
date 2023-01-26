class Point {
  #x: number;
  #y: number;

  constructor(x: number, y: number) {
    this.#x = x;
    this.#y = y;
  }

  get x() {
    return this.#x;
  }

  get y() {
    return this.#y;
  }

  toJSON() {
    return {
      x: this.#x,
      y: this.#y,
    };
  }
}

const input = Deno.readTextFileSync(Deno.args[0]).trim();

console.log("Fewest Steps:", solve(parse(input)));

function solve(input: string[][]) {
  const map: number[][] = Array(input.length).fill(null).map((_) =>
    Array(input[0].length).fill(false)
  );
  let S: Point;
  let E: Point;
  for (let y = 0; y < input.length; y++) {
    for (let x = 0; x < input[y].length; x++) {
      const cell = input[y][x];
      if (cell === "S") {
        S = new Point(x, y);
        // Starting position S has elevation a
        input[y][x] = "a";
      } else if (cell === "E") {
        E = new Point(x, y);
        // Destination E has elevation z
        input[y][x] = "z";
      }
      // While we are looping, re-encode chars to ints, to make our elevation comparisons easier
      map[y][x] = input[y][x].charCodeAt(0);
    }
  }
  const path = getShortestPath(S!, E!);
  return path.length;

  function getNeighbors(p: Point) {
    return [
      new Point(p.x, p.y - 1),
      new Point(p.x - 1, p.y),
      new Point(p.x + 1, p.y),
      new Point(p.x, p.y + 1),
    ].filter((coord) => typeof map[coord.y]?.[coord.x] !== "undefined");
  }

  function buildFrontier(from: Point) {
    const frontier: Point[] = [];
    frontier.push(from);

    const came_from: Map<string, string> = new Map();
    came_from.set(JSON.stringify(from), "start");

    while (frontier.length > 0) {
      const current = frontier.shift() ?? new Point(0, 0);
      const current_val = map[current.y][current.x];

      const neighbors = getNeighbors(current);
      for (const next of neighbors) {
        const next_cell = map[next.y][next.x];
        const next_id = JSON.stringify(next);

        if (next_cell - current_val > 1 || came_from.has(next_id)) {
          continue;
        }

        const current_id = JSON.stringify(current);
        frontier.push(next);
        came_from.set(next_id, current_id);
      }
    }
    return came_from;
  }

  function getShortestPath(from: Point, to: Point) {
    const came_from = buildFrontier(from);
    let current = JSON.stringify(to);
    const path: string[] = [];

    while (current !== JSON.stringify(from)) {
      path.push(current);
      if (!(current = came_from.get(current)!)) {
        break;
      }
    }

    // An undefined current means no possible path was found
    if (current === undefined) {
      return [];
    }

    // path won't include the from Point
    return path.reverse();
  }
}

function parse(input: string) {
  const map = [];
  for (const line of input.split("\n")) {
    map.push(line.trim().split(""));
  }
  return map;
}
