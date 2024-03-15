export enum APICommandMethod {
  ERROR      = "error",
	SHOW_GET   = "show_get",
	SHOW_SET   = "show_set",
	SHOW_LOAD  = "show_load"
}

export interface APIRequest {
  method: APICommandMethod;
  path: string;
  data: string;
}
