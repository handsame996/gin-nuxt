export interface AuthResponse extends resType {
  token: string; // 这里不需要可选符，因为认证相关接口一定返回token
}
