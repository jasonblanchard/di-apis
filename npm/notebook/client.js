"use strict";
// tslint:disable
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
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
exports.__esModule = true;
exports.NotebookClient = void 0;
var request = __importStar(require("superagent"));
/**
 *
 * @class NotebookClient
 * @param {(string)} [domainOrOptions] - The project domain.
 */
var NotebookClient = /** @class */ (function () {
    function NotebookClient(domain, logger) {
        this.logger = logger;
        this.domain = "";
        this.errorHandlers = [];
        if (domain) {
            this.domain = domain;
        }
    }
    NotebookClient.prototype.getDomain = function () {
        return this.domain;
    };
    NotebookClient.prototype.addErrorHandler = function (handler) {
        this.errorHandlers.push(handler);
    };
    NotebookClient.prototype.setRequestHeadersHandler = function (handler) {
        this.requestHeadersHandler = handler;
    };
    NotebookClient.prototype.setConfigureAgentHandler = function (handler) {
        this.configureAgentHandler = handler;
    };
    NotebookClient.prototype.setConfigureRequestHandler = function (handler) {
        this.configureRequestHandler = handler;
    };
    NotebookClient.prototype.request = function (method, url, body, headers, queryParameters, form, reject, resolve, opts) {
        var _this = this;
        if (this.logger) {
            this.logger.log("Call " + method + " " + url);
        }
        var agent = this.configureAgentHandler ?
            this.configureAgentHandler(request["default"]) :
            request["default"];
        var req = agent(method, url);
        if (this.configureRequestHandler) {
            req = this.configureRequestHandler(req);
        }
        req = req.query(queryParameters);
        if (this.requestHeadersHandler) {
            headers = this.requestHeadersHandler(__assign({}, headers));
        }
        req.set(headers);
        if (body) {
            req.send(body);
            if (typeof (body) === 'object' && !(body.constructor.name === 'Buffer')) {
                headers['content-type'] = 'application/json';
            }
        }
        if (Object.keys(form).length > 0) {
            req.type('form');
            req.send(form);
        }
        if (opts.$retries && opts.$retries > 0) {
            req.retry(opts.$retries);
        }
        if (opts.$timeout && opts.$timeout > 0 || opts.$deadline && opts.$deadline > 0) {
            req.timeout({
                deadline: opts.$deadline,
                response: opts.$timeout
            });
        }
        req.end(function (error, response) {
            // an error will also be emitted for a 4xx and 5xx status code
            // the error object will then have error.status and error.response fields
            // see superagent error handling: https://github.com/visionmedia/superagent/blob/master/docs/index.md#error-handling
            if (error) {
                reject(error);
                _this.errorHandlers.forEach(function (handler) { return handler(error); });
            }
            else {
                resolve(response);
            }
        });
    };
    NotebookClient.prototype.convertParameterCollectionFormat = function (param, collectionFormat) {
        if (Array.isArray(param) && param.length >= 2) {
            switch (collectionFormat) {
                case "csv":
                    return param.join(",");
                case "ssv":
                    return param.join(" ");
                case "tsv":
                    return param.join("\t");
                case "pipes":
                    return param.join("|");
                default:
                    return param;
            }
        }
        return param;
    };
    NotebookClient.prototype.Notebook_ListEntriesURL = function (parameters) {
        var queryParameters = {};
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        if (parameters['pageSize'] !== undefined) {
            queryParameters['page_size'] = this.convertParameterCollectionFormat(parameters['pageSize'], '');
        }
        if (parameters['pageToken'] !== undefined) {
            queryParameters['page_token'] = this.convertParameterCollectionFormat(parameters['pageToken'], '');
        }
        if (parameters.$queryParameters) {
            queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
        }
        var keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(function (key) { return key + '=' + encodeURIComponent(queryParameters[key]); }).join('&')) : '');
    };
    /**
     *
     * @method
     * @name NotebookClient#Notebook_ListEntries
     * @param {integer} pageSize -
     * @param {string} pageToken -
     */
    NotebookClient.prototype.Notebook_ListEntries = function (parameters) {
        var _this = this;
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        var body;
        var queryParameters = {};
        var headers = {};
        var form = {};
        return new Promise(function (resolve, reject) {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';
            if (parameters['pageSize'] !== undefined) {
                queryParameters['page_size'] = _this.convertParameterCollectionFormat(parameters['pageSize'], '');
            }
            if (parameters['pageToken'] !== undefined) {
                queryParameters['page_token'] = _this.convertParameterCollectionFormat(parameters['pageToken'], '');
            }
            if (parameters.$queryParameters) {
                queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
            }
            _this.request('GET', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    };
    NotebookClient.prototype.Notebook_CreateEntryURL = function (parameters) {
        var queryParameters = {};
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        if (parameters.$queryParameters) {
            queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
        }
        queryParameters = {};
        var keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(function (key) { return key + '=' + encodeURIComponent(queryParameters[key]); }).join('&')) : '');
    };
    /**
     *
     * @method
     * @name NotebookClient#Notebook_CreateEntry
     * @param {} body -
     */
    NotebookClient.prototype.Notebook_CreateEntry = function (parameters) {
        var _this = this;
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        var body;
        var queryParameters = {};
        var headers = {};
        var form = {};
        return new Promise(function (resolve, reject) {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';
            if (parameters['body'] !== undefined) {
                body = parameters['body'];
            }
            if (parameters['body'] === undefined) {
                reject(new Error('Missing required  parameter: body'));
                return;
            }
            if (parameters.$queryParameters) {
                queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
            }
            form = queryParameters;
            queryParameters = {};
            _this.request('POST', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    };
    NotebookClient.prototype.Notebook_GetEntryURL = function (parameters) {
        var queryParameters = {};
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        path = path.replace('{id}', "" + encodeURIComponent(this.convertParameterCollectionFormat(parameters['id'], '').toString()));
        if (parameters.$queryParameters) {
            queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
        }
        var keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(function (key) { return key + '=' + encodeURIComponent(queryParameters[key]); }).join('&')) : '');
    };
    /**
     *
     * @method
     * @name NotebookClient#Notebook_GetEntry
     * @param {string} id -
     */
    NotebookClient.prototype.Notebook_GetEntry = function (parameters) {
        var _this = this;
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        var body;
        var queryParameters = {};
        var headers = {};
        var form = {};
        return new Promise(function (resolve, reject) {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';
            path = path.replace('{id}', "" + encodeURIComponent(_this.convertParameterCollectionFormat(parameters['id'], '').toString()));
            if (parameters['id'] === undefined) {
                reject(new Error('Missing required  parameter: id'));
                return;
            }
            if (parameters.$queryParameters) {
                queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
            }
            _this.request('GET', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    };
    NotebookClient.prototype.Notebook_DeleteEntryURL = function (parameters) {
        var queryParameters = {};
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        path = path.replace('{id}', "" + encodeURIComponent(this.convertParameterCollectionFormat(parameters['id'], '').toString()));
        if (parameters.$queryParameters) {
            queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
        }
        var keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(function (key) { return key + '=' + encodeURIComponent(queryParameters[key]); }).join('&')) : '');
    };
    /**
     *
     * @method
     * @name NotebookClient#Notebook_DeleteEntry
     * @param {string} id -
     */
    NotebookClient.prototype.Notebook_DeleteEntry = function (parameters) {
        var _this = this;
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        var body;
        var queryParameters = {};
        var headers = {};
        var form = {};
        return new Promise(function (resolve, reject) {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';
            path = path.replace('{id}', "" + encodeURIComponent(_this.convertParameterCollectionFormat(parameters['id'], '').toString()));
            if (parameters['id'] === undefined) {
                reject(new Error('Missing required  parameter: id'));
                return;
            }
            if (parameters.$queryParameters) {
                queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
            }
            _this.request('DELETE', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    };
    NotebookClient.prototype.Notebook_UpdateEntryURL = function (parameters) {
        var queryParameters = {};
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        path = path.replace('{id}', "" + encodeURIComponent(this.convertParameterCollectionFormat(parameters['id'], '').toString()));
        if (parameters.$queryParameters) {
            queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
        }
        var keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(function (key) { return key + '=' + encodeURIComponent(queryParameters[key]); }).join('&')) : '');
    };
    /**
     *
     * @method
     * @name NotebookClient#Notebook_UpdateEntry
     * @param {string} id -
     * @param {} body -
     */
    NotebookClient.prototype.Notebook_UpdateEntry = function (parameters) {
        var _this = this;
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        var body;
        var queryParameters = {};
        var headers = {};
        var form = {};
        return new Promise(function (resolve, reject) {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';
            path = path.replace('{id}', "" + encodeURIComponent(_this.convertParameterCollectionFormat(parameters['id'], '').toString()));
            if (parameters['id'] === undefined) {
                reject(new Error('Missing required  parameter: id'));
                return;
            }
            if (parameters['body'] !== undefined) {
                body = parameters['body'];
            }
            if (parameters['body'] === undefined) {
                reject(new Error('Missing required  parameter: body'));
                return;
            }
            if (parameters.$queryParameters) {
                queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
            }
            _this.request('PATCH', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    };
    NotebookClient.prototype.Notebook_UndeleteEntryURL = function (parameters) {
        var queryParameters = {};
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries/{id}:undelete';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        path = path.replace('{id}', "" + encodeURIComponent(this.convertParameterCollectionFormat(parameters['id'], '').toString()));
        if (parameters.$queryParameters) {
            queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
        }
        queryParameters = {};
        var keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(function (key) { return key + '=' + encodeURIComponent(queryParameters[key]); }).join('&')) : '');
    };
    /**
     *
     * @method
     * @name NotebookClient#Notebook_UndeleteEntry
     * @param {string} id -
     */
    NotebookClient.prototype.Notebook_UndeleteEntry = function (parameters) {
        var _this = this;
        var domain = parameters.$domain ? parameters.$domain : this.domain;
        var path = '/v2/entries/{id}:undelete';
        if (parameters.$path) {
            path = (typeof (parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }
        var body;
        var queryParameters = {};
        var headers = {};
        var form = {};
        return new Promise(function (resolve, reject) {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';
            path = path.replace('{id}', "" + encodeURIComponent(_this.convertParameterCollectionFormat(parameters['id'], '').toString()));
            if (parameters['id'] === undefined) {
                reject(new Error('Missing required  parameter: id'));
                return;
            }
            if (parameters.$queryParameters) {
                queryParameters = __assign(__assign({}, queryParameters), parameters.$queryParameters);
            }
            form = queryParameters;
            queryParameters = {};
            _this.request('POST', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    };
    return NotebookClient;
}());
exports.NotebookClient = NotebookClient;
exports["default"] = NotebookClient;
