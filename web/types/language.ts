export const supportedLocales = ['en', 'zh'] as const
export type Locale = typeof supportedLocales[number]