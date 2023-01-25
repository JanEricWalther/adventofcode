export type WorryFunction = (oldVal: number) => number;

class Item {
  worry: number;

  constructor(worry: number) {
    this.worry = worry;
  }
}

class Monkey {
  id: number;
  items: Item[];
  divisibleBy: number;
  worryFn: WorryFunction;
  if_false_throw_to: number;
  if_true_throw_to: number;
  inspectionCount: number;

  constructor(
    config: {
      id: number;
      items: number[];
      worryFn: WorryFunction;
      divisibleBy: number;
      ifFalse: number;
      ifTrue: number;
    },
  ) {
    this.id = config.id;
    this.items = config.items.map((i) => new Item(i));
    this.divisibleBy = config.divisibleBy;
    this.if_false_throw_to = config.ifFalse;
    this.if_true_throw_to = config.ifTrue;
    this.worryFn = config.worryFn;
    this.inspectionCount = 0;
  }
}

class Game {
  monkeys: Array<Monkey>;
  isPartOne: boolean;
  mod: number;

  constructor(monkeys: Array<Monkey>, isPartOne: boolean) {
    this.monkeys = monkeys;
    this.isPartOne = isPartOne;

    // divisibleBy are always Primes
    // so we can mod the item by the Product of all divisibleBys 
    // to keep the number at a managable Size
    this.mod = monkeys.map((monkey) => monkey.divisibleBy).reduce(
      (a, b) => a * b,
      1,
    );
  }

  playRound = () => {
    for (const monkey of this.monkeys) {
      for (const item of monkey.items) {
        item.worry = monkey.worryFn(item.worry);
        if (this.isPartOne) {
          item.worry = Math.trunc(item.worry / 3);
        } else {
          item.worry %= this.mod; 
        }
        if (item.worry % monkey.divisibleBy === 0) {
          this.monkeys[monkey.if_true_throw_to].items.push(item);
        } else {
          this.monkeys[monkey.if_false_throw_to].items.push(item);
        }
      }
      monkey.inspectionCount += monkey.items.length;
      monkey.items = [];
    }
  };

  get inspectionCounts() {
    return this.monkeys.map((monkey) => monkey.inspectionCount).sort((a, b) =>
      b - a
    );
  }
}

export { Game, Monkey };
