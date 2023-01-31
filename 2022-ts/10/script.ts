const input = Deno.readTextFileSync(Deno.args[0]).trim();

console.log(`Signal Strength: ${solve(input.split("\n"))}`);

function solve(instructions: string[]): number {
  let cycle = 1;
  let signalStrength = 0;
  let x_reg = 1;
  console.log('');
  
  for (const instruction of instructions) {
    const [opCode, arg] = instruction.trim().split(" ");
    checkCycle();
    switch (opCode) {
      case "addx":
        cycle++;
        checkCycle();
        x_reg += +arg;
        cycle++;
        break;
      case "noop":
        cycle++;
        break;
      default:
        throw new Error("Invalid input.");
    }
  }
  console.log('');
  return signalStrength;

  function checkCycle() {
    if ((cycle - 20) % 40 === 0) {
      signalStrength += cycle * x_reg;
    }
    draw();
  }

  function draw() {
    const encoder = new TextEncoder();
    const pos = cycle > 40 ? (cycle % 40) - 1 : cycle - 1;
    const pixel = Math.abs(x_reg - pos) < 2 ? "#" : ".";

    Deno.stdout.writeSync(encoder.encode(pixel));
    if (cycle % 40 === 0) {
      Deno.stdout.writeSync(encoder.encode("\n"));
      
    }
  }
}
