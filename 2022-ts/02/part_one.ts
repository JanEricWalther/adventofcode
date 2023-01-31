const input = await Deno.readTextFile("input.txt");
let total_score = 0;

const points = Object.freeze({
  DRAW: 3,
  WIN: 6,
  X: 1,
  Y: 2,
  Z: 3,
});

for (const line of input.split("\n")) {
  if (line === '')
    continue;
  const opp = line[0];
  const me = line[2];

  switch (me) {
    case "X":
      switch (opp) {
        case "A":
          total_score += points.DRAW;
          break;
        case "B":
          // LOSS
          break;
        case "C":
          total_score += points.WIN;
          break;
      }
      break;
    case "Y":
      switch (opp) {
        case "A":
          total_score += points.WIN;
          break;
        case "B":
          total_score += points.DRAW;
          break;
        case "C":
          // LOSS
          break;
      }
      break;
    case "Z":
      switch (opp) {
        case "A":
          // LOSS
          break;
        case "B":
          total_score += points.WIN;
          break;
        case "C":
          total_score += points.DRAW;
          break;
      }
      break;
    default:
      console.error(line);
      throw new DOMException("undefined input");
  }
  total_score += points[me];
}

console.log(`Part 1: ${total_score}`);
