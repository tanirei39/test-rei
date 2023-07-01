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
    <input type="submit" value="投稿">
  </form>
  <button onclick="history.back()">戻る</button>
</body>
</html>