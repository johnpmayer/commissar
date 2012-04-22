
// chat

define(
  ["jquery"],
  function(){
    
    function initChat(path, logElem, inputElem) {
      
      function message(msg) {
        logElem.append("<p>" + msg + "</p>");
      }
      
      var url = "ws://" + window.location.host + "/" + path;
      
      var socket = new WebSocket(url);
            
      socket.onmessage = function(msg) {
        message("server: " + msg.data);
      };
            
      function send() {
        var text = inputElem.val();
        if(text.length > 0) {
          socket.send(text);
          inputElem.val("");
        }
      }
      
      inputElem.keypress(function(event) {
        if (event.keyCode == '13') {  
          send();
        }
      });
      
    }
    
    return {
      initChat:initChat
    
    }
    
  }
);

