export interface APIResponse {
  error: string | null;
  [key: string]: any;
}

export interface ParamMap {
  [key: string]: any;
}

export async function requestAPI(
  url: string,
  params?: ParamMap
): Promise<APIResponse> {
  const requestURL = params
    ? `/api${url}?` + new URLSearchParams(params)
    : `/api${url}`;
  const response = await fetch(requestURL);
  const responseJSON = (await response.json()) as APIResponse;
  return responseJSON;
}

export async function requestAPIForm(
  url: string,
  formData: FormData,
  keys: string[]
): Promise<APIResponse> {
  let params: ParamMap = {};
  for (const key of keys) {
    params[key] = formData.get(key);
  }

  return await requestAPI(url, params);
}
