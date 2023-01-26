const input = Deno.readTextFileSync(Deno.args[0]).trim();

console.log("Fewest Steps:", solve(parse(input)));

function parse(input: string) {
  // map[y][x]
  // const map: Point[][] = [];
  const map = [];

  for (const line of input.split("\n")) {
    //   const row: Point[] = [];
    //   for (const char of line.trim().split("")) {
    //     const val = char.charCodeAt(0);
    //     row.push(new Point(val));
    //   }
    map.push(line.trim().split(""));
  }
  return map;
}

function solve(input: string[][]) {
  const map: number[][] = Array(input.length).fill(null).map((_) =>
    Array(input[0].length).fill(false)
  );
  let S;
  let E;
  for (let y = 0; y < input.length; y++) {
    for (let x = 0; x < input[y].length; x++) {
      const cell = input[y][x];
      if (cell === "S") {
        S = { x, y };
        // Your current position (`S`) has elevation `a`
        input[y][x] = "a";
      } else if (cell === "E") {
        E = { x, y };
        // And the location that should get the best signal (`E`) has elevation `z`
        input[y][x] = "z";
      }
      // While we are looping, re-encode chars to ints, to make our elevation comparisons easier

      map[y][x] = input[y][x].charCodeAt(0);
    }
  }
  if (!S || !E) {
    throw new Error("invalid input");
  }

  const path = getShortestPath(S.x, S.y, E.x, E.y);
  return path.length;

  function toId(x: number, y: number) {
    return `${x},${y}`;
  }

  function getNeighbors(x: number, y: number) {
    return [
      { x: x, y: y - 1 },
      { x: x - 1, y: y },
      { x: x + 1, y: y },
      { x: x, y: y + 1 },
    ].filter((coord) => typeof map[coord.y]?.[coord.x] !== "undefined");
  }

  function buildFrontier(from_x: number, from_y: number) {
    const frontier = [];
    frontier.push({ x: from_x, y: from_y });

    const came_from = new Map();
    came_from.set(toId(from_x, from_y), null);
    while (frontier.length > 0) {
      const current = frontier.shift() ?? { x: 0, y: 0 };
      const current_val = map[current.y][current.x];

      const neighbors = getNeighbors(current.x, current.y);
      for (const next of neighbors) {
        const next_cell = map[next.y][next.x];
        const next_id = toId(next.x, next.y);

        if (next_cell - current_val > 1 || came_from.has(next_id)) {
          continue;
        }
        // Coord is walkable
        const current_id = toId(current.x, current.y);
        frontier.push(next);
        came_from.set(next_id, current_id);
      }
    }
    return came_from;
  }

  function getShortestPath(
    from_x: number,
    from_y: number,
    to_x: number,
    to_y: number,
  ) {
    const from_id = toId(from_x, from_y);
    const to_id = toId(to_x, to_y);
    const came_from = buildFrontier(from_x, from_y);
    let current = to_id;

    const path = [];
    while (current !== undefined && current !== from_id) {
      path.push(current);
      current = came_from.get(current);
    }

    // An undefined `current` means it wasn't possible to have a path `from` -> `to`, return an empty path
    if (current === undefined) {
      return [];
    }

    // Finally, put `from` first, and `to` last
    // Note our path won't include the `from` position
    return path.reverse();
  }
}

class Point {
  #height: number;
  possibleNeighbors: Point[];

  constructor(height: number) {
    this.#height = height;
    this.possibleNeighbors = [];
  }

  get height() {
    return this.#height;
  }
}
