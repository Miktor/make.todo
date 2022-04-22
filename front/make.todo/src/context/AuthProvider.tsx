import { createContext } from 'react';

interface AuthContextInterface {
  authorized: boolean;
}

export const AuthContext = createContext<AuthContextInterface>(
  {} as AuthContextInterface
);

const defaultAuthContext: AuthContextInterface = {
  authorized: false,
};

type AuthProviderProps = {
  children: React.ReactNode;
};

export function AuthProvider({ children }: AuthProviderProps): JSX.Element {
  return (
    <AuthContext.Provider value={defaultAuthContext}>
      {children}
    </AuthContext.Provider>
  );
}
