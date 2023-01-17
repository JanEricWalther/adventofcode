const input = await Deno.readTextFile("input.txt");
const inputArr = input.split("\n");
let sum_of_priorities = 0;
let sum_of_badge_priorites = 0;

const map_char_to_priority: { [key: string]: number } = {
  a: 1,
  b: 2,
  c: 3,
  d: 4,
  e: 5,
  f: 6,
  g: 7,
  h: 8,
  i: 9,
  j: 10,
  k: 11,
  l: 12,
  m: 13,
  n: 14,
  o: 15,
  p: 16,
  q: 17,
  r: 18,
  s: 19,
  t: 20,
  u: 21,
  v: 22,
  w: 23,
  x: 24,
  y: 25,
  z: 26,
  A: 27,
  B: 28,
  C: 29,
  D: 30,
  E: 31,
  F: 32,
  G: 33,
  H: 34,
  I: 35,
  J: 36,
  K: 37,
  L: 38,
  M: 39,
  N: 40,
  O: 41,
  P: 42,
  Q: 43,
  R: 44,
  S: 45,
  T: 46,
  U: 47,
  V: 48,
  W: 49,
  X: 50,
  Y: 51,
  Z: 52,
};

one:
for (const line of inputArr) {
  const first_half = line.substring(0, line.length / 2);
  const second_half = line.substring(line.length / 2);

  for (const char of first_half.split("")) {
    if (second_half.includes(char)) {
      sum_of_priorities += map_char_to_priority[char];
      continue one;
    }
  }
}

two:
for (let i = 0; i < inputArr.length; i += 3) {
  for (const char of inputArr[i]) {
    if (inputArr[i + 1].includes(char) && inputArr[i + 2].includes(char)) {
      sum_of_badge_priorites += map_char_to_priority[char];
      continue two;
    }
  }
}

console.log(`Part 1: ${sum_of_priorities}`);
console.log(`Part 2: ${sum_of_badge_priorites}`);
