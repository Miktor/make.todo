import API from '.';

export interface UserLoginResponse {
  token: string;
}

export interface UserLoginRequest {
  login: string;
  password: string;
}

export interface RegisterRequest {
  EmailHash: string;
  PasswordHash: string;
}

export abstract class Auth extends API {
  private static prefix = 'auth';

  public static async login(request: UserLoginRequest): Promise<Response> {
    return Auth.apiCall(`${Auth.prefix}/login`, request);
  }

  public static async register(request: RegisterRequest): Promise<Response> {
    return Auth.apiCall(`${Auth.prefix}/register`, request);
  }
}
