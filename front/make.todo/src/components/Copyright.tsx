import * as React from 'react';
import Link from '@mui/material/Link';
import Typography from '@mui/material/Typography';
import { Box, SxProps, Theme } from '@mui/material';

interface CopyrightProps {
  sx: SxProps<Theme>;
}

export default function Copyright({
  sx = [],
}: CopyrightProps): React.ReactElement {
  return (
    <Typography variant="body2" color="text.secondary" align="center" sx={sx}>
      {'Copyright © '}
      <Link color="inherit" href="https://mui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}.
    </Typography>
  );
}
