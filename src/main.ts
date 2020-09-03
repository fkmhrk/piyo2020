console.log("OK");

declare class Go {
  importObject: any;
  run(inst: any): Promise<any>;
}

declare function setImage(key: string, img: HTMLImageElement): void;
declare function start(): void;

const loadImage = (src: string) =>
  new Promise((resolve: (img: HTMLImageElement) => void, reject) => {
    const img = new Image();
    img.src = src;
    img.onload = () => {
      resolve(img);
    };
  });

if (!WebAssembly.instantiateStreaming) {
  // polyfill
  (<any>WebAssembly.instantiateStreaming) = async (
    resp: any,
    importObject: any
  ) => {
    const source = await (await resp).arrayBuffer();
    return await WebAssembly.instantiate(source, importObject);
  };
}

// boot
(async () => {
  const playerImg = await loadImage("./player.png");

  const go = new Go();
  let mod: any;
  let inst: any;

  try {
    const result = await WebAssembly.instantiateStreaming(
      fetch("./main.wasm"),
      go.importObject
    );
    mod = result.module;
    inst = result.instance;
    go.run(inst);

    setImage("player", playerImg);

    const button = document.querySelector("#start") as HTMLButtonElement;
    button.addEventListener("click", () => {
      button.hidden = true;
      start();
    });
  } catch (err) {
    console.error(err);
  }
})();
