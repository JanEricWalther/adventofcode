const input = await Deno.readTextFile("input.txt");
let total_score = 0;

const points = Object.freeze({
  X: 0,
  Y: 3,
  Z: 6,
  A: 1,
  B: 2,
  C: 3,
});

for (const line of input.split("\n")) {
  if (line === '')
    continue;
  const opp = line[0];
  const me = line[2];

  switch (opp) {
    case 'A':
      if (me === 'X')
        total_score += points.C;
      else if (me === 'Y')
        total_score += points[opp];
      else 
        total_score += points.B;
      break;
    case 'B':
      if (me === 'X')
        total_score += points.A;
      else if (me === 'Y')
        total_score += points[opp];
      else 
        total_score += points.C;
      break;
    case 'C':
      if (me === 'X')
        total_score += points.B;
      else if (me === 'Y')
        total_score += points[opp];
      else 
        total_score += points.A;
      break;
    default:
      console.error(line);
      throw new Error("undefined input");
  }
  if (me !== 'X' && me !== 'Y' && me !== 'Z')
    throw new Error("undefined input");
  
  total_score += points[me];
}

console.log(`Part 2: ${total_score}`);
