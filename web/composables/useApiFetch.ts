import type { FetchError } from 'ofetch'
import { message } from 'ant-design-vue';
import type { ApiParams, resType } from '~/types/request'



export function useApiFetch<T = any>(params: ApiParams) {
  const config = useRuntimeConfig()
  const baseURL = params.baseURL || config.public.apiBase
  
  // 构建完整URL（包括查询参数）
  const fullUrl = new URL(params.url, baseURL)
  if (params.params) {
    Object.entries(params.params).forEach(([key, value]) => {
      fullUrl.searchParams.append(key, String(value))
    })
  }

  // 定义错误处理函数并明确类型
  const onResponseError = ({ response }: { response: FetchError<resType<T>> }) => {
    console.error(`API请求错误 [${response.status}]`, response.data?.msg)
    message.error(`API请求错误 [${response.status}] ${response.data?.msg}`)
    if (response.status === 401) {
      message.warning('用户未授权，需要重新登录')
      // 跳转到登录页
      navigateTo('/login')
    }
  }

  // 合并默认选项，使用更精确的类型
  const fetchOptions = {
    method: params.method,
    headers: {
      'Content-Type': 'application/json',
      ...params.headers
    },
    body: params.data,
    timeout: params.timeout || 60000,
    responseType: params.responseType || 'json',
    onResponseError,
  }

  // 使用resType<T>作为返回类型
  return useFetch<resType<T>>(fullUrl.toString(), fetchOptions)
}