/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/main.ts");
/******/ })
/************************************************************************/
/******/ ({

/***/ "./src/main.ts":
/*!*********************!*\
  !*** ./src/main.ts ***!
  \*********************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
eval("\nvar __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {\n    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\nconsole.log(\"OK\");\nconst loadImage = (src) => new Promise((resolve, reject) => {\n    const img = new Image();\n    img.src = src;\n    img.onload = () => {\n        resolve(img);\n    };\n});\nconst toDateStr = (timestamp) => {\n    const date = new Date(timestamp * 1000);\n    const addZero = (v) => {\n        if (v < 10) {\n            return `0${v}`;\n        }\n        else {\n            return `${v}`;\n        }\n    };\n    return `${date.getFullYear()}/${addZero(date.getMonth() + 1)}/${addZero(date.getDate())}`;\n};\nconst showTopScores = (scores) => {\n    const tbody = document.querySelector(\"#score_list\");\n    tbody.innerHTML = \"\";\n    for (let i = 0; i < 5 && i < scores.length; i++) {\n        const s = scores[i];\n        const tr = document.createElement(\"tr\");\n        const rankTh = document.createElement(\"th\");\n        rankTh.innerText = `${i + 1}`;\n        tr.appendChild(rankTh);\n        const scoreTd = document.createElement(\"td\");\n        scoreTd.innerText = `${s.score}`;\n        tr.appendChild(scoreTd);\n        const dateTd = document.createElement(\"td\");\n        dateTd.innerText = toDateStr(s.date);\n        tr.appendChild(dateTd);\n        tbody.appendChild(tr);\n    }\n};\nwindow.showResult = (stage, score, resultStr) => {\n    const result = JSON.parse(resultStr);\n    showTopScores(result.scores);\n    const button = document.querySelector(\"#tweet-button\");\n    button.onclick = () => {\n        window.location.href = `https://twitter.com/intent/tweet?text=I reached Stage ${stage} and got ${score} Points!&url=https:%2f%2ffkmhrk.github.io%2fgo-wasm-stg%2f`;\n    };\n};\nwindow.updateResult = (resultStr) => {\n    const result = JSON.parse(resultStr);\n    document.querySelector(\"#playCount\").innerHTML = `${result.start_count}`;\n    document.querySelector(\"#deathCount\").innerHTML = `${result.death_count}`;\n    document.querySelector(\"#days\").innerHTML = `${result.days}`;\n};\nif (!WebAssembly.instantiateStreaming) {\n    // polyfill\n    WebAssembly.instantiateStreaming = (resp, importObject) => __awaiter(void 0, void 0, void 0, function* () {\n        const source = yield (yield resp).arrayBuffer();\n        return yield WebAssembly.instantiate(source, importObject);\n    });\n}\n// boot\n(() => __awaiter(void 0, void 0, void 0, function* () {\n    const [playerImg, heartImg, numberImg, e1Img, e2Img, e11Img, e12Img, item1Img,] = yield Promise.all([\n        loadImage(\"./player.png\"),\n        loadImage(\"./heart.png\"),\n        loadImage(\"./number.png\"),\n        loadImage(\"./enemy1.png\"),\n        loadImage(\"./enemy2.png\"),\n        loadImage(\"./enemy11.png\"),\n        loadImage(\"./enemy12.png\"),\n        loadImage(\"./item1.png\"),\n    ]);\n    const go = new Go();\n    let mod;\n    let inst;\n    try {\n        const result = yield WebAssembly.instantiateStreaming(fetch(\"./main.wasm\"), go.importObject);\n        mod = result.module;\n        inst = result.instance;\n        go.run(inst);\n        setImage(\"player\", playerImg, 24, 24);\n        setImage(\"heart\", heartImg, 18, 18);\n        setImage(\"number\", numberImg, 18, 18);\n        setImage(\"enemy1\", e1Img, 24, 24);\n        setImage(\"enemy2\", e2Img, 24, 24);\n        setImage(\"enemy11\", e11Img, 40, 40);\n        setImage(\"enemy12\", e12Img, 40, 40);\n        setImage(\"item1\", item1Img, 12, 12);\n        const button = document.querySelector(\"#start\");\n        button.addEventListener(\"click\", () => {\n            button.hidden = true;\n            start();\n        });\n        const restartButton = document.querySelector(\"#restart\");\n        restartButton.addEventListener(\"click\", () => {\n            const block = document.querySelector(\"#gameover-block\");\n            block.style.display = \"none\";\n            restart();\n        });\n    }\n    catch (err) {\n        console.error(err);\n    }\n}))();\n\n\n//# sourceURL=webpack:///./src/main.ts?");

/***/ })

/******/ });