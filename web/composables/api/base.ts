import type { resType } from '~/types/api'
import type { AuthResponse } from '~/types/response/base'
export function useBaseApi() {
  // 用户登录
  const login = (formData: { user_name: string; password: string }) => {

    return useApiFetch<AuthResponse>({
      method: "post",
      url: '/base/login',
      data: formData
    })
  }
  // 测试jwt
  const ping = () => {
    return useApiFetch<resType>({
      method: "post",
      url: '/ping1',
    })
  }
  return {
    login,
    ping
  }
}