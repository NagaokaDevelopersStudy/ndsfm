(function(){
  var feedUrl = "https://pcr.apple.com/id1415356362"
  var itunesUrl = "https://itunes.apple.com/jp/podcast/ndsfm/id1415356362";

  var link = document.getElementById("subscribe_link");
  if (window.navigator.platform.match(/(MacIntel|iPhone|iPad|iPod)/)) {
    link.href = itunesUrl;
  } else {
    link.href = feedUrl;
  }
})();
