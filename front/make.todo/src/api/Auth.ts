import API from '.';

export interface UserLoginResponse {
  token: string;
}

export interface UserLoginRequest {
  login: string;
  password: string;
}

export interface RegisterRequest {
  login: string;
  password: string;
}

export abstract class Auth extends API {
  private static prefix = 'auth';

  public static async login(
    request: UserLoginRequest
  ): Promise<UserLoginResponse> {
    return Auth.apiCall(`${Auth.prefix}/login`, request);
  }

  public static async register(
    request: RegisterRequest
  ): Promise<UserLoginResponse> {
    return Auth.apiCall(`${Auth.prefix}/register`, request);
  }
}
