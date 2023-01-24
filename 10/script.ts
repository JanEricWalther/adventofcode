const input = Deno.readTextFileSync(Deno.args[0]).trim();

console.log(`Part 1: ${signalStrength(input.split("\n"))}`);

function signalStrength(instructions: string[]): number {
  let cycle = 1;
  let signalStrength = 0;
  let x_reg = 1;

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

  return signalStrength;

  function checkCycle() {
    if ((cycle - 20) % 40 === 0) {
      signalStrength += cycle * x_reg;
    }
    draw();
  }

  function draw() {
    const encoder = new TextEncoder();
    let pos = cycle - 1;
    if (cycle > 40) {
      pos = (cycle % 40) - 1;
    }

    let pixel = ".";
    if (Math.abs(x_reg - pos) < 2) {
      pixel = "#";
    }

    Deno.stdout.writeSync(encoder.encode(pixel));
    if (cycle % 40 === 0) {
      Deno.stdout.writeSync(encoder.encode("\n"));
    }
  }
}
