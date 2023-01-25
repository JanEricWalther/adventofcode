import { Game, Monkey, WorryFunction } from "./monkey.ts";

const input = Deno.readTextFileSync(Deno.args[0]).trim();

console.log(monkeyBusiness(input));

function monkeyBusiness(input: string): number {
  const monkeys: Monkey[] = [];
  for (const config of parseInput(input)) {
    monkeys.push(new Monkey(config));
  }

  const game = new Game(monkeys);
  for (const i of Array(20).keys()) {
    game.playRound();
  }
  const counts = game.inspectionCounts.sort((a, b) => b - a);

  return counts[0] * counts[1];
}

function parseInput(input: string) {
  const store = [];
  for (const block of input.split("\n\n")) {
    /*
    Monkey 0:
      Starting items: 79, 98
      Operation: new = old * 19
      Test: divisible by 23
        If true: throw to monkey 2
        If false: throw to monkey 3
    */
    const lines = block.trim().split("\n");
    const monkey = lines.reduce<
      {
        id: number;
        items: number[];
        worryFn: WorryFunction;
        divisibleBy: number;
        ifTrue: number;
        ifFalse: number;
      }
    >((acc, line, i) => {
      line = line.trim();

      switch (i) {
        case 0: {
          const id = +line.match(/Monkey (\d+):/)![1];
          acc.id = id;
          break;
        }
        case 1: {
          const items = line.split(":")[1].split(", ").map((i) => {
            return +i.trim();
          });
          acc.items = items;
          break;
        }
        case 2: {
          const operation = line.split(":")[1].trim();
          // new = old * 19
          const [, , left, op, right] = operation.split(" ");
          const worryFn = (oldVal: number) => {
            const leftVal = left === "old" ? oldVal : +left;
            const rightVal = right === "old" ? oldVal : +right;
            // op is only + or *
            return op === "+" ? leftVal + rightVal : leftVal * rightVal;
          };
          acc.worryFn = worryFn;
          break;
        }
        case 3: {
          const divisibleBy = line.split(" ").at(-1)!;
          acc.divisibleBy = +divisibleBy;
          break;
        }
        case 4: {
          const ifTrue = line.split(" ").at(-1)!;
          acc.ifTrue = +ifTrue;
          break;
        }
        case 5: {
          const ifFalse = line.split(" ").at(-1)!;
          acc.ifFalse = +ifFalse;
          break;
        }
        default:
          throw new Error("Invalid input.");
      }
      return acc;
    }, {
      id: -1,
      items: [],
      worryFn: () => 0,
      divisibleBy: 0,
      ifTrue: -1,
      ifFalse: -1,
    });
    store.push(monkey);
  }
  return store;
}
