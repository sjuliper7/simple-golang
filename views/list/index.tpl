<div class="container">
    <div class="row">
        <div class="col-md-12">
            <h1>Todo List</h1>
        </div>
    </div>

    <div class="row">
        <div class="col-md-12">

            <table class="table table-hover">
                <thead>
                    <tr>
                        <th>Todo</th>
                        <th rowspan="2">Action</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .list}}
                    <tr>
                        <td><a href="/list-view/{{.Id}}">{{.Title}}</a>
                            from {{.Created}}
                        </td>
                        <td>
                            <a href="/list-edit/{{.Id}}" class="btn btn-default">Edit</a>
                        </td>
                        <td>
                            <a href="/list-delete/{{.Id}}" class="btn btn-danger">Delete</a>
                        </td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
    </div>
</div>