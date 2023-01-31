const input = await Deno.readTextFile("input.txt");

let greatestCalories = 0;
let calories = 0;

for (const line of input.split("\n")) {
  if (line === "") {
    greatestCalories = calories > greatestCalories ? calories : greatestCalories;
    calories = 0;
    continue;
  }
  calories += Number.parseInt(line);
  
}

console.log(greatestCalories);
