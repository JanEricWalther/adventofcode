const input = await Deno.readTextFile("input.txt");

const elves: number[] = [];
let calories = 0;

for (const line of input.split("\n")) {
  if (line === "") {
    elves.push(calories);
    calories = 0;
    continue;
  }
  calories += Number.parseInt(line);
}

elves.sort((a, b) => b - a);

console.log(`Part 2: ${
  elves.slice(0, 3)
    .reduce((a, b) => a + b, 0)
}`);
