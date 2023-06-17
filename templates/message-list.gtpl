<!DOCTYPE html>
<html>
<head>
    <title>TestRei 投稿一覧</title>
</head>
<body>
  <h2>投稿一覧</h2>
  <table class="list">
    <tr>
      <th>ID</th>
      <th>メッセージ</th>
    </tr>
    {{ range . }}
    <tr>
      <td>{{.Id}}</td>
      <td>{{.Text}}</td>
    </tr>
    {{ end }}
  </table>
  <button onclick="location.href='./message-form'">新規投稿</button>
</body>
</html>