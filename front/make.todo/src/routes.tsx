import { Navigate, useRoutes } from 'react-router-dom';
import AuthorizationLayout from './layouts/AuthorizationLayout';
import DashboardLayout from './layouts/MainLayout';
import Dashboard from './pages/Dashboard';
import Login from './pages/Login';
import Register from './pages/Register';

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
const Router = () =>
  useRoutes([
    {
      path: '/app',
      element: <DashboardLayout />,
      children: [{ path: '/app', element: <Dashboard /> }],
    },
    {
      path: '/',
      element: <AuthorizationLayout />,
      children: [
        { path: '/', element: <Navigate to="/app" /> },
        { path: 'login', element: <Login /> },
        { path: 'register', element: <Register /> },
        { path: '*', element: <Navigate to="/404" /> },
      ],
    },
    { path: '*', element: <Navigate to="/404" replace /> },
  ]);

export default Router;
