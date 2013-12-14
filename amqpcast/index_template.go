package amqpcast

const indexTemplate = `
<!DOCTYPE html>
<html>

<head>
<title>amqpcast</title>
<meta charset="utf-8">
<script src="http://code.jquery.com/jquery-2.0.3.min.js"></script>
<style>
#messages {
  background-color: black;
  color: #3f0;
  padding : 1em;
}

.time {
  color: #ccc;
  margin-right: 1em;
}
</style>
</head>

<body>

<h1>amqpcast</h1>

<pre id="messages"></pre>

<script>
$(function() {
    var ws = new WebSocket("ws://" + window.location.host + "/ws")

    ws.onopen = function() {
        console.log("websocket open");
    };

    ws.onmessage = function(e) {
        var dateSpan = $('<span/>').addClass('time').append(
            (new Date()).toISOString()),
            messageSpan = $('<span/>').addClass('message').append(e.data);
        $('#messages').prepend(
            $('<br/>')).prepend(
            messageSpan).prepend(
            dateSpan);
    };

    ws.onclose = function(e) {
        console.log("closed");
        console.log(e);
    };

    ws.onerror = function(e) {
        console.log("error");
        console.log(e);
    }
});
</script>

</body>
</html>
`
