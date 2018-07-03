{{template "html/header.html"}}
<body>
    {{template "html/topnav.html"}}

    <div class="container">
        <div class="left">
            <a href="/list-new" class="btn btn-info">New List</a>
        </div>
    </div>

    {{.LayoutContent}}
</body>

</html>