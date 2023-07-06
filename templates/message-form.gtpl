<!DOCTYPE html>
<html>
<head>
    <title>TestRei {{.Title1}}</title>
</head>
<body>
  <h2>{{.Title2}}</h2>
  <form action="{{.Action}}" method="post">
    <table>
      <tr>
        <td>メッセージ</td>
        <td><input type="text" name="message_text" value="{{.Message}}"></td>
      </tr>
    </table>
    <input type="submit" value={{.DoneText}}>
  </form>
  <button onclick="history.back()">戻る</button>
</body>
</html>