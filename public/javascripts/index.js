/**
 *  @type {HTMLVideoElement}
 */
const videoElement = document.getElementById('video');
let increaseRateInterval;

videoElement.onkeydown = function (ev) {
  switch (ev.code) {
    case 'KeyF':
      return videoElement.requestFullscreen();
  }
};

videoElement.onloadeddata = function (ev) {
  console.log('onloadeddata', ev);
};

videoElement.onpause = function (ev) {
  clearInterval(increaseRateInterval);
};

videoElement.onprogress = function (ev) {
  console.log(videoElement);
  console.log(ev);
};

async function setVideoSource(src) {
  const hadPause = videoElement.paused;
  const hadTime = videoElement.currentTime;
  const poster = videoElement.poster;
  videoElement.poster = "";
  videoElement.src = src;
  if (hadPause) {
    return;
  }
  videoElement.currentTime = hadTime;
  await videoElement.play();
  videoElement.poster = poster;
  videoElement.DOCUMENT_FRAGMENT_NODE
}
