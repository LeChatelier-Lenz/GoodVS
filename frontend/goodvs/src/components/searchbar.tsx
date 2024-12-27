import Paper from '@mui/material/Paper';
import InputBase from '@mui/material/InputBase';
import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import SearchIcon from '@mui/icons-material/Search';
import DirectionsIcon from '@mui/icons-material/Directions';
import {useRef} from "react";

export default function CustomizedInputBase() {
    const searchRef = useRef("");

    const handleChange = () => {
        const search = document.getElementById('product') as HTMLInputElement;
        searchRef.current = search.value;
        // console.log("search value: ",searchRef.current);
    }

    const handleSearch = () => {
        console.log("search value: ",searchRef.current);
    }

    return (
        <Paper
            component="form"
            sx={{ p: '2px 4px', display: 'flex', alignItems: 'center', width: 600,
                minWidth: 300, maxWidth: 800, height: 50, borderRadius: 5, margin: 'auto'

        }}
            onSubmit={handleSearch}
        >
            <IconButton sx={{ p: '20px' }} aria-label="menu">
                <MenuIcon />
            </IconButton>
            <InputBase
                sx={{ ml: 2, flex: 2 }}
                size="medium"
                placeholder="Search Product"
                inputProps={{ 'aria-label': 'product' }}
                id = "product"
                name = "product"
                onChange={handleChange}
            />
            <IconButton type="button" sx={{ p: '10px' }} aria-label="search"
                onClick = {handleSearch}
            >

                <SearchIcon />
            </IconButton>
            <Divider sx={{ height: 32, m: 0.5 }} orientation="vertical" />
            <IconButton color="primary" sx={{ p: '10px' }} aria-label="directions">
                <DirectionsIcon />
            </IconButton>
        </Paper>
    );
}