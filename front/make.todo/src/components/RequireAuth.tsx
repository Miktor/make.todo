import { useLocation, Navigate, Outlet } from 'react-router-dom';
import { Location } from 'history';
import useAuth from '../hooks/UseAuth';

export interface RequireAuthState {
  from: Location;
}

export function RequireAuth() {
  const auth = useAuth();
  const location = useLocation();

  const state: RequireAuthState = { from: location };
  return auth?.authorized ? (
    <Outlet />
  ) : (
    <Navigate to="/login" state={state} replace />
  );
}
