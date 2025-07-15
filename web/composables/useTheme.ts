export const useTheme = () => {
  const isDark = useState('isDark', () => {
    // 从本地存储读取偏好，默认浅色
    return localStorage.getItem('theme') === 'dark'
  })
  // 响应式触发器，用于强制更新 themeVars
  const themeVersion = ref(0)
  // 切换主题
  const toggleTheme = () => {
    isDark.value = !isDark.value
    const theme = isDark.value ? 'dark' : 'light'
    document.documentElement.classList.toggle('dark-theme', isDark.value)
    localStorage.setItem('theme', theme)
    themeVersion.value++
  }

  // 初始化主题
  onMounted(() => {
    if (isDark.value) {
      document.documentElement.classList.add('dark-theme')
    }
    themeVersion.value++
  })

  // 获取当前主题的所有变量
  const themeVars = computed(() => {
    themeVersion.value
    // 获取根元素的计算样式
    const styles = getComputedStyle(document.documentElement)
    
    // 提取需要的主题变量（与 less 中定义的对应）
    return {
      primaryColor: styles.getPropertyValue('--primary-color').trim() ,
      primaryLight: styles.getPropertyValue('--primary-light').trim(),
      primaryDark: styles.getPropertyValue('--primary-dark').trim(),
      bgColor: styles.getPropertyValue('--bg-color').trim(),
      textPrimary: styles.getPropertyValue('--text-primary').trim(),
      successColor: styles.getPropertyValue('--success-color').trim(),
      warningColor: styles.getPropertyValue('--warning-color').trim(),
      dangerColor: styles.getPropertyValue('--danger-color').trim(),
      textSecondary: styles.getPropertyValue('--text-secondary').trim(),
      borderColor: styles.getPropertyValue('--border-color').trim()
    }
  })

  return { isDark, themeVars, toggleTheme }
}



export const useAntTheme = () => {
  const isDark = useState('isDark', () => {
    // 从本地存储读取偏好，默认浅色
    return localStorage.getItem('theme') === 'dark'
  })
  // 响应式触发器，用于强制更新 themeVars
  const themeVersion = ref(0)
  // 切换主题
  const toggleTheme = () => {
    isDark.value = !isDark.value
    const theme = isDark.value ? 'dark' : 'light'
    document.documentElement.classList.toggle('dark-theme', isDark.value)
    localStorage.setItem('theme', theme)
    themeVersion.value++
  }

  // 初始化主题
  onMounted(() => {
    if (isDark.value) {
      document.documentElement.classList.add('dark-theme')
    }
    themeVersion.value++
  })

  // 获取当前主题的所有变量
  const themeVars = computed(() => {
    themeVersion.value
    // 获取根元素的计算样式
    const styles = getComputedStyle(document.documentElement)
    
    // 提取需要的主题变量（与 less 中定义的对应）
    
    return {
      // 主色调系列
    colorPrimary: styles.getPropertyValue('--primary-color').trim(),
    colorPrimaryLight: styles.getPropertyValue('--primary-light').trim(),
    colorPrimaryLighter: styles.getPropertyValue('--primary-lighter').trim(),
    colorPrimaryDark: styles.getPropertyValue('--primary-dark').trim(),
    colorPrimaryDarker: styles.getPropertyValue('--primary-darker').trim(),

    // 背景色系列
    colorBgContainer: styles.getPropertyValue('--bg-color').trim(),
    colorBgElevated: styles.getPropertyValue('--bg-secondary').trim(),
    colorBgTextActive: styles.getPropertyValue('--bg-tertiary').trim(),

    // 文本色系列
    colorText: styles.getPropertyValue('--text-primary').trim(),
    colorTextSecondary: styles.getPropertyValue('--text-secondary').trim(),
    colorTextTertiary: styles.getPropertyValue('--text-tertiary').trim(),
    colorTextDisabled: styles.getPropertyValue('--text-disabled').trim(),

    // 功能色系列（成功）
    colorSuccess: styles.getPropertyValue('--success-color').trim(),
    colorSuccessLight: styles.getPropertyValue('--success-light').trim(),
    colorSuccessDark: styles.getPropertyValue('--success-dark').trim(),

    // 功能色系列（警告）
    colorWarning: styles.getPropertyValue('--warning-color').trim(),
    colorWarningLight: styles.getPropertyValue('--warning-light').trim(),
    colorWarningDark: styles.getPropertyValue('--warning-dark').trim(),

    // 功能色系列（危险/错误）
    colorError: styles.getPropertyValue('--danger-color').trim(),
    colorErrorLight: styles.getPropertyValue('--danger-light').trim(),
    colorErrorDark: styles.getPropertyValue('--danger-dark').trim(),

    // 边框色系列
    colorBorder: styles.getPropertyValue('--border-color').trim(),
    colorBorderDark: styles.getPropertyValue('--border-dark').trim(),
    }
  })

  return { isDark, themeVars, toggleTheme }
}