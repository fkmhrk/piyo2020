importScripts(
  "https://storage.googleapis.com/workbox-cdn/releases/5.1.2/workbox-sw.js"
);

workbox.setConfig({
  debug: false,
});
const { registerRoute } = workbox.routing;
const { CacheFirst, NetworkFirst } = workbox.strategies;
const { CacheableResponse } = workbox.cacheableResponse;

registerRoute(new RegExp(".+\\.wasm$"), new NetworkFirst());
registerRoute(new RegExp(".+\\.js$"), new NetworkFirst());
registerRoute(new RegExp(".+\\.html$"), new NetworkFirst());
registerRoute(new RegExp(".+\\.png$"), new CacheFirst());
