var localIP = "127.0.0.1"

function jump2analyse() {
    url = 'http://'+localIP+':8080/analyse'
    window.location.href=url;
}

function jump2monitor() {
    url = 'http://'+localIP+':8080/monitor'
    window.location.href=url;
}