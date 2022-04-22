import { Navigate, Route, Routes } from 'react-router-dom';
import { RequireAuth } from './components/RequireAuth';
import AuthorizationLayout from './layouts/AuthorizationLayout';
import DashboardLayout from './layouts/MainLayout';
import Dashboard from './pages/Dashboard';
import Login from './pages/Login';
import Register from './pages/Register';

export default function Router(): React.ReactElement {
  return (
    <Routes>
      <Route path="/app" element={<DashboardLayout />}>
        <Route element={<RequireAuth />}>
          <Route path="" element={<Dashboard />} />
        </Route>
      </Route>

      <Route path="/" element={<AuthorizationLayout />}>
        <Route path="" element={<Navigate to="/app" />} />
        <Route path="login" element={<Login />} />
        <Route path="register" element={<Register />} />
      </Route>

      <Route path="*" element={<Navigate to="/404" replace />} />
    </Routes>
  );
}
