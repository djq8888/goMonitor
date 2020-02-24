var localIP = "127.0.0.1"

function showfiles() {
    var xhr=new XMLHttpRequest();
    xhr.onreadystatechange=function () {
        if (xhr.readyState==4){
            //alert(xhr.responseText);
            document.getElementById("files").innerHTML = xhr.responseText;
        }
    };
    xhr.open('get','http://'+localIP+':8080/showFiles');
    xhr.send(null);
}

function showlog(){
    var url = "http://"+localIP+":8080/showLog?name="+document.getElementById('filename').value;
    window.location.href=url;
}

function parselog(){
    var filename = document.getElementById('filename').value
    if (filename == "")
    {
        document.getElementById('filename').focus();
        alert("please input filename!");
        return false
    }
    var from = document.getElementById('from').value
    if (from == "")
    {
        document.getElementById('from').focus();
        alert("please input from!");
        return false
    }
    var url = "http://"+localIP+":8080/parseLog?name="+filename+"&from="+from;
    var to = document.getElementById('to').value
    if (to != "")
    {
        url += "&to="+document.getElementById('to').value
    }
    window.location.href=url;
}