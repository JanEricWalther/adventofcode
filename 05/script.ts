[
  ["D", "M", "S", "Z", "R", "F", "W", "N"],
  ["W", "P", "Q", "G", "S"],
  ["W", "R", "V", "Q", "F", "N", "J", "C"],
  ["F", "Z", "P", "C", "G", "D", "L"],
  ["T", "P", "S"],
  ["H", "D", "F", "W", "R", "L"],
  ["Z", "N", "D", "C"],
  ["W", "N", "R", "F", "V", "S", "J", "Q"],
  ["R", "M", "S", "G", "Z", "W", "V"],
];

const input = await Deno.readTextFile("input.txt");
const [raw_stacks, raw_instructions] = input.split("\n\n");
const stacks = parseStacks(raw_stacks);

for (const line of raw_instructions.trim().split("\n")) {
  // parseMove9000(line);
  parseMove9001(line);
}

let out = "";
for (const stack of stacks) {
  out += stack.pop();
}
console.log(out);

// Part 1
function parseMove9000(input: string) {
  const move = input.split(" ");
  const [, count, , from, , to] = move;

  for (let i = +count; i > 0; i--) {
    const item = stacks[+from - 1].pop() || "";
    stacks[+to - 1].push(item);
  }
}

// Part 2
function parseMove9001(input: string) {
  const move = input.split(" ");
  const [, countS, , fromS, , toS] = move;
  const count = Number.parseInt(countS);
  const from = Number.parseInt(fromS) - 1;
  const to = Number.parseInt(toS) - 1;
  
  const items = stacks[from].splice(-count, count) || "";
  stacks[to].push(...items);
}

function parseStacks(input: string): string[][] {
  const CRATE_SIZE = "[A]".length;
  const stacks: string[][] = [];

  for (const line of input.split("\n")) {
    if (line.includes("1")) {
      continue;
    }

    for (let i = 0; i < line.length; i += CRATE_SIZE + 1) {
      const crate = line.substring(i, i + CRATE_SIZE);
      const idx = i / (CRATE_SIZE + 1);

      if (!stacks[idx]) {
        stacks[idx] = [];
      }
      if (crate.trim()) {
        const [, letter] = crate.split("");
        stacks[idx].push(letter);
      }
    }
  }

  for (const stack of stacks) {
    stack.reverse();
  }
  return stacks;
}
