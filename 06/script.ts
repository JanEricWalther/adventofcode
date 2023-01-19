const input = await Deno.readTextFile("input.txt");
const NUMBER_OF_DISTINCT_CHARS = 14;

console.log(solve(input));

function solve(input: string): number {
  let pos = NUMBER_OF_DISTINCT_CHARS;
  const buff: string[] = input.substring(0, NUMBER_OF_DISTINCT_CHARS).split("");

  for (const char of input.substring(NUMBER_OF_DISTINCT_CHARS).split("")) {
    if (!hasDuplicates(buff)) {
      return pos;
    }
    buff.shift();
    buff.push(char);
    pos++;
  }
  return 0;
}

function hasDuplicates(arr: string[]) {
  return new Set(arr).size !== arr.length;
}
