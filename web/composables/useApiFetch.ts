import type { FetchError } from 'ofetch'
import { message } from 'ant-design-vue';
import type { ApiParams, resType } from '~/types/api'
import type { UseFetchOptions } from 'nuxt/app'

export function useApiFetch<T = any>(params: ApiParams) {
  const config = useRuntimeConfig()
  const baseURL = params.baseURL || config.public.apiBase

  var fullUrl = baseURL + params.url
  if (params.params) {
   const searchParams = new URLSearchParams();
    Object.entries(params.params).forEach(([key, value]) => {
    searchParams.append(key, String(value));
    });
    fullUrl += (fullUrl.includes('?') ? '&' : '?') + searchParams.toString();
  }

  // 定义错误处理函数并明确类型
  const onResponseError = ({ response }: { response: FetchError<resType<T>> }) => {
    message.error(`API请求错误 [${response.status}] ${response.data?.msg ? response.data?.msg :""}`)
    if (response.status === 401) {
      message.warning('用户未授权，需要重新登录')
      // 跳转到登录页
      navigateTo('/login')
    }
  }

  // 处理正确响应
  const onResponse = (ctx: { response: { _data: resType<T>,status: number } }) => {
    const { _data,status } = ctx.response;
    
    if (status >= 200 && status < 300) {
      if (_data.code !== 0) {
        console.log('发起请求:', fullUrl, '触发时间:', new Date().toISOString(), '请求体:', params.data);
        console.log('onResponse函数被调用，来源:', new Error().stack); // 打印调用栈
        message.warning(_data.msg || '请求失败');
      } else {
        message.success(_data.msg || '请求成功')
      }
    }
    
    // 可以在这里添加其他全局处理逻辑，如日志记录
    // console.log('[API Response]', _data);
    
    // 可选：可以直接修改响应数据，例如只返回data字段
    // ctx.response._data = _data.data;
  };

  // 合并默认选项，使用更精确的类型
  const fetchOptions: UseFetchOptions<resType<T>> = {
    method: params.method,
    headers: {
      'Content-Type': 'application/json',
      ...params.headers
    },
    body: params.data,
    timeout: params.timeout || 60000,
    responseType: params.responseType || 'json',
    watch:false,
    onResponseError,
    onResponse
  }

  // 使用resType<T>作为返回类型
  return useFetch<resType<T>>(fullUrl.toString(), fetchOptions)
}