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
      <th>削除</th>
      <th>編集</th>
    </tr>
    {{ range . }}
    <tr>
      <td>{{.Id}}</td>
      <td>{{.Text}}</td>
      <td> <button onclick="location.href='./message-delete/{{.Id}}'">削除</button></td>
      <td> <button onclick="location.href='./message-edit/{{.Id}}'">編集</button></td>
    </tr>
    {{ end }}
  </table>
  <button onclick="location.href='./message-form'">新規投稿</button>
</body>
</html>

<script>
  function OneChanged(){
  }
 </script>