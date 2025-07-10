import { useApiFetch } from '~/composables/useApiFetch'
import { ApiMethod, } from '~/types/request'
import type { resType } from '~/types/request' // 仅导入类型
export function useUserApi() {
  // 获取用户列表
  const fetchUserList = () => {
    return useApiFetch<{ users: resType }>({
    method: ApiMethod.POST, // 指定请求方法
    url: '/base/login'      // 指定请求URL
    // 其他可选参数根据需要添加
  })
  }

  // 用户登录
  const login = (formData: { username: string; password: string }) => {
    return useApiFetch<{ token: string }>({
      method: ApiMethod.POST,
      url: '/api/login',
      data: formData
    })
  }

  return {
    fetchUserList,
    login
  }
}