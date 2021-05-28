// tslint:disable

import * as request from "superagent";
import {
    SuperAgentStatic,
    SuperAgentRequest,
    Response
} from "superagent";

export type RequestHeaders = {
    [header: string]: string;
}
export type RequestHeadersHandler = (headers: RequestHeaders) => RequestHeaders;

export type ConfigureAgentHandler = (agent: SuperAgentStatic) => SuperAgentStatic;

export type ConfigureRequestHandler = (agent: SuperAgentRequest) => SuperAgentRequest;

export type CallbackHandler = (err: any, res ? : request.Response) => void;

export type protobufAny = {
    'type_url' ? : string;
    'value' ? : string;
} & {
    [key: string]: any;
};

export type rpcStatus = {
    'code' ? : number;
    'message' ? : string;
    'details' ? : Array < protobufAny >
    ;
} & {
    [key: string]: any;
};

export type v2DeleteEntryResponse = {
    'entry' ? : v2Entry;
} & {
    [key: string]: any;
};

export type v2Entry = {
    'id' ? : string;
    'text' ? : string;
    'creator_id' ? : string;
    'created_at' ? : string;
    'updated_at' ? : string;
    'delete_time' ? : string;
} & {
    [key: string]: any;
};

export type v2ListEntriesResponse = {
    'entries' ? : Array < v2Entry >
    ;
    'next_page_token' ? : string;
    'total_size' ? : number;
    'has_next_page' ? : boolean;
} & {
    [key: string]: any;
};

export type Logger = {
    log: (line: string) => any
};

export interface ResponseWithBody < S extends number, T > extends Response {
    status: S;
    body: T;
}

export type QueryParameters = {
    [param: string]: any
};

export interface CommonRequestOptions {
    $queryParameters ? : QueryParameters;
    $domain ? : string;
    $path ? : string | ((path: string) => string);
    $retries ? : number; // number of retries; see: https://github.com/visionmedia/superagent/blob/master/docs/index.md#retrying-requests
    $timeout ? : number; // request timeout in milliseconds; see: https://github.com/visionmedia/superagent/blob/master/docs/index.md#timeouts
    $deadline ? : number; // request deadline in milliseconds; see: https://github.com/visionmedia/superagent/blob/master/docs/index.md#timeouts
}

/**
 * 
 * @class NotebookClient
 * @param {(string)} [domainOrOptions] - The project domain.
 */
export class NotebookClient {

    private domain: string = "";
    private errorHandlers: CallbackHandler[] = [];
    private requestHeadersHandler ? : RequestHeadersHandler;
    private configureAgentHandler ? : ConfigureAgentHandler;
    private configureRequestHandler ? : ConfigureRequestHandler;

    constructor(domain ? : string, private logger ? : Logger) {
        if (domain) {
            this.domain = domain;
        }
    }

    getDomain() {
        return this.domain;
    }

    addErrorHandler(handler: CallbackHandler) {
        this.errorHandlers.push(handler);
    }

    setRequestHeadersHandler(handler: RequestHeadersHandler) {
        this.requestHeadersHandler = handler;
    }

    setConfigureAgentHandler(handler: ConfigureAgentHandler) {
        this.configureAgentHandler = handler;
    }

    setConfigureRequestHandler(handler: ConfigureRequestHandler) {
        this.configureRequestHandler = handler;
    }

    private request(method: string, url: string, body: any, headers: RequestHeaders, queryParameters: QueryParameters, form: any, reject: CallbackHandler, resolve: CallbackHandler, opts: CommonRequestOptions) {
        if (this.logger) {
            this.logger.log(`Call ${method} ${url}`);
        }

        const agent = this.configureAgentHandler ?
            this.configureAgentHandler(request.default) :
            request.default;

        let req = agent(method, url);
        if (this.configureRequestHandler) {
            req = this.configureRequestHandler(req);
        }

        req = req.query(queryParameters);

        if (this.requestHeadersHandler) {
            headers = this.requestHeadersHandler({
                ...headers
            });
        }

        req.set(headers);

        if (body) {
            req.send(body);

            if (typeof(body) === 'object' && !(body.constructor.name === 'Buffer')) {
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

        req.end((error, response) => {
            // an error will also be emitted for a 4xx and 5xx status code
            // the error object will then have error.status and error.response fields
            // see superagent error handling: https://github.com/visionmedia/superagent/blob/master/docs/index.md#error-handling
            if (error) {
                reject(error);
                this.errorHandlers.forEach(handler => handler(error));
            } else {
                resolve(response);
            }
        });
    }

    private convertParameterCollectionFormat < T > (param: T, collectionFormat: string | undefined): T | string {
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
    }

    Notebook_ListEntriesURL(parameters: {
        'pageSize' ? : number,
        'pageToken' ? : string,
    } & CommonRequestOptions): string {
        let queryParameters: QueryParameters = {};
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        if (parameters['pageSize'] !== undefined) {
            queryParameters['page_size'] = this.convertParameterCollectionFormat(
                parameters['pageSize'],
                ''
            );
        }

        if (parameters['pageToken'] !== undefined) {
            queryParameters['page_token'] = this.convertParameterCollectionFormat(
                parameters['pageToken'],
                ''
            );
        }

        if (parameters.$queryParameters) {
            queryParameters = {
                ...queryParameters,
                ...parameters.$queryParameters
            };
        }

        let keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(key => key + '=' + encodeURIComponent(queryParameters[key])).join('&')) : '');
    }

    /**
     * 
     * @method
     * @name NotebookClient#Notebook_ListEntries
     * @param {integer} pageSize - 
     * @param {string} pageToken - 
     */
    Notebook_ListEntries(parameters: {
        'pageSize' ? : number,
        'pageToken' ? : string,
    } & CommonRequestOptions): Promise < ResponseWithBody < 200, v2ListEntriesResponse > | ResponseWithBody < number, rpcStatus >> {
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        let body: any;
        let queryParameters: QueryParameters = {};
        let headers: RequestHeaders = {};
        let form: any = {};
        return new Promise((resolve, reject) => {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';

            if (parameters['pageSize'] !== undefined) {
                queryParameters['page_size'] = this.convertParameterCollectionFormat(
                    parameters['pageSize'],
                    ''
                );
            }

            if (parameters['pageToken'] !== undefined) {
                queryParameters['page_token'] = this.convertParameterCollectionFormat(
                    parameters['pageToken'],
                    ''
                );
            }

            if (parameters.$queryParameters) {
                queryParameters = {
                    ...queryParameters,
                    ...parameters.$queryParameters
                };
            }

            this.request('GET', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    }

    Notebook_CreateEntryURL(parameters: {
        'body': v2Entry,
    } & CommonRequestOptions): string {
        let queryParameters: QueryParameters = {};
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        if (parameters.$queryParameters) {
            queryParameters = {
                ...queryParameters,
                ...parameters.$queryParameters
            };
        }

        queryParameters = {};

        let keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(key => key + '=' + encodeURIComponent(queryParameters[key])).join('&')) : '');
    }

    /**
     * 
     * @method
     * @name NotebookClient#Notebook_CreateEntry
     * @param {} body - 
     */
    Notebook_CreateEntry(parameters: {
        'body': v2Entry,
    } & CommonRequestOptions): Promise < ResponseWithBody < 200, v2Entry > | ResponseWithBody < number, rpcStatus >> {
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        let body: any;
        let queryParameters: QueryParameters = {};
        let headers: RequestHeaders = {};
        let form: any = {};
        return new Promise((resolve, reject) => {
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
                queryParameters = {
                    ...queryParameters,
                    ...parameters.$queryParameters
                };
            }

            form = queryParameters;
            queryParameters = {};

            this.request('POST', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    }

    Notebook_GetEntryURL(parameters: {
        'id': string,
    } & CommonRequestOptions): string {
        let queryParameters: QueryParameters = {};
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        path = path.replace(
            '{id}',
            `${encodeURIComponent(this.convertParameterCollectionFormat(
                        parameters['id'],
                        ''
                    ).toString())}`
        );

        if (parameters.$queryParameters) {
            queryParameters = {
                ...queryParameters,
                ...parameters.$queryParameters
            };
        }

        let keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(key => key + '=' + encodeURIComponent(queryParameters[key])).join('&')) : '');
    }

    /**
     * 
     * @method
     * @name NotebookClient#Notebook_GetEntry
     * @param {string} id - 
     */
    Notebook_GetEntry(parameters: {
        'id': string,
    } & CommonRequestOptions): Promise < ResponseWithBody < 200, v2Entry > | ResponseWithBody < number, rpcStatus >> {
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        let body: any;
        let queryParameters: QueryParameters = {};
        let headers: RequestHeaders = {};
        let form: any = {};
        return new Promise((resolve, reject) => {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';

            path = path.replace(
                '{id}',
                `${encodeURIComponent(this.convertParameterCollectionFormat(
                    parameters['id'],
                    ''
                ).toString())}`
            );

            if (parameters['id'] === undefined) {
                reject(new Error('Missing required  parameter: id'));
                return;
            }

            if (parameters.$queryParameters) {
                queryParameters = {
                    ...queryParameters,
                    ...parameters.$queryParameters
                };
            }

            this.request('GET', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    }

    Notebook_DeleteEntryURL(parameters: {
        'id': string,
    } & CommonRequestOptions): string {
        let queryParameters: QueryParameters = {};
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        path = path.replace(
            '{id}',
            `${encodeURIComponent(this.convertParameterCollectionFormat(
                        parameters['id'],
                        ''
                    ).toString())}`
        );

        if (parameters.$queryParameters) {
            queryParameters = {
                ...queryParameters,
                ...parameters.$queryParameters
            };
        }

        let keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(key => key + '=' + encodeURIComponent(queryParameters[key])).join('&')) : '');
    }

    /**
     * 
     * @method
     * @name NotebookClient#Notebook_DeleteEntry
     * @param {string} id - 
     */
    Notebook_DeleteEntry(parameters: {
        'id': string,
    } & CommonRequestOptions): Promise < ResponseWithBody < 200, v2DeleteEntryResponse > | ResponseWithBody < number, rpcStatus >> {
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        let body: any;
        let queryParameters: QueryParameters = {};
        let headers: RequestHeaders = {};
        let form: any = {};
        return new Promise((resolve, reject) => {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';

            path = path.replace(
                '{id}',
                `${encodeURIComponent(this.convertParameterCollectionFormat(
                    parameters['id'],
                    ''
                ).toString())}`
            );

            if (parameters['id'] === undefined) {
                reject(new Error('Missing required  parameter: id'));
                return;
            }

            if (parameters.$queryParameters) {
                queryParameters = {
                    ...queryParameters,
                    ...parameters.$queryParameters
                };
            }

            this.request('DELETE', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    }

    Notebook_UpdateEntryURL(parameters: {
        'id': string,
        'body': v2Entry,
    } & CommonRequestOptions): string {
        let queryParameters: QueryParameters = {};
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        path = path.replace(
            '{id}',
            `${encodeURIComponent(this.convertParameterCollectionFormat(
                        parameters['id'],
                        ''
                    ).toString())}`
        );

        if (parameters.$queryParameters) {
            queryParameters = {
                ...queryParameters,
                ...parameters.$queryParameters
            };
        }

        let keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(key => key + '=' + encodeURIComponent(queryParameters[key])).join('&')) : '');
    }

    /**
     * 
     * @method
     * @name NotebookClient#Notebook_UpdateEntry
     * @param {string} id - 
     * @param {} body - 
     */
    Notebook_UpdateEntry(parameters: {
        'id': string,
        'body': v2Entry,
    } & CommonRequestOptions): Promise < ResponseWithBody < 200, v2Entry > | ResponseWithBody < number, rpcStatus >> {
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries/{id}';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        let body: any;
        let queryParameters: QueryParameters = {};
        let headers: RequestHeaders = {};
        let form: any = {};
        return new Promise((resolve, reject) => {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';

            path = path.replace(
                '{id}',
                `${encodeURIComponent(this.convertParameterCollectionFormat(
                    parameters['id'],
                    ''
                ).toString())}`
            );

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
                queryParameters = {
                    ...queryParameters,
                    ...parameters.$queryParameters
                };
            }

            this.request('PATCH', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    }

    Notebook_UndeleteEntryURL(parameters: {
        'id': string,
    } & CommonRequestOptions): string {
        let queryParameters: QueryParameters = {};
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries/{id}:undelete';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        path = path.replace(
            '{id}',
            `${encodeURIComponent(this.convertParameterCollectionFormat(
                        parameters['id'],
                        ''
                    ).toString())}`
        );

        if (parameters.$queryParameters) {
            queryParameters = {
                ...queryParameters,
                ...parameters.$queryParameters
            };
        }

        queryParameters = {};

        let keys = Object.keys(queryParameters);
        return domain + path + (keys.length > 0 ? '?' + (keys.map(key => key + '=' + encodeURIComponent(queryParameters[key])).join('&')) : '');
    }

    /**
     * 
     * @method
     * @name NotebookClient#Notebook_UndeleteEntry
     * @param {string} id - 
     */
    Notebook_UndeleteEntry(parameters: {
        'id': string,
    } & CommonRequestOptions): Promise < ResponseWithBody < 200, v2Entry > | ResponseWithBody < number, rpcStatus >> {
        const domain = parameters.$domain ? parameters.$domain : this.domain;
        let path = '/v2/entries/{id}:undelete';
        if (parameters.$path) {
            path = (typeof(parameters.$path) === 'function') ? parameters.$path(path) : parameters.$path;
        }

        let body: any;
        let queryParameters: QueryParameters = {};
        let headers: RequestHeaders = {};
        let form: any = {};
        return new Promise((resolve, reject) => {
            headers['accept'] = 'application/json';
            headers['content-type'] = 'application/json';

            path = path.replace(
                '{id}',
                `${encodeURIComponent(this.convertParameterCollectionFormat(
                    parameters['id'],
                    ''
                ).toString())}`
            );

            if (parameters['id'] === undefined) {
                reject(new Error('Missing required  parameter: id'));
                return;
            }

            if (parameters.$queryParameters) {
                queryParameters = {
                    ...queryParameters,
                    ...parameters.$queryParameters
                };
            }

            form = queryParameters;
            queryParameters = {};

            this.request('POST', domain + path, body, headers, queryParameters, form, reject, resolve, parameters);
        });
    }

}

export default NotebookClient;
