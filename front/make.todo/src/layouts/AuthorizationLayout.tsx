import * as React from 'react';
import { Container, CssBaseline } from '@mui/material';
import { Outlet } from 'react-router-dom';
import Copyright from '../components/Copyright';

export default function AuthorizationLayout(): React.ReactElement {
  return (
    <Container
      component="main"
      sx={{
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      <Outlet />
      <Copyright sx={{ mt: 5 }} />
    </Container>
  );
}
