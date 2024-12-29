import { styled } from '@mui/material/styles';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
// import MenuIcon from '@mui/icons-material/Menu';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
// import SearchIcon from '@mui/icons-material/Search';
import MoreIcon from '@mui/icons-material/MoreVert';

const StyledToolbar = styled(Toolbar)(({ theme }) => ({
    alignItems: 'flex-start',
    paddingTop: theme.spacing(1),
    paddingBottom: theme.spacing(2),
    // Override media queries injected by theme.mixins.toolbar
    '@media all': {
        minHeight: 80,
    },
}));

export default function Header() {
    return (
        <Box
            position="fixed"
            sx={{
            flexGrow: 1 ,width: "100%",
            gap: { xs: 4, sm: 8 },
            top: 0,
            zIndex: 100,
        }}>
            <AppBar position="static">
                <StyledToolbar>
                    <IconButton
                        size="large"
                        edge="start"
                        color="inherit"
                        aria-label="open drawer"
                        sx={{ mr: 2 }}
                        onClick={() => {
                            window.location.href = "/";
                        }}
                    >
                        <ArrowBackIcon
                            sx={{
                                fontSize: 40,
                            }}
                        />
                    </IconButton>
                    <Typography
                        variant="h3"
                        noWrap
                        component="div"
                        sx={{
                            flexGrow: 1,
                            alignSelf: 'flex-end',
                            fontFamily:"fantasy"
                    }}
                    >
                        Let's GoodVS!
                    </Typography>
                    <IconButton size="large" aria-label="search" color="inherit"
                        onClick={() => {
                            if(localStorage.getItem('userID') === null) {
                                window.location.href = "/signin";
                            }else {
                                window.location.href = "/user";
                            }
                        }}
                    >
                        <AccountCircleIcon
                            sx={{
                                fontSize: 40,
                            }}
                        />
                    </IconButton>
                    <IconButton
                        size="large"
                        aria-label="display more actions"
                        edge="end"
                        color="inherit"
                    >
                        <MoreIcon
                            sx={{
                                fontSize: 40,
                            }}
                        />
                    </IconButton>
                </StyledToolbar>
            </AppBar>
        </Box>
    );
}