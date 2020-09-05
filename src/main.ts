console.log("OK");

declare class Go {
  importObject: any;
  run(inst: any): Promise<any>;
}

declare function setImage(
  key: string,
  img: HTMLImageElement,
  width: number,
  height: number
): void;
declare function start(): void;
declare function restart(): void;

const loadImage = (src: string) =>
  new Promise((resolve: (img: HTMLImageElement) => void, reject) => {
    const img = new Image();
    img.src = src;
    img.onload = () => {
      resolve(img);
    };
  });

(<any>window).setShareText = (stage: number, score: number) => {
  const button = document.querySelector("#tweet-button") as HTMLButtonElement;
  button.onclick = () => {
    window.location.href = `https://twitter.com/intent/tweet?text=I reached Stage ${stage} and got ${score} Points!&url=https:%2f%2ffkmhrk.github.io%2fgo-wasm-stg%2f`;
  };
};

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
  const [playerImg, heartImg, numberImg, e1Img, e2Img] = await Promise.all([
    loadImage("./player.png"),
    loadImage("./heart.png"),
    loadImage("./number.png"),
    loadImage("./enemy1.png"),
    loadImage("./enemy2.png"),
  ]);

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

    setImage("player", playerImg, 24, 24);
    setImage("heart", heartImg, 18, 18);
    setImage("number", numberImg, 18, 18);
    setImage("enemy1", e1Img, 24, 24);
    setImage("enemy2", e2Img, 24, 24);

    const button = document.querySelector("#start") as HTMLButtonElement;
    button.addEventListener("click", () => {
      button.hidden = true;
      start();
    });
    const restartButton = document.querySelector(
      "#restart"
    ) as HTMLButtonElement;
    restartButton.addEventListener("click", () => {
      const block = document.querySelector("#gameover-block") as HTMLDivElement;
      block.style.display = "none";
      restart();
    });
  } catch (err) {
    console.error(err);
  }
})();
