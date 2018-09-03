(function(){
  var player = document.getElementById('player');
  if(!player) {
    return;
  }

  document.getElementById('rate1.0').addEventListener('click', function(){
    console.log('1.0');
    player.defaultPlaybackRate = 1.0;
    player.playbackRate = 1.0;
  });

  document.getElementById('rate1.5').addEventListener('click', function(){
    console.log('x1.5');
    player.defaultPlaybackRate = 1.5;
    player.playbackRate = 1.5;
  });

  document.getElementById('rate1.75').addEventListener('click', function(){
    console.log('x1.75');
    player.defaultPlaybackRate = 1.75;
    player.playbackRate = 1.75;
  });

  document.getElementById('rate2.0').addEventListener('click', function(){
    console.log('x2.0');
    player.defaultPlaybackRate = 2.0;
    player.playbackRate = 2.0;
  });

  document.getElementById('rate10').addEventListener('click', function(){
    console.log('x10');
    player.defaultPlaybackRate = 10;
    player.playbackRate = 10;
  });

})();
