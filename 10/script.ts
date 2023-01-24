const input = Deno.readTextFileSync(Deno.args[0]).trim();

console.log(`Part 1: ${signalStrength(input.split("\n"))}`);

function signalStrength(instructions: string[]): number {
  let cycle = 1;
  let signalStrength = 0;
  let x_reg = 1;

  for (const instruction of instructions) {
    const [opCode, arg] = instruction.trim().split(" ");

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
    checkCycle();
  }

  return signalStrength;

  function checkCycle() {
    if ((cycle - 20) % 40 === 0) {
      signalStrength += cycle * x_reg;
      console.log(`Cycle: ${cycle}\nX: ${x_reg}\nSignal: ${signalStrength}`);
    }
  }
}
