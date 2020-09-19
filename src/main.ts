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

declare function gtag(...args: any): void;

interface IResultScore {
  score: number;
  date: number;
}

const isAnalyticsEnabled = () => {
  return localStorage.getItem("ga") === "true";
};

const enableAnalytics = () => {
  localStorage.setItem("ga", "true");
};

const startAnalytics = () => {
  gtag("js", new Date());
  gtag("config", "UA-177923604-1");
};

const loadImage = (src: string) =>
  new Promise((resolve: (img: HTMLImageElement) => void, reject) => {
    const img = new Image();
    img.src = src;
    img.onload = () => {
      resolve(img);
    };
  });

const toDateStr = (timestamp: number) => {
  const date = new Date(timestamp * 1000);
  const addZero = (v: number) => {
    if (v < 10) {
      return `0${v}`;
    } else {
      return `${v}`;
    }
  };
  return `${date.getFullYear()}/${addZero(date.getMonth() + 1)}/${addZero(
    date.getDate()
  )}`;
};
const showTopScores = (scores: IResultScore[]) => {
  const tbody = document.querySelector("#score_list")!;
  tbody.innerHTML = "";
  for (let i = 0; i < 5 && i < scores.length; i++) {
    const s = scores[i];
    const tr = document.createElement("tr");
    const rankTh = document.createElement("th");
    rankTh.innerText = `${i + 1}`;
    tr.appendChild(rankTh);

    const scoreTd = document.createElement("td");
    scoreTd.innerText = `${s.score}`;
    tr.appendChild(scoreTd);

    const dateTd = document.createElement("td");
    dateTd.innerText = toDateStr(s.date);
    tr.appendChild(dateTd);

    tbody.appendChild(tr);
  }
};

(<any>window).event = (action: string) => {
  gtag("event", action);
};

(<any>window).showResult = (
  stage: number,
  score: number,
  resultStr: string
) => {
  const result = JSON.parse(resultStr);
  showTopScores(result.scores);
  const button = document.querySelector("#tweet-button") as HTMLButtonElement;
  button.onclick = () => {
    window.location.href = `https://twitter.com/intent/tweet?text=I reached Stage ${stage} and got ${score} Points!&url=https:%2f%2ffkmhrk.github.io%2fgo-wasm-stg%2f&hashtags=piyo2020`;
  };
};

(<any>window).updateResult = (resultStr: string) => {
  const result = JSON.parse(resultStr);
  document.querySelector("#playCount")!!.innerHTML = `${result.start_count}`;
  document.querySelector("#deathCount")!!.innerHTML = `${result.death_count}`;
  document.querySelector("#days")!!.innerHTML = `${result.days}`;
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

const initEventListeners = () => {
  const button = document.querySelector("#start") as HTMLButtonElement;
  button.addEventListener("click", () => {
    button.hidden = true;
    if (isAnalyticsEnabled()) {
      start();
    } else {
      // show analytics
      const block = document.querySelector(
        "#analytics_block"
      ) as HTMLDivElement;
      block.style.display = "flex";
    }
  });

  const restartButton = document.querySelector("#restart") as HTMLButtonElement;
  restartButton.addEventListener("click", () => {
    const block = document.querySelector("#gameover-block") as HTMLDivElement;
    block.style.display = "none";
    restart();
  });

  const acceptButton = document.querySelector("#accept") as HTMLButtonElement;
  acceptButton.addEventListener("click", () => {
    const block = document.querySelector("#analytics_block") as HTMLDivElement;
    block.style.display = "none";
    enableAnalytics();
    startAnalytics();
    start();
  });
};

// boot
(async () => {
  if (isAnalyticsEnabled()) {
    startAnalytics();
  }

  if ("serviceWorker" in navigator) {
    window.addEventListener("load", () => {
      navigator.serviceWorker.register("./sw.js");
    });
  }

  const titleImg = await loadImage("./title.png");
  const canvas = document.querySelector("#canvas") as HTMLCanvasElement;
  const context = canvas.getContext("2d")!;
  context.drawImage(titleImg, 0, 0, 640, 960, 0, 0, 320, 480);

  const [
    playerImg,
    heartImg,
    numberImg,
    e1Img,
    e2Img,
    e11Img,
    e12Img,
    e13Img,
    e14Img,
    item1Img,
  ] = await Promise.all([
    loadImage("./player.png"),
    loadImage("./heart.png"),
    loadImage("./number.png"),
    loadImage("./enemy1.png"),
    loadImage("./enemy2.png"),
    loadImage("./enemy11.png"),
    loadImage("./enemy12.png"),
    loadImage("./enemy13.png"),
    loadImage("./enemy14.png"),
    loadImage("./item1.png"),
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
    setImage("enemy11", e11Img, 40, 40);
    setImage("enemy12", e12Img, 40, 40);
    setImage("enemy13", e13Img, 40, 40);
    setImage("enemy14", e14Img, 40, 40);
    setImage("item1", item1Img, 12, 12);

    initEventListeners();
  } catch (err) {
    console.error(err);
  }
})();
