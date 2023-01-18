const input = await Deno.readTextFile("input.txt");
let double_assignments = 0;
let overlaping_assignments = 0;

for (const line of input.trim().split("\n")) {
  const [left, right] = line.split(",");
  const elf1 = left.split("-");
  const elf2 = right.split("-");

  if (contains(elf1, elf2) || contains(elf2, elf1)) {
    double_assignments++;
  }
  if (overlaps(elf1, elf2)) {
    overlaping_assignments++;
  }
}

console.log(`Part 1: ${double_assignments}`);
console.log(`Part 2: ${overlaping_assignments}`);

function contains(a: string[], b: string[]): boolean {
  return (Number.parseInt(a[0]) >= Number.parseInt(b[0]) &&
    Number.parseInt(a[1]) <= Number.parseInt(b[1]));
}

function overlaps(a: string[], b: string[]): boolean {
  const a0 = Number.parseInt(a[0]);
  const a1 = Number.parseInt(a[1]);
  const b0 = Number.parseInt(b[0]);
  const b1 = Number.parseInt(b[1]);

  if (a0 === b0 || a0 === b1 || a1 === b0 || a1 === b1) {
    return true;
  } else if (a1 >= b0 && a0 <= b0) {
    return true;
  } else if (b1 >= a0 && b0 <= a0) {
    return true;
  }
  return false;
}
