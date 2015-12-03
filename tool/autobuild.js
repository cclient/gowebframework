var watch = require('glob-watcher');
var exec = require('child_process').exec;
var sys = require('sys');
var matchregs = new Array(9);
var godirpath = process.env.GODIR;
var command = process.env.COMMAND;
for (var i = 0; i < 10; i++) {
    var perkey = "";
    var tokay = "";
    for (var j = 0; j <= i; j++) {
        perkey += "*";
        tokay = perkey + "/" + tokay
    }
    matchregs[i] = godirpath + "/" + tokay.substring(0, tokay.length - 1);
}
function startwatch() {
    var w = watch(matchregs, function(evt) {
        console.log("1")
        child = exec(command, function(error, stdout, stderr) {
            sys.print('error: \n' + error + "\n");
            sys.print('stdout: \n' + stdout + "\n");
            sys.print('stderr: \n' + stderr + "\n");
        })
        setTimeout(function() {
            w.end();
            setTimeout(function() {
                startwatch();
            }, 1000)
        }, 2000)
    })
}

startwatch()
