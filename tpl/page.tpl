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
    <h2>Comment List</h2>
    <div id="comment-list">
        {{range .Comments}}
        <div id="comment-node comment-node-shown">
            <p>Commented by {{.Name}} ({{.Email}})</p>
            <p>{{.CommentText}}</p>
            <HR style=" border:1 dashed #987cb9" width="100%" color=#987cb9 SIZE=1>
        </div>
        {{end}}
    </div>

    <h2>Add Comment</h2>
    <form id="form-newcomment">
        <p><input hidden="true" type="text" name="page_id" value="{{.ID}}"></p>
        <p><input placeholder="your name" type="text" name="name" required="required"></p>
        <p><input placeholder="your email" type="email" name="email" required="required"></p>
        <p><textarea rows="5" placeholder="input your comments here" name="content" required="required"></textarea></p>
        <p><input type="submit" name="submit"></p>
    </form>
    <script src="/assets/js/jquery-3.1.1.js"></script>
</body>

<script>
    $('form#form-newcomment').submit(
        function addComment(e) {
            e.preventDefault();
            let form = $('form#form-newcomment');
            $.ajax({
                type: "POST",
                url: "/api/comments",
                data: form.serialize(),
                success: function(msg) {
                    //alert(msg.Fields.isAdded)
                    if (msg.Fields.isAdded.toLowerCase() != "true") {
                        alert("Failed to comment.")
                        return false;
                    }
                    let params = {};
                    $.each(form.serializeArray(), function(i, field) {
                        params[field.name] = field.value;
                    });


                    $("#comment-list").append('<div id="comment-node comment-node-shown"><p>Commented by ' + params['name'] + '(' + params['email'] + ')</p><p>' + params['content'] + '</p><HR style=" border:1 dashed #987cb9" width="100%" color=#987cb9 SIZE=1></div></div>');
                }
            });
            return;
        });
</script>

</html>