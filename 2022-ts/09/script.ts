const input = Deno.readTextFileSync(Deno.args[0]).trim();

console.log(`Part 1: ${visitedPositions(input)}`);

function visitedPositions(instructions: string): number {
  const visited = new Set<string>();
  const head: { x: number; y: number } = { x: 0, y: 0 };
  const tail: { x: number; y: number } = { x: 0, y: 0 };

  visited.add(JSON.stringify({ x: 0, y: 0 }));
  let prevHead = { x: 0, y: 0 };

  for (const instruction of instructions.split("\n")) {
    const [dir, steps] = instruction.trim().split(" ");
    switch (dir) {
      case "L":
        moveLeft(+steps);
        break;
      case "R":
        moveRight(+steps);
        break;
      case "U":
        moveUp(+steps);
        break;
      case "D":
        moveDown(+steps);
        break;
      default:
        throw new Error("undefined Input");
    }
  }
  return visited.size;

  function moveRight(count: number): void {
    while (count > 0) {
      prevHead = { x: head.x, y: head.y };
      head.x++;
      checkMove();
      count--;
    }
  }

  function moveLeft(count: number): void {
    while (count > 0) {
      prevHead = { x: head.x, y: head.y };
      head.x--;
      checkMove();
      count--;
    }
  }

  function moveUp(count: number): void {
    while (count > 0) {
      prevHead = { x: head.x, y: head.y };
      head.y--;
      checkMove();
      count--;
    }
  }

  function moveDown(count: number): void {
    while (count > 0) {
      prevHead = { x: head.x, y: head.y };
      head.y++;
      checkMove();
      count--;
    }
  }

  function checkMove(): void {
    if (Math.abs(head.x - tail.x) < 2 && Math.abs(head.y - tail.y) < 2) {
      return;
    }
    tail.x = prevHead.x;
    tail.y = prevHead.y;
    visited.add(JSON.stringify({ x: tail.x, y: tail.y }));
  }
}
