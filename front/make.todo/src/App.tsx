import { CssBaseline } from '@mui/material';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { ReactElement } from 'react';
import Router from './routes';

const mdTheme = createTheme();

function App(): ReactElement {
  return (
    <ThemeProvider theme={mdTheme}>
      <CssBaseline />
      <Router />
    </ThemeProvider>
  );
}

export default App;
