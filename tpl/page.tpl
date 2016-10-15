<html>

<head>
    <title>{{.Title}}</title>
</head>

<body>
    <h1>{{.Title}}</h1>
    <div>
        <p>{{.Content}}</p>
    </div>
    <div>{{.Date}}</div>
    <h2>Comments</h2>
    <form action="/api/comments" method="POST">
        <p><input hidden="true" type="text" name="page_id" value="{{.ID}}"></p>
        <p><input placeholder="your name" type="text" name="name" value="" required="required"></p>
        <p><input placeholder="your email" type="email" name="email" value="" required="required"></p>
        <p><textarea rows="5" placeholder="input your comments here" name="content" required="required"></textarea></p>
        <p><input type="submit" name="submit"></p>
    </form>


</body>

</html>