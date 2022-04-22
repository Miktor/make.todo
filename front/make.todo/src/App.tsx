import { CssBaseline } from '@mui/material';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { ReactElement } from 'react';
import { AuthProvider } from './context/AuthProvider';
import Router from './routes';

const mdTheme = createTheme();

function App(): ReactElement {
  return (
    <ThemeProvider theme={mdTheme}>
      <AuthProvider>
        <CssBaseline />
        <Router />
      </AuthProvider>
    </ThemeProvider>
  );
}

export default App;
