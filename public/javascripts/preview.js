/** @type {HTMLVideoElement} */
const videoElement = document.getElementById('video');

/** @type {HTMLInputElement} */
const videoUpload = document.getElementById('video-upload');

videoUpload.addEventListener('change', function (event) {
  if (event.target.files && event.target.files[0]) {
    var reader = new FileReader();

    reader.onload = function (e) {
      console.log('loaded');
      videoElement.src = e.target.result;
      videoElement.load();
    }.bind(this);

    reader.readAsDataURL(event.target.files[0]);
  }
});
