export default defineNuxtRouteMiddleware((to, from) => {
  // 获取 token（从 localStorage 或 cookies 中，根据实际存储方式调整）
  const token = useCookie('token').value || localStorage.getItem('token')

  // 定义无需登录即可访问的白名单路由（如登录页、注册页、404页等）
  const whiteList = ['/login', '/register', '/404']

  // 判断：若没有 token，且目标路由不在白名单中，则强制跳转到登录页
  if (!token && !whiteList.includes(to.path)) {
    return navigateTo('/login')
  }

  // 可选：若已登录却访问登录页，自动跳转到首页
  if (token && to.path === '/login') {
    return navigateTo('/admin')
  }
})