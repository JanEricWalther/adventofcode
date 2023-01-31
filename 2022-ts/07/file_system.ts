class File {
  name: string;
  #size: number;

  constructor(name: string, size: number) {
    this.#size = size;
    this.name = name;
  }
  get size(): number {
    return this.#size;
  }
}

class Dir {
  name: string;
  parentDir: Dir | null;
  contents: Array<Dir | File>;

  constructor(name: string, parentDir: Dir | null) {
    this.name = name;
    this.parentDir = parentDir;
    this.contents = [];
  }

  push(entry: File | Dir) {
    this.contents.push(entry);
  }

  get size(): number {
    return this.contents.map((entry) => entry.size).reduce((a, b) => a + b, 0);
  }
}

class FileSystem {
  root: Dir;

  constructor(instructions: string[]) {
    this.root = new Dir("/", null);
    this.build(instructions);
  }

  get size() {
    return this.root.size;
  }

  build(instructions: string[]) {
    let currentDir: Dir = this.root;

    const cd = (param: string) => {
      switch (param) {
        case "/":
          currentDir = this.root;
          break;
        case "..":
          currentDir = currentDir.parentDir ?? currentDir;
          break;
        default: {
          const newDir = currentDir.contents.find((entry) =>
            entry instanceof Dir && entry.name === param
          );
          currentDir = newDir instanceof Dir ? newDir : currentDir;
        }
      }
    };

    for (const line of instructions) {
      switch (true) {
        case line.startsWith("$ cd "):
          cd(line.split(" ")[2]);
          break;
        case line.startsWith("$ ls"):
          break;
        case line.startsWith("dir"):
          currentDir.push(new Dir(line.split(" ")[1], currentDir));
          break;
        default: {
          // this should be a File
          const [size, name] = line.split(" ");
          currentDir.push(new File(name, +size));
        }
      }
    }
  }

  static *walk(dir: Dir): Generator<Dir | File, void, unknown> {
    for (const entry of dir.contents) {
      yield entry;
      if (entry instanceof Dir) {
        yield* FileSystem.walk(entry);
      }
    }
  }

  *[Symbol.iterator]() {
    yield* FileSystem.walk(this.root);
  }
}

export { Dir, File, FileSystem };
