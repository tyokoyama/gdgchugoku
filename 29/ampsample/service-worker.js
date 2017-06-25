var dataCacheName = 'weatherData-v1';
var cacheName = 'weatherPWA-step-5-2';
var filesToCache = [
  '/',
  '/app/pwacodelab/index.html',
  '/app/pwacodelab/scripts/app.js',
  '/app/pwacodelab/styles/inline.css',
  '/app/pwacodelab/images/clear.png',
  '/app/pwacodelab/images/cloudy-scattered-showers.png',
  '/app/pwacodelab/images/cloudy.png',
  '/app/pwacodelab/images/fog.png',
  '/app/pwacodelab/images/ic\_add\_white\_24px.svg',
  '/app/pwacodelab/images/ic\_refresh\_white\_24px.svg',
  '/app/pwacodelab/images/partly-cloudy.png',
  '/app/pwacodelab/images/rain.png',
  '/app/pwacodelab/images/scattered-showers.png',
  '/app/pwacodelab/images/sleet.png',
  '/app/pwacodelab/images/snow.png',
  '/app/pwacodelab/images/thunderstorm.png',
  '/app/pwacodelab/images/wind.png',
  '/app/pwacodelab/images/icons/icon-144x144.png'
];

self.addEventListener('install', function(e) {
  console.log('[ServiceWorker] Install');
  e.waitUntil(
    caches.open(cacheName).then(function(cache) {
      console.log('[ServiceWorker] Caching app shell');
      return cache.addAll(filesToCache);
    })
  );
});

self.addEventListener('activate', function(e) {
  console.log('[ServiceWorker] Activate');
  e.waitUntil(
    caches.keys().then(function(keyList) {
      return Promise.all(keyList.map(function(key) {
        console.log('[ServiceWorker] Removing old cache', key);
        if (key !== cacheName) {
          return caches.delete(key);
        }
      }));
    })
  );
  return self.clients.claim();
});

self.addEventListener('fetch', function(e) {
  console.log('[ServiceWorker] Fetch', e.request.url);
  var dataUrl = 'https://query.yahooapis.com/v1/public/yql';
  if (e.request.url.indexOf(dataUrl) === 0) {
    e.respondWith(
    fetch(e.request)
        .then(function(response) {
        return caches.open(dataCacheName).then(function(cache) {
            cache.put(e.request.url, response.clone());
            console.log('[ServiceWorker] Fetched&Cached Data');
            return response;
            });
        })
    );
  } else {
    e.respondWith(
      caches.match(e.request).then(function(response) {
        return response || fetch(e.request);
    })
  );
}});