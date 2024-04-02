export enum APICommandMethod {
  ERROR       = "error",
	SHOW_GET    = "show_get",
	SHOW_SET    = "show_set",
	SHOW_LOAD   = "show_load",
  SHOW_LIST   = "show_list",
  SHOW_SAVE   = "show_save",
  SHOW_SAVEAS = "show_saveas"
}

export interface APIRequest {
  method: APICommandMethod;
  path: string;
  data: string;
}
