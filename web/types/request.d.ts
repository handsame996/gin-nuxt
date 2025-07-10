export enum ApiMethod {
  POST = "post",
  GET = "get",
  DELETE = "delete",
  PUT = "put",
  PATCH = "patch"
}

export type HttpHeaders = {
  [key: string]: string;
};

export interface ApiParams {
  method: ApiMethod;
  headers?: HttpHeaders;
  url: string;
  params?: Record<string, any>;
  data?: any;
  timeout?: number;
  withCredentials?: boolean;
  responseType?: 'json' | 'text' | 'blob' | 'arrayBuffer';
  baseURL?: string;
  validateStatus?: (status: number) => boolean;
}

export interface resType <T = any>{
  code: number;
  data: T;
  msg: string;
  err: string;
}