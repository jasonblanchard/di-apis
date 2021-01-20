"use strict";
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
exports.__esModule = true;
var client_1 = require("../client");
var notebookClient = new client_1.NotebookClient("https://di3.blanktech.net/notebook");
notebookClient.setRequestHeadersHandler(function (headers) { return (__assign(__assign({}, headers), { "Authorization": "Bearer " + process.env.API_TOKEN })); });
notebookClient.Notebook_GetEntry({
    id: "8351"
}).then(function (response) {
    console.log("Body: ", response.body);
})["catch"](function (error) {
    console.log("Error: ", error.status, error.response.body);
});
