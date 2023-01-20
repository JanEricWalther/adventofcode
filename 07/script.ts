import { Dir, File, FileSystem } from "./file_system.ts";
const DISK_SIZE = 70_000_000;
const UPDATE_SIZE = 30_000_000;

const input = Deno.readTextFileSync(Deno.args[0]);
const drive = new FileSystem(input.split("\n"));
const dirs = [...drive].filter((entry) => entry instanceof Dir);

console.info(`Total Size: ${drive.size}`);
console.log(`Part 1: ${sumOfSmallDirs(dirs)}`);
console.log(`Part 2: ${smallestDirForUpdate(dirs)}`);

function sumOfSmallDirs(dirs: (Dir | File)[]) {
  let sum = 0;
  for (const item of dirs) {
    const size = item.size;
    if (size <= 100_000) {
      sum += size;
    }
  }
  return sum;
}

function smallestDirForUpdate(dirs: (Dir | File)[]): number {
  const spaceNeededForUpdate = -(DISK_SIZE - UPDATE_SIZE - drive.size);
  const possibleDirs: number[] = [];
  for (const dir of dirs) {
    if (dir.size >= spaceNeededForUpdate)
      possibleDirs.push(dir.size);
  }
  return Math.min(...possibleDirs);
}
