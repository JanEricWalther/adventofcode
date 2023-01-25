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

  constructor(monkeys: Array<Monkey>) {
    this.monkeys = monkeys;    
  }

  playRound = () => {
    for (const monkey of this.monkeys) {
      for (const item of monkey.items) {
        item.worry = Math.floor(monkey.worryFn(item.worry) / 3);
        if (item.worry % monkey.divisibleBy === 0) {
          this.monkeys[monkey.if_true_throw_to].items.push(item);
        } else {
          this.monkeys[monkey.if_false_throw_to].items.push(item);
        }
        monkey.inspectionCount++;
      }
      monkey.items = [];
    }
  }

  get inspectionCounts() {
    return this.monkeys.map(monkey => monkey.inspectionCount);
  }
}

export { Game,  Monkey };
