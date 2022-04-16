import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import LayersIcon from '@mui/icons-material/Layers';

export const mainListItems = (
  <ListItemButton>
    <ListItemIcon>
      <LayersIcon />
    </ListItemIcon>
    <ListItemText primary="Todos" />
  </ListItemButton>
);

export const secondaryListItems = (
  <>
    <ListItemButton>
      <ListItemText primary="Log In" />
    </ListItemButton>
    <ListItemButton>
      <ListItemText primary="Log Out" />
    </ListItemButton>
  </>
);
