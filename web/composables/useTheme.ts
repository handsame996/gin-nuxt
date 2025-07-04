export const useTheme = () => {
  const isDark = useState('isDark', () => {
    // 从本地存储读取偏好，默认浅色
    return localStorage.getItem('theme') === 'dark'
  })

  // 切换主题
  const toggleTheme = () => {
    isDark.value = !isDark.value
    const theme = isDark.value ? 'dark' : 'light'
    document.documentElement.classList.toggle('dark-theme', isDark.value)
    localStorage.setItem('theme', theme)
  }

  // 初始化主题
  onMounted(() => {
    if (isDark.value) {
      document.documentElement.classList.add('dark-theme')
    }
  })

  return { isDark, toggleTheme }
}