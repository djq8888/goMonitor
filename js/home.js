function showlog(){
    var url="http://127.0.0.1:8080/showLog?name="+document.getElementById('filename').value;
    window.location.href=url;
}

function parselog(){
    var url = "http://127.0.0.1:8080/parseLog?from="+document.getElementById('from').value+"&to="+document.getElementById('to').value;
    window.location.href=url;
}