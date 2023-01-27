const input = Deno.readTextFileSync(Deno.args[0]).trim();

console.log(sumOfPackets(parse(input)));
console.log(dividerPacketPositions(parse(input)));

function sumOfPackets(packets: number[][][]) {
  const rightOrder: number[] = [];
  let pairIndex = 1;

  for (const pair of packets) {
    const [leftPack, rightPack] = pair;
    const cmp = compare(leftPack, rightPack);
    if (cmp) {
      rightOrder.push(pairIndex);
    }
    pairIndex++;
  }

  return rightOrder.reduce((a, b) => a + b, 0);
}

function dividerPacketPositions(packets: any[]) {
  const divider2 = [[2]];
  const divider6 = [[6]];

  packets = [...packets.flat(), divider2, divider6];
  packets.sort((a, b) => {
    const cmp = compare(a, b);
    if (cmp) {
      return -1;
    } else if (cmp !== null) {
      // false && not null
      return 1;
    } else {
      return 0;
    }
  });

  const pos2 = packets.indexOf(divider2)+1;
  const pos6 = packets.indexOf(divider6)+1;
  return pos2*pos6;
}

function compare(left: any[], right: any[]): boolean | null {
  // Loop the upper bounds of left and right
  for (let i = 0; i < Math.max(left.length, right.length); i++) {
    const leftItem = left[i];
    const rightItem = right[i];

    // Out of bounds checks
    if (leftItem === undefined && rightItem !== undefined) {
      return true;
    } else if (rightItem === undefined) {
      return false;
    }

    if (typeof leftItem === "number" && typeof rightItem === "number") {
      if (leftItem < rightItem) {
        return true;
      } else if (leftItem > rightItem) {
        return false;
      } else {
        // Same value, continue checking next
        continue;
      }
    }

    // Recursion!
    let substep = null;

    if (Array.isArray(leftItem) && Array.isArray(rightItem)) {
      // [] - []
      substep = compare(leftItem, rightItem);
    } else if (Array.isArray(leftItem) && !Array.isArray(rightItem)) {
      // [] - number
      substep = compare(leftItem, [rightItem]);
    } else if (!Array.isArray(leftItem) && Array.isArray(rightItem)) {
      // number - []
      substep = compare([leftItem], rightItem);
    }

    // A null result means to pop out of our recursive step onto the next item
    if (substep !== null) {
      return substep;
    }
  }

  // This (sub)list was neither right nor wrong
  return null;
}

function parse(input: string) {
  // deno-lint-ignore no-explicit-any
  const packets: any[] = [];
  for (const pair of input.split("\n\n")) {
    packets.push(pair.split("\n").map((line) => JSON.parse(line)));
  }
  return packets;
}
