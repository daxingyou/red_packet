(function (b, a, c) {
    b.danidemo = b.extend({}, {
        addLog: function (j, f, i) {
            var h = new Date();
            var e = b("<li />", {"class": "demo-" + f});
            var g = "[" + h.getHours() + ":" + h.getMinutes() + ":" + h.getSeconds() + "] ";
            g += i;
            e.html(g);
            b(j).prepend(e)
        }, addFile: function (g, e, d) {
            var f = '<div id="demo-file' + e + '"><img src="http://placehold.it/48.png" class="demo-image-preview" /><span class="demo-file-id">#' + e + "</span> - " + d.name + ' <span class="demo-file-size">(' + b.danidemo.humanizeSize(d.size) + ')</span><br />Status: <span class="demo-file-status">Waiting to upload</span><div class="progress progress-striped active"><div class="progress-bar" role="progressbar" style="width: 0%;"><span class="sr-only">0% Complete</span></div></div></div>';
            var e = b(g).attr("file-counter");
            if (!e) {
                b(g).empty();
                e = 0
            }
            e++;
            b(g).attr("file-counter", e);
            b(g).prepend(f)
        }, updateFileStatus: function (e, d, f) {
            b("#demo-file" + e).find("span.demo-file-status").html(f).addClass("demo-file-status-" + d)
        }, updateFileProgress: function (d, e) {
            b("#demo-file" + d).find("div.progress-bar").width(e);
            b("#demo-file" + d).find("span.sr-only").html(e + " Complete")
        }, humanizeSize: function (e) {
            var d = Math.floor(Math.log(e) / Math.log(1024));
            return (e / Math.pow(1024, d)).toFixed(2) * 1 + " " + ["B", "kB", "MB", "GB", "TB"][d]
        }
    }, b.danidemo)
})(jQuery, this);