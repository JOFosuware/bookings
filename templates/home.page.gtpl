{{template "base" .}}

{{define "content"}}
    <div class="container">
        <div class="row">
            <div class="col">
                <h1>Welcome to the homepage</h1>
                <p>This is the homepage text</p>
                <p>
                    {{index .StringMap "test"}}
                </p>

                
            </div>
        </div>
    </div>
{{end}}