const input = await Deno.readTextFile("input.txt");
let double_assignments = 0;

for (const line of input.trim().split("\n")) {
  const [left, right] = line.split(",");
  const elf1 = left.split("-");
  const elf2 = right.split("-");

  if (contains(elf1, elf2) || contains(elf2, elf1)) {
    double_assignments++;
  }
}

console.log(`Part 1: ${double_assignments}`);

function contains(a : string[], b : string[]): boolean {
  return (Number.parseInt(a[0]) >= Number.parseInt(b[0]) &&
    Number.parseInt(a[1]) <= Number.parseInt(b[1]));
}
