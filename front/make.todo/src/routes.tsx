import {
  Routes, Route, Navigate, useRoutes,
} from 'react-router-dom';
import DashboardLayout from './layouts/MainLayout';
import Dashboard from './pages/Dashboard';
import SignIn from './pages/SignIn';

const Router = () => useRoutes([
  {
    path: '/app',
    element: <DashboardLayout />,
    children: [{ path: '/app', element: <Dashboard /> }],
  },
  {
    path: '/',
    children: [
      { path: '/', element: <Navigate to="/app" /> },
      { path: 'login', element: <SignIn /> },
      { path: '*', element: <Navigate to="/404" /> },
    ],
  },
  { path: '*', element: <Navigate to="/404" replace /> },
]);

export default Router;
