import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import BookOutlinedIcon from '@mui/icons-material/BookOutlined';
import Button from '@material-ui/core/Button';
import Box from '@material-ui/core/Box';
import { styled } from '@mui/material/styles';
import { blue } from '@mui/material/colors';

const ColorButton = styled(Button)(({ theme }) => ({
    color: theme.palette.getContrastText(blue[600]),
    width: 130,
    'cursor': 'pointer',
    backgroundColor: blue[600],
    '&:hover': {
        backgroundColor: blue[800],
    },
}));

async function Login() {
    let STATE = require('uuid').v4();
    let REDIRECT_URL = 'http://127.0.0.1:3001';
    let CLIENT_ID;
    if (process.env.REACT_APP_CLIENT_ID)
        CLIENT_ID = process.env.REACT_APP_CLIENT_ID;
    else
        throw new Error("CLIENT_ID environment variable is not set");
    window.location.href = `https://api.intra.42.fr/oauth/authorize?client_id=${CLIENT_ID}&redirect_uri=${REDIRECT_URL}&scope=public&state=${STATE}&response_type=code`;
}

export default function Header(props) {
    return (
        <Box sx={{
            mt: 1,
            pl: 1,
            pr: 1,
            display: "flex",
            flexDirection: 'row',
        }}>
            <Grid container direction="row">
                <Grid item >
                    <BookOutlinedIcon fontSize="large" />
                </Grid >
                <Grid item >
                    <Typography variant="h5" component="div" gutterBottom>
                        42diploma
                    </Typography>
                </Grid >
            </Grid >
            <Box sx={{
            }}>
                {props.user === null ?
                    <ColorButton variant="contained" size="large" onClick={Login}>Sign in</ColorButton>
                    :
                    <ColorButton variant="contained" size="large" onClick={() => {
                        props.setUser(null);
                        props.setProjectsDoable(null);
                    }}>Sign out</ColorButton>
                }
            </Box>
        </Box>
    );
}
