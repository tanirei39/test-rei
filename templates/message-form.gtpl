<!DOCTYPE html>
<html>
<head>
    <title>TestRei 投稿画面</title>
</head>
<body>
  <h2>メッセージの送信</h2>
  <form action="message-confirm" method="post">
    <table>
      <tr>
        <td>メッセージ</td>
        <td><input type="text" name="message_text"></td>
      </tr>
    </table>
    <button onclick="sendMessage();">投稿</button>
  </form>
  <button onclick="history.back()">戻る</button>

  <script>
      let ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
      function sendMessage(){
        ws.send(JSON.stringify({
          "event": "event-name",
          "AtrNameInFrontend": "Hello World!",
        }))
      }
    </script>
</body>
</html>