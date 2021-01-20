var fs = require("fs");
var CodeGen = require("swagger-typescript-codegen").CodeGen;

var file = "gen/pb-go/notebook.swagger.json";
var swagger = JSON.parse(fs.readFileSync(file, "UTF-8"));
var tsSourceCode = CodeGen.getTypescriptCode({
    className: "NotebookClient",
    swagger: swagger,
});
console.log(tsSourceCode);