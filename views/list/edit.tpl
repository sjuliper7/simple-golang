<div class="container">
    <h1>Edit {{.list.Title}}</h1>

    <h4>New Blog Post</h4>
    <form action="" method="post">
        <div class="form-group">
            <label>Title</label>
            <input class="form-control" type="text" name="title" value="{{.list.Title}}">
        </div>
        <div class="form-group">
            <label>Content</label>
            <textarea name="description" class="form-control" colspan="3" rowspan="10">{{.list.Description}}</textarea>
            <input type="hidden" name="id" value="{{.list.Id}}"><br>
            <input type="submit" value="Submit" class="btn btn-default">
        </div>
    </form>
</div>