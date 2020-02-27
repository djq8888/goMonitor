var localIP = "127.0.0.1"

function showMonitors() {
    var xhr=new XMLHttpRequest();
    xhr.onreadystatechange=function () {
        if (xhr.readyState==4){
            //alert(xhr.responseText);
            //document.getElementById("showBox").innerHTML = xhr.responseText;
            var filelist = new Array();
            filelist = xhr.responseText.split("\r\n");
            var select = document.getElementById('files');
            var options = "<option>==select one==</option>";
            for(var i=0; i<filelist.length; i++){
                options +='<option>'+filelist[i]+'</option>';
            }
            console.log(options)
            select.innerHTML = options;
        }
    };
    xhr.open('get','http://'+localIP+':8080/showMonitors');
    xhr.send(null);
}

function showCPU(){
    var select = document.getElementById('files');
    if (select.options.length == 0)
    {
        alert("请先获取文件列表并选择文件！");
        return false;
    }
    var index = select.selectedIndex;
    var filename = select.options[index].text;
    if (filename == "==select one==")
    {
        alert("请先选择文件!");
        return false
    }
    var url = "http://"+localIP+":8080/getCPU?name="+filename;
    var xhr = new XMLHttpRequest();
    var res = new Array();
    xhr.onreadystatechange=function () {
        if (xhr.readyState==4){
            //alert(xhr.responseText);
            //document.getElementById("files").innerHTML = xhr.responseText;
            res = xhr.responseText.split("\r\n");
            var xAxis = new Array();
            var yAxis = new Array();
            var avg = 0, sum = 0, max = 0;
            var line = document.getElementById("cpu");
            console.log(res.length)
            for (var i=0, j=0; j<res.length - 1 ;i++, j+=Math.floor(res.length/1800) + 1)
            {
                xAxis[i] = i + 1;
                yAxis[i] = Number(res[j+1])
                sum += yAxis[i]
                if (yAxis[i] > max)
                    max = yAxis[i];
            }
            console.log(xAxis.length)
            avg = sum / (xAxis.length - 1);
            var statistic = "avg is:" + avg.toString() + "\r\n" + "max is:" + max.toString();
            document.getElementById("showBox").innerHTML = statistic;
            var datas = {
                labels: xAxis,//标签
                values: yAxis,//值
                txtSet: {//绘制文本设置
                    txtfont: "14px microsoft yahei",
                    txtalgin: "center",
                    txtbaseline: "middle",
                    txtColor:"#000000"
                },
                bgSet:{//绘制背景线设置
                    lineColor:"#C0C0C0",
                    lineWidth:1,

                },
                lineColor:"#000000",//折线颜色
                circleColor:"black",//折线上原点颜色
                yAxis:{//y轴表示什么，及绘制文本的位置
                    x:50,
                    y:11,
                    title:"CPU利用率"
                }
            };
            lineChart(line,datas);//画折线图
        }
    };
    xhr.open('get',url);
    xhr.send(null);
}

function showMEM(){
    var select = document.getElementById('files');
    if (select.options.length == 0)
    {
        alert("请先获取文件列表并选择文件！");
        return false;
    }
    var index = select.selectedIndex;
    var filename = select.options[index].text;
    if (filename == "==select one==")
    {
        alert("请先选择文件!");
        return false
    }
    var url = "http://"+localIP+":8080/getMEM?name="+filename;
    var xhr = new XMLHttpRequest();
    var res = new Array();
    xhr.onreadystatechange=function () {
        if (xhr.readyState==4){
            //alert(xhr.responseText);
            //document.getElementById("files").innerHTML = xhr.responseText;
            res = xhr.responseText.split("\r\n");
            var xAxis = new Array();
            var yAxis = new Array();
            var avg = 0, sum = 0, max = 0;
            var line = document.getElementById("mem");
            console.log(res.length)
            for (var i=0, j=0; j<res.length - 1 ;i++, j+=Math.floor(res.length/1800) + 1)
            {
                xAxis[i] = i + 1;
                yAxis[i] = Number(res[j+1])
                sum += yAxis[i]
                if (yAxis[i] > max)
                    max = yAxis[i];
            }
            console.log(xAxis.length)
            avg = sum / (xAxis.length - 1);
            var statistic = "avg is:" + avg.toString() + "\r\n" + "max is:" + max.toString();
            document.getElementById("showBox").innerHTML = statistic;
            var datas = {
                labels: xAxis,//标签
                values: yAxis,//值
                txtSet: {//绘制文本设置
                    txtfont: "14px microsoft yahei",
                    txtalgin: "center",
                    txtbaseline: "middle",
                    txtColor:"#000000"
                },
                bgSet:{//绘制背景线设置
                    lineColor:"#C0C0C0",
                    lineWidth:1,

                },
                lineColor:"#000000",//折线颜色
                circleColor:"black",//折线上原点颜色
                yAxis:{//y轴表示什么，及绘制文本的位置
                    x:50,
                    y:11,
                    title:"內存占用率"
                }
            };
            lineChart(line,datas);//画折线图
        }
    };
    xhr.open('get',url);
    xhr.send(null);
}

function lineChart(elem, data) {
    if (elem.getContext) {
        var ctx = elem.getContext("2d"),
            labels = data.labels,//数值对应标签
            values = data.values,//数值
            len = labels.length,//标签/数值个数
            elemWidth = elem.width,//画布宽度
            elemHeight = elem.height,//画布高度
            gridHeight = Math.ceil(elemHeight / 5),//每行之间高度
            gridWidth = Math.floor(elemWidth / len),//每列之间看度
            actualHeight = 4 * gridHeight + 20;//绘制区域实际高度
        //设置绘制直线的属性
        ctx.strokeStyle = data.bgSet.lineColor;
        ctx.lineWidth = data.bgSet.lineWidth;
        //设置绘制文本的属性
        ctx.font = data.txtSet.txtfont;
        ctx.textAlign = data.txtSet.txtalgin;
        ctx.txtbaseline = data.txtSet.txtbaseline;
        //绘制背景
        //绘制背景横线
        ctx.beginPath();
        for (var i = 0; i < 5; i++) {
            var hgridY = gridHeight * i + 20,
                hgridX = gridWidth * len;
            ctx.moveTo(0, hgridY);
            ctx.lineTo(hgridX, hgridY);
        }
        ctx.stroke();

        //绘制背景的竖线，表示每个label
        // ctx.beginPath();
        // for (var j = 0; j < len + 1; j++) {
        //     var vgridX = gridWidth * j,
        //         vgridY = actualHeight;
        //     ctx.moveTo(vgridX, vgridY);
        //     ctx.lineTo(vgridX, vgridY + 10);
        // }
        // ctx.stroke();
        //绘制x轴标签
        ctx.fillStyle = data.txtSet.txtColor;
        for (var k = 0; k < len; k += Math.round(len/16)+1) {
            var txtX = gridWidth * (k + 0.5),
                txtY = actualHeight + 15;
            ctx.fillText(labels[k], txtX, txtY);
        }
        ctx.fill();

        //绘制y轴标签
        ctx.beginPath();
        for (var i = 0; i < 5; i++) {
            var hgridY = gridHeight * i + 20,
                hgridX = gridWidth * len;
            ctx.fillText(Math.round(25 * (4-i)), hgridX-15, hgridY);
        }

        //将数据与坐标对应
        var cData = new Array();
        for (var i = 0; i < len; i++) {
            cData[i] = values[i] / 25 * gridHeight;
        }

        ctx.stroke();
        //绘制折线
        ctx.strokeStyle = data.lineColor;
        ctx.beginPath();
        var pointX = gridWidth / 2,
            pointY = actualHeight - cData[0];
        ctx.moveTo(pointX, pointY);
        for (var i = 1; i < len; i++) {
            pointX += gridWidth;
            pointY = actualHeight - cData[i];
            ctx.lineTo(pointX, pointY);
        }
        ctx.stroke();
        //绘制坐标圆形
        ctx.beginPath();
        ctx.fillStyle = data.circleColor; //圆点的颜色
        for (var i = 0; i < len; i++) {
            var circleX = gridWidth / 2 + gridWidth * i,
                circleY = actualHeight - cData[i];
            ctx.moveTo(circleX, circleY); //假如不每次绘制之前确定开始绘制新路径，可以每次绘制之前移动到新的圆心
            ctx.arc(circleX, circleY, 1, 0, Math.PI * 2, false);
        }
        ctx.fill();
        //绘制坐标圆形对应的值
        // ctx.beginPath();
        // ctx.fillStyle = data.txtSet.txtColor;; //文本颜色
        // for (var i = 0; i < len; i++) {
        //     var circleX = gridWidth / 2 + gridWidth * i,
        //         circleY = actualHeight - cData[i];
        //     ctx.fillText(values[i], circleX, circleY - 8);
        // }
        // ctx.fill();
        //绘制y轴代表什么
        ctx.fillText(data.yAxis.title, data.yAxis.x, data.yAxis.y);
        ctx.fill();

    }
}